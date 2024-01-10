package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

var validate = validator.New()

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email     string
	Password  string
	Access    string
	CreatedAt time.time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
}
