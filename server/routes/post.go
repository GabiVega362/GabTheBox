package routes

import "github.com/gin-gonic/gin"

var routesPOST = map[string]func(*gin.Context){

	"/login": func(ctx *gin.Context) {
		ctx.String(200, "Esto es el login por POST")
	},
}
