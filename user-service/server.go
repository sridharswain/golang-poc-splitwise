package main

import (
	"user-service/config"
	"user-service/model"

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
	for i := 0; i < 100; i++ {
		go func() {
			user := model.User{Name: "User1", Email: "user1@gmail.com"}
			db.Create(&user)
		}()
	}
	// Start Router
	startRouter()
}
