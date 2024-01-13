package models

import (
	"time"
)

type User struct {
	UserID    uint8
	Username  string
	Password  string
	Access    string
	CreatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
}

type LoginRequet struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
