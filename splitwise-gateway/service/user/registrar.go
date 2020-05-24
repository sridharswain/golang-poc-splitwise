package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Print -> prints Hello world
func Print() {
	fmt.Println("Helo World")
}

//Login -> logs in a user
func Login(c *gin.Context) {

}

//Logout -> logs out a user
func Logout(c *gin.Context) {

}
