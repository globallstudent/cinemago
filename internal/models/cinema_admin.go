package models

import "github.com/google/uuid"

type CinemaAdmin struct {
	ID       uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CinemaID uuid.UUID `json:"cinema_id" gorm:"default:uuid_generate_v4()"`
	Role     string    `json:"role"`
}
