package model

// User -> entity of the user
type User struct {
	id    int32  `gorm:"primary_key"`
	Name  string `gorm:"column:user_name"`
	Email string `gorm:"column:email"`
}

// TableName Get table Name
func (user *User) TableName() string {
	return "user_profile"
}
