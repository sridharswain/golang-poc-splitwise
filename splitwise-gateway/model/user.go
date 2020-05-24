package model

import "github.com/jinzhu/gorm"

// User -> dto of the user
type User struct {
	gorm.Model
	Name  string
	Email string
}
