package server

import (
	"github.com/gabivega362/gabthebox/server/routes"
	"github.com/gin-gonic/gin"
)

func ListenAndServe(address string) {
	router := gin.Default()
	router.LoadHTMLGlob("./server/templates/*")
	routes.SetRoutes(router)

	router.Run(address)
}
