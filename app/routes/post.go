package routes

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/gabivega362/gabthebox/app/config"
	"github.com/gin-gonic/gin"
) // paquete para crear servidores WEB

// Register es una estructura que contiene los datos del formulario HTML que se pasan por POST al reguistrarte
type Register struct {
	// Password es la password de usuario
	Password string `form:"pass"`
	// Email es el email de usuario
	Email string `form:"email"`
}
type LabParams struct {
	Action     string `form:"action"`
	Enviroment string `form:"enviroment"`
}

// routesPOST es un array asociativo (map) que asocia cada ruta(clave) con las acciones a realizar antes de devolver la respuesta (valor)
var routesPOST = map[string]func(*gin.Context){

	"/login": func(gctx *gin.Context) {
		// cuando se accede a la ruta "/login" por POST
		// devuelve el codigo de estado 200 OK junto a un mensaje en texto plano
		gctx.String(200, "Esto es el login por POST")
	},
	"/register": func(gctx *gin.Context) {
		// cuando se accede a la ruta "/register" por POST
		// obtenemos los datos del formulario y los guardamos dentro de una variable llamada data
		var data Register
		// Bind es una función que tiene Gin para rellenar la variable especificada entre parentesis, Bind nos obliga a hacerlo con punteros por lo que se utiliza el & el cual CREA un puntero (data)
		gctx.Bind(&data)
		// obtenemos la configuración del programa
		cfg := gctx.MustGet("Config").(*config.Config)
		// Comprobamos que el usuario no exista
		stmt, err := cfg.Database.Prepare("SELECT id FROM users WHERE email = $1")
		if err != nil {
			gctx.String(http.StatusInternalServerError, "Internal Server Error = 0db01")
			return
		}
		// Cierra la sentencia preparada antes de salir de la función
		defer stmt.Close()
		id := -1
		err = stmt.QueryRow(data.Email).Scan(&id)
		if err != nil && err != sql.ErrNoRows {
			gctx.String(http.StatusInternalServerError, "Internal Server Error = 0db02")
			return
		} 
		if id != -1 {
			// si el usuario existe, redirigimos a la pagina de registro con un error
			gctx.Redirect(http.StatusFound, "/register?error")
			return
		}
		
		// si el usuario no existe, lo creamos
		stmt2, err := cfg.Database.Prepare("INSERT INTO users (email, password) VALUES ($1, $2)")
		if err != nil {
			gctx.String(http.StatusInternalServerError, "Internal Server Error = 0db03")
			return
		}
		defer stmt2.Close()
		_, err = stmt2.Exec(data.Email, data.Password)
		if err != nil {
			gctx.String(http.StatusInternalServerError, "Internal Server Error = 0db04")
			return
		}
		// redirigimos el usuario a la pagina login
		gctx.Redirect(http.StatusFound, "/login?success=true")
	},
	"/lab": func(gctx *gin.Context) {
		// obtenemos los datos del formulario
		var params LabParams
		gctx.Bind(&params)
		// obtenemos la configuracion de la aplicación
		cfg, exists := gctx.Get("Config")
		if !exists {
			gctx.String(http.StatusInternalServerError, "Internal Server Error:0x1")
			return
		}
		dockerClient := cfg.(*config.Config).Docker

		// comprobamos la accion (FIXME: check user, check if already deployed)
		switch params.Action {
		case "Start":
			// obtenemos los contenedores que están en ejecucion
			containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{})
			if err != nil {
				gctx.String(http.StatusInternalServerError, "Internal Server Error:0x2")
				return
			}
			// comprobamos si el contenedor ya está en ejecucion
			i := 0
			found := false
			for i < len(containers) && !found {
				if containers[i].Image == params.Enviroment {
					found = true
				}
				i++
			}
			// si el contenedor no está en ejecucion lo iniciamos
			if !found {
				// buscamos la imagen
				_, err := dockerClient.ImagePull(context.Background(), params.Enviroment, types.ImagePullOptions{})
				if err != nil {
					gctx.String(http.StatusInternalServerError, "Internal Server Error:0x3")
					return
				}
				// creamos el contenedor
				res, err := dockerClient.ContainerCreate(context.Background(), &container.Config{
					Image: params.Enviroment,
				}, nil, nil, nil, strings.ReplaceAll(params.Enviroment, "/", "_"))
				if err != nil {
					panic(err)
					gctx.String(http.StatusInternalServerError, "Internal Server Error:0x4")
					return
				}
				//iniciamos el contendor
				err = dockerClient.ContainerStart(context.Background(), res.ID, types.ContainerStartOptions{})
				if err != nil {
					gctx.String(http.StatusInternalServerError, "Internal Server Error:0x5")
					return
				}
			}
		}
		// redirigimos al usuario a la pagina del lab
		gctx.Redirect(http.StatusFound, "/lab")
	},
}
