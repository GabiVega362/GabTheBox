package routes

import (
	"fmt"
	"net/http"

	"github.com/gabivega362/gabthebox/app/config"
	"github.com/gin-gonic/gin"
) // paquete para crear servidores WEB

// routesPOST es un array asociativo (map) que asocia cada ruta(clave) con las acciones a realizar antes de devolver la respuesta (valor)
var routesPOST = map[string]func(*gin.Context){

	"/login": func(gctx *gin.Context) {
		// Guardamos los datos del formulario en la variable "data"
		var data struct {
			User string `form:"user"`
			Pass string `form:"pass"`
		}
		gctx.Bind(&data)

		// obtenemos el cliente de la base de datos desde la configuracion de la aplicacion que hemos guardado en Gin
		databaseClient := gctx.MustGet("Config").(*config.Config).Database
		// logeamos al usuario
		if id, err := databaseClient.UserLogin(data.User, data.Pass); err != nil {
			gctx.Redirect(http.StatusFound, "/login?error=true")
		} else {
			gctx.SetCookie("GTBSESSID", id, 3600, "/", "", false, true)
			gctx.Redirect(http.StatusFound, "/lab")
		}
	},
	"/register": func(gctx *gin.Context) {
		// cuando se accede a la ruta "/register" por POST
		// obtenemos los datos del formulario y los guardamos dentro de una variable llamada data
		var data struct {
			User  string `form:"user"`
			Email string `form:"email"`
			Pass  string `form:"pass"`
		}
		// Bind es una función que tiene Gin para rellenar la variable especificada entre parentesis, Bind nos obliga a hacerlo con punteros por lo que se utiliza el & el cual CREA un puntero (data)
		gctx.Bind(&data)
		// obtenemos el cliente de la base de datos desde la configuracion de la aplicacion que hemos guardado en Gin
		databaseClient := gctx.MustGet("Config").(*config.Config).Database
		// registramos el usuario o error en la DB
		if success, err := databaseClient.UserResgister(data.User, data.Email, data.Pass); err != nil || !success {
			// Usuario duplicado
			gctx.Redirect(http.StatusFound, "/register?error=true")
		} else {
			// Usuario creado
			gctx.Redirect(http.StatusFound, "/login?success=true")
		}
	},

	"/lab": func(gctx *gin.Context) {
		// comprobamos que el usuario esta autenticado
		if !IsAuthenticated(gctx) {
			gctx.String(http.StatusForbidden, "You shall not pass!")
			return
		}
		// obtenemos los datos del formulario
		var data struct {
			Action string `form:"action"`
			Lab    string `form:"lab"`
		}
		gctx.Bind(&data)
		// obtenemos la configuracion de la aplicacion que hemos guardado en Gin
		cfg := gctx.MustGet("Config").(*config.Config)
		// obtenemos la sesion del usuario
		sessid, _ := GetSessionId(gctx)
		// comprobamos la accion a realizar
		var gErr error
		switch data.Action {
		case "Encender":
			// comprobamos si ya ha deployeado una
			if _, alreadyDeployed, err := cfg.Database.LabsGetContainerByUser(sessid); err != nil || alreadyDeployed {
				fmt.Println(alreadyDeployed, err)
				gctx.Redirect(http.StatusFound, "/lab?error")
				return
			}
			// obtenemos la imagen del laboratorio
			image, err := cfg.Database.LabsGetImageById(data.Lab)
			if err != nil {
				gctx.String(http.StatusInternalServerError, "Internal server error: 0db05")
				return
			}
			// Llamamos la funcion publica, definida en labs.go para levantar una nueva instancia
			container, port, err := cfg.Docker.StartLab(sessid, image)
			if err != nil {
				gErr = err
				break
			}
			// guardamos la relacion entre el usuario y el laboratorio en la BBDD
			if err = cfg.Database.LabsUserStarted(sessid, data.Lab, container, port); err != nil {
				_ = cfg.Docker.StopLab(container)
				gctx.String(http.StatusInternalServerError, "Internal server error: 0x21")
				return
			}

		case "Detener":
			// obtenemos el ID del contenedor
			container, alreadyDeployed, err := cfg.Database.LabsGetContainerByUser(sessid)
			if err != nil {
				gctx.String(http.StatusInternalServerError, "Internal server error: 0x22")
				return
			}
			if !alreadyDeployed {
				gctx.Redirect(http.StatusFound, "/lab")
				return
			}
			// borramos la relacion entre el usuario y la base de datos
			if err = cfg.Database.LabsUserStopped(sessid); err != nil {
				gctx.String(http.StatusInternalServerError, "Internal server error: 0x22")
				return
			}
			// paramos el contenedor
			gErr = cfg.Docker.StopLab(container)
		}

		if gErr != nil {
			gctx.String(http.StatusInternalServerError, "Internal server error: 0x20")
			panic(gErr)
		}
		// redirigimos al usuario a la pagina del lab
		gctx.Redirect(http.StatusFound, "/lab?success")
	},
}
