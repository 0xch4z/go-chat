package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data)
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
