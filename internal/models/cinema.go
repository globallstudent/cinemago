package models

import (
	"github.com/google/uuid"
	"time"
)

type Cinema struct {
	ID            uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name          string    `json:"name" gorm:"unique;not null"`
	Address       string    `json:"address" gorm:"not null"`
	City          string    `json:"city" gorm:"not null"`
	Latitude      float64   `json:"latitude" gorm:"not null"`
	Longitude     float64   `json:"longitude" gorm:"not null"`
	Description   string    `json:"description" gorm:"not null"`
	Website       string    `json:"website" gorm:"not null"`
	IsActive      bool      `json:"is_active" gorm:"not null"`
	AverageRating float64   `json:"average_rating" gorm:"not null;type:numeric(3, 2);default:0.0"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"not null"`
}
