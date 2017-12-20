package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// gin.SetMode(gin.ReleaseMode)

	initializeRoutes()

	router.Run(":8080")
}
