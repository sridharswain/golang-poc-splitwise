package dto

import (
	"encoding/json"
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
