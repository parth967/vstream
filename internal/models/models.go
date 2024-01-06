package models

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	ID             int
	Email          string
	Password       string
	FavoritePhrase string
}
