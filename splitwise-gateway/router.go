   z  package main

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"

	"splitwise-gateway/service/user"
)

//startStatusRoutes -> starts all status routes
func startStatusRoutes(router *gin.RouterGroup) {

	// Gives status of server
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Server is Running",
			"time":   time.Now().String(),
		})
	})
}

//startUserRegistrarRoutes -> starts all registrar routes
func startUserRegistrarRoutes(router *gin.RouterGroup) {
	router.POST("/logout", user.Logout)
	router.POST("/login", user.Login)
}

func startUserRouter(router *gin.RouterGroup) {
	router.POST("/register", user.Register)
	router.GET("/:id", user.GetUser)
	router.PUT("/:id", user.UpdateUser)
	router.DELETE("/:id", user.DeleteUser)
}

//StartRouter -> starts router
func StartRouter(router *gin.Engine) {
	
	startStatusRoutes(router.Group("/status"))
	startUserRegistrarRoutes(router.Group("/registrar"))
	startUserRouter(router.Group("/user"))
}
