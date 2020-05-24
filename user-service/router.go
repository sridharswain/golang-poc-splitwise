package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func startStatusRoute(router *gin.RouterGroup) {
	// Gives status of server
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Server is Running",
			"time":   time.Now().String(),
		})
	})
}

// func start

//Start -> starts all routers
func Start(router *gin.Engine) {
	startStatusRoute(router.Group("/status"))
}
