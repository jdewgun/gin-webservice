package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jdewgun/gin-webservice/middleware"
)

func main() {
	router := gin.Default()
	router.GET("/hello", middleware.SayHello)
	router.GET("/os", middleware.DisplayRuntime)
	router.Run(":5000")
}
