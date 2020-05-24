package dto

import "encoding/json"

//Error -> dto to encapsulate error
type Error struct {
	Message string `json:"message"`
}

//JSON -> converts error dto to json string
func (e *Error) JSON() string {
	jsonData, _ := json.Marshal(e)
	return string(jsonData)
}
