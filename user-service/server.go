package main

import (
	"user-service/config"

	"github.com/gin-gonic/gin"
)

func startRouter() {
	router := gin.Default()
	Start(router)
	router.Run(":8081")
}

func main() {
	// Load Config
	config.Load()

	// Start DB
	config.StartDB()
	db := config.GetDB()
	defer db.Close()

	// Start Router
	startRouter()
}
