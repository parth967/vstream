package models

import (
	"time"
)

type Document struct {
	RepoId       uint8
	RepoLocation string
	RepoName     string
	CreatedAt    time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
}
