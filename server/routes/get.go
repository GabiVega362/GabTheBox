package routes

import "github.com/gin-gonic/gin"

var routesGET = map[string]func(*gin.Context){
	"/": func(ctx *gin.Context) {
		name := ctx.Query("name")

		ctx.HTML(200, "index.tmpl", gin.H{
			"name": name,
		})
	},

	"/login": func(ctx *gin.Context) {
		ctx.HTML(200, "login.tmpl", gin.H{})
	},

	"/register": func(ctx *gin.Context) {
		ctx.HTML(200, "register.tmpl", gin.H{})
	},
}
