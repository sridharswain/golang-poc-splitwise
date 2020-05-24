package user

import (
	"fmt"
	"net/http"
	"splitwise-gateway/dto"

	"github.com/gin-gonic/gin"
)

//Register -> signs up a user
func Register(c *gin.Context) {
	var registerDto dto.RegisterUserDto
	if err := c.ShouldBindJSON(&registerDto); err != nil {
		error := &dto.Error{
			Message: "Invalid Input",
		}
		c.JSON(http.StatusBadRequest, error.JSON())
		return
	}
	fmt.Println(registerDto.JSON())
	c.Data(http.StatusOK, "application/json", registerDto.JSON())
}

//GetUser -> gets user from user service
func GetUser(c *gin.Context) {

}

//UpdateUser -> updates user's data
func UpdateUser(c *gin.Context) {

}

//DeleteUser -> deletes user
func DeleteUser(c *gin.Context) {

}
