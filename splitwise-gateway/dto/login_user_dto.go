package dto

// LoginUserDto -> dto for user login
type LoginUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
