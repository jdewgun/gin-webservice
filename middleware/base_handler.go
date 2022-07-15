package middleware

import (
	"github.com/gin-gonic/gin"
	"runtime"
)

func SayHello(context *gin.Context) {
	context.JSON(200, gin.H{
		"payload": "Hello Gin",
	})
}

func DisplayRuntime(context *gin.Context) {
	context.JSON(200, runtime.GOOS)
}
