package main

import "github.com/gin-gonic/gin"

func ejercicio1 (ctx *gin.Context) {
		ctx.String(200, "pong") 
}

func ejercicio2 (ctx *gin.Context) {
	nombre := ctx.Query("nombre")
	apellido := ctx.Query("apellido")
	ctx.String(200, "Hola %s %s", nombre, apellido)
}

func main() {
	router := gin.Default()
	router.GET("/ping", ejercicio1)
	router.POST("/saludo", ejercicio2)
	router.Run()
}