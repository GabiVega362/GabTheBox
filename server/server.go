package server

import (
	"github.com/gabivega362/gabthebox/server/config"
	"github.com/gabivega362/gabthebox/server/routes"
	"github.com/gin-gonic/gin"
)

func ListenAndServe(ctx *config.Context) {

	router := gin.Default()
	router.LoadHTMLGlob("./server/templates/*")
	routes.SetRoutes(router)

	router.Run(ctx.Args.Address)
}
