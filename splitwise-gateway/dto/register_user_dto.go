package dto

import (
	"encoding/json"

	"github.com/fatih/structs"
)

// RegisterUserDto -> dto to register a user
type RegisterUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//JSON -> converts dto to Json
func (registerUserDto *RegisterUserDto) JSON() []byte {
	jsonData, _ := json.Marshal(registerUserDto)
	return jsonData
}

//Map -> returns map of the struct
func (registerUserDto *RegisterUserDto) Map() map[string]interface{} {
	return structs.Map(registerUserDto)
}
