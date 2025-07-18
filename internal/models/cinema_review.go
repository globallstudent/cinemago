package models

import (
	"github.com/google/uuid"
	"time"
)

type CinemaReview struct {
	ID         uuid.UUID  `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID     uuid.UUID  `json:"user_id" gorm:"type:uuid;not null"`
	CinemaID   uuid.UUID  `json:"cinema_id" gorm:"type:uuid;not null"`
	Rating     int        `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 10 "`
	Comment    string     `json:"comment"`
	ReviewedAt *time.Time `json:"reviewed_at"`
}
