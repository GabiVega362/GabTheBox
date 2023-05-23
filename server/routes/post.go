package routes

import "github.com/gin-gonic/gin" // paquete para crear servidores WEB

// Register es una estructura que contiene los datos del formulario HTML que se pasan por POST al reguistrarte
type Register struct {
	// Username es el nombre de usuario
	Username string `form:"user"`
	// Password es la password de usuario
	Password string `form:"pass"`
	// Email es el email de usuario
	Email string `form:"email"`
}
// routesPOST es un array asociativo (map) que asocia cada ruta(clave) con las acciones a realizar antes de devolver la respuesta (valor)
var routesPOST = map[string]func(*gin.Context){

	"/login": func(ctx *gin.Context) {
		// cuando se accede a la ruta "/login" por POST
		// devuelve el codigo de estado 200 OK junto a un mensaje en texto plano
		ctx.String(200, "Esto es el login por POST")
	},
	"/register": func(ctx *gin.Context) {
		// cuando se accede a la ruta "/register" por POST
		// obtenemos los datos del formulario y los guardamos dentro de una variable llamada data
		var data Register
		// Bind es una función que tiene Gin para rellenar la variable especificada entre parentesis, Bind nos obliga a hacerlo con punteros por lo que se utiliza el & el cual CREA un puntero (data)
		ctx.Bind(&data)
		// devuelve el codigo de estado 200 OK junto a un mensaje en texto plano
		ctx.String(200, "Me ha llegado que el usuario " + data.Username + "con el email " + data.Email + " y contraseña " + data.Password + " se ha registrado")
	},
}
