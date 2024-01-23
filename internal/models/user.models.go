package models

import (
	"time"
)

type User struct {
	UserID    uint8
	Username  string
	Name      string
	Password  string
	Access    string
	CreatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
}
