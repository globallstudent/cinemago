package models

import "github.com/google/uuid"

type Seat struct {
	ID              uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CinemaID        uuid.UUID `json:"cinema_id" gorm:"type:uuid"`
	RoomNumber      int       `json:"room_number"`
	SeatType        string    `json:"seat_type" gorm:"type:varchar(255);default:'standard'"`
	PriceMultiplier float64   `json:"price_multiplier" gorm:"default:1.0"`
}
