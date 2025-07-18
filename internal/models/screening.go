package models

import (
	"github.com/google/uuid"
	"time"
)

type Screening struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;type:uuid"`
	MovieID     uuid.UUID `json:"movie_id" gorm:"type:uuid"`
	CinemaID    uuid.UUID `json:"cinema_id" gorm:"type:uuid"`
	StartTime   time.Time `json:"start_time"`
	RoomNumber  int       `json:"room_number"`
	TicketPrice float64   `json:"ticket_price"`
	Capacity    int       `json:"capacity"`
	CreatedAt   time.Time `json:"created_at"`
}
