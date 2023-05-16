package routes

import "github.com/gin-gonic/gin"

// Enrutador automatico para todas las rutas definidas en routesGET
func SetRoutes(router *gin.Engine) {
	for path, handler := range routesGET {
		router.GET(path, handler)
	}

	for path, handler := range routesGET {
		router.POST(path, handler)
	}
}
