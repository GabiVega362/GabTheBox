package routes

import (
	"net/http"

	"github.com/gabivega362/gabthebox/app/config"
	"github.com/gin-gonic/gin"
) // paquete para crear servidores WEB

// routesGET es un array asociativo (map) que asocia cada ruta(clave) con las acciones a realizar antes de devolver la respuesta (valor)
var routesGET = map[string]func(*gin.Context){
	//	PATH	ACCION
	"/": func(gctx *gin.Context) {
		// cuando se accede a la ruta "/"
		// devuelve el código de estado HTTP 200 OK y la plantilla index.tmpl
		gctx.HTML(200, "index.tmpl", gin.H{})
	},

	"/login": func(gctx *gin.Context) {
		// cuando se accede a la ruta "/login"
		// devuelve el código de estado HTTP 200 OK y la plantilla login.tmpl
		gctx.HTML(200, "login.tmpl", gin.H{})
	},

	"/logout": func(gctx *gin.Context) {
		gctx.SetCookie("GTBSESSID", "", -1, "/", "", false, true)
		gctx.Redirect(http.StatusFound, "/")
	},

	"/register": func(gctx *gin.Context) {
		// cuando se accede a la ruta "/register"
		// devuelve el código de estado HTTP 200 OK y la plantilla register.tmpl
		gctx.HTML(200, "register.tmpl", gin.H{})
	},

	"/lab": func(gctx *gin.Context) {

		databaseClient := gctx.MustGet("Config").(*config.Config).Database
		if labs, err := databaseClient.GetAllLabs(); err != nil {
			gctx.String(http.StatusInternalServerError, "Internal Server Error: 0db04")
		} else {
			// devuelve el código de estado HTTP 200 OK y la plantilla lab.tmpl
			gctx.HTML(http.StatusOK, "lab.tmpl", gin.H{
				"Labs": labs,
			})
		}

	},
}
