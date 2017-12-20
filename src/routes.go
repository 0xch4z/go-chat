package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Route type
type Route string

// Router constants
const (
	roomRoute   = "/room/:roomid"
	streamRoute = "/stream/:roomid"
)

func initializeRoutes() {
	router = gin.Default()

	router.Use(static.Serve("/static", static.LocalFile("static", false)))

	router.GET("/foo", testRoute)

	roomRoute := router.Group("/room")
	{
		roomRoute.GET("/:room_id", roomGET)
		roomRoute.POST("/:room_id", roomPOST)
		roomRoute.DELETE("/:room_id", roomDELETE)
	}

	router.GET("/stream/:room_id", stream)

	router.LoadHTMLGlob("templates/*")
}

func testRoute(c *gin.Context) {
	render(c, gin.H{}, "index.tmpl")
}
