package server

import (
	"github.com/gin-gonic/gin" // paquete para crear servidores web

	"github.com/gabivega362/gabthebox/server/config"
	"github.com/gabivega362/gabthebox/server/routes"
	
)

// ListenAndServe pone el servidor a la escucha en la direccion especificada. La dirección está definida dentro de la estructura context. Esta función es bloqueante, por lo que no continuará hasta que el servidor se detenga
func ListenAndServe(ctx *config.Context) {

	// creamos un enrutador WEB con la configuración por defecto que provee el framework Gin
	router := gin.Default()
	// especificamos al enrutador donde se encuentran las plantillas HTML
	router.LoadHTMLGlob("./server/templates/*")
	// añadimos las rutas(GET, POST, etc...) al enrutador
	routes.SetRoutes(router)

	// encendemos el enrutador para que maneje todas las peticiones que le llegan a la dirección especificada
	router.Run(ctx.Args.Address)
}
