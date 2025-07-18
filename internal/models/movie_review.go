package models

import (
	"github.com/google/uuid"
	"time"
)

type MovieReview struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	MovieID    string    `json:"movie_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Rating     int       `json:"rating" gorm:"check between 1 AND 10"`
	Comment    string    `json:"comment"`
	ReviewedAt time.Time `json:"reviewed_at"`
}
