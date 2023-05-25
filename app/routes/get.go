package routes

import (
	"github.com/gabivega362/gabthebox/app/docker"
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

	"/register": func(gctx *gin.Context) {
		// cuando se accede a la ruta "/register"
		// devuelve el código de estado HTTP 200 OK y la plantilla register.tmpl
		gctx.HTML(200, "register.tmpl", gin.H{})
	},

	"/lab": func(gctx *gin.Context) {
		// cuando se accede a la ruta "/lab"
		// devuelve el código de estado HTTP 200 OK y la plantilla lab.tmpl
		gctx.HTML(200, "lab.tmpl", gin.H{
			"Enviroments": []docker.Enviroment{
				*docker.NewEnviroment("gability", "GabTheBox", "Vulnerable inviroment for learning web hacking"),
				*docker.NewEnviroment("vulnerables/web-dvwa", "DVWA", "Damn Vulnerable Web Applications"),
			},
		})
	},
}
