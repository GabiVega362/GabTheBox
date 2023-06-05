package routes

import (
	"fmt"
	"net/http"

	"github.com/gabivega362/gabthebox/app/config"
	"github.com/gin-gonic/gin"
) // paquete para crear servidores WEB

type LabParams struct {
	Action     string `form:"action"`
	Enviroment string `form:"enviroment"`
}

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
			gctx.SetCookie("GTBSESSID", fmt.Sprint(id), 3600, "/", "", false, true)
			gctx.Redirect(http.StatusFound, "/lab")
		}
	},
	"/register": func(gctx *gin.Context) {
		// cuando se accede a la ruta "/register" por POST
		// obtenemos los datos del formulario y los guardamos dentro de una variable llamada data
		var data struct {
			Email string `form:"email"`
			Pass  string `form:"pass"`
		}
		// Bind es una función que tiene Gin para rellenar la variable especificada entre parentesis, Bind nos obliga a hacerlo con punteros por lo que se utiliza el & el cual CREA un puntero (data)
		gctx.Bind(&data)
		// obtenemos el cliente de la base de datos desde la configuracion de la aplicacion que hemos guardado en Gin
		databaseClient := gctx.MustGet("Config").(*config.Config).Database
		// registramos el usuario
		if success, err := databaseClient.UserResgister(data.Email, data.Pass); err != nil {
			// Fallo en la base de datos
			gctx.String(http.StatusInternalServerError, "Internal Server Error: 0db3")
		} else if !success {
			// Usuario duplicado
			gctx.Redirect(http.StatusFound, "/register?error=true")
		} else {
			// Usuario creado
			gctx.Redirect(http.StatusFound, "/login?success=true")
		}
	},

	"/lab": func(gctx *gin.Context) {
		// obtenemos los datos del formulario
		var params LabParams
		gctx.Bind(&params)
		// obtenemos el cliente de docker desde la configuracion de la aplicaciones que hemos guardado en Gin
		dockerClient := gctx.MustGet("Config").(*config.Config).Docker
		// comprobamos la accion (FIXME: pass the username)
		var err error
		switch params.Action {
		case "Start":
			// Llamamos la funcion publica, definida en client.go para levantar una nueva instancia
			err = dockerClient.StartLab("nobody", params.Enviroment)
		case "Stop":
			err = dockerClient.StopLab("nobody", params.Enviroment)
		}
		if err != nil {
			gctx.String(http.StatusInternalServerError, "Internal server error: 0x20")
		}
		// redirigimos al usuario a la pagina del lab
		gctx.Redirect(http.StatusFound, "/lab")
	},
}
