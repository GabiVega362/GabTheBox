package routes

import "github.com/gin-gonic/gin" // paquete para crear servidores web

// SetRoutes recibe un enrutador WEB y le añade las rutas GET y POST
func SetRoutes(router *gin.Engine) {
	// para cada clave-valor (path-handler) en el array asociativo routesGET, lo añadimos al enrutador como rutas GET
	for path, handler := range routesGET {
		router.GET(path, handler)
	}
	
	// para cada clave-valor (path-handler) en el array asociativo routesPOST, lo añadimos al enrutador como rutas POST
	for path, handler := range routesPOST {
		router.POST(path, handler)
	}
}
