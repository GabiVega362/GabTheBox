package server

import (
	"github.com/gin-gonic/gin" // paquete para crear servidores web

	"github.com/gabivega362/gabthebox/app/config"
	"github.com/gabivega362/gabthebox/app/routes"
)

// ListenAndServe pone el servidor a la escucha en la direccion especificada. La dirección está definida dentro de la estructura context. Esta función es bloqueante, por lo que no continuará hasta que el servidor se detenga
func ListenAndServe(cfg *config.Config) {

	// creamos un enrutador WEB con la configuración por defecto que provee el framework Gin
	router := gin.Default()
	// le especificamos al enrutador cual es nuestro contexto
	router.Use(func (gctx *gin.Context) {
		gctx.Set("Config", cfg)
	})
	// especificamos al enrutador donde se encuentran las plantillas HTML
	router.LoadHTMLGlob("./app/templates/*")
	// le especificamos donde estan los archivos estaticos (img, css,etc)
	router.Static("/static", "./app/public")
	// añadimos las rutas(GET, POST, etc...) al enrutador
	routes.SetRoutes(router)

	// encendemos el enrutador para que maneje todas las peticiones que le llegan a la dirección especificada
	router.Run(cfg.Args.Address)
}
