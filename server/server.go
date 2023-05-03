package server

import "github.com/gin-gonic/gin"

func ListenAndServe(address string) {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Bienvenidos a GabTheBox!",
		})
	})

	router.Run(address)
}
