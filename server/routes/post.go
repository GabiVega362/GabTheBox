package routes

import "github.com/gin-gonic/gin"

type Register struct {
	Username string `form:"user"`
	Password string `form:"pass"`
	Email string `form:"email"`
}

var routesPOST = map[string]func(*gin.Context){

	"/login": func(ctx *gin.Context) {
		ctx.String(200, "Esto es el login por POST")
	},
	"/register": func(ctx *gin.Context) {
		var data Register
		ctx.Bind(&data)
		ctx.String(200, "Me ha llegado que el usuario " + data.Username + "con el email " + data.Email + " y contraseña " + data.Password + " se ha registrado")
	},
}
