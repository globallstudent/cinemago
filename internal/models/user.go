package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name          string    `gorm:"not null"`
	Email         string    `gorm:"unique;not null"`
	Phone         string
	PasswordHash  string `gorm:"not null"`
	Role          string `gorm:"type:text;default:user"`
	IsActive      bool   `gorm:"default:true"`
	LasLogin      *time.Time
	EmailVerified bool `gorm:"default:true"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
