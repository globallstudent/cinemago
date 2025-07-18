package models

import (
	"github.com/google/uuid"
	"time"
)

type Booking struct {
	ID            uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID        uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	ScreeningID   uuid.UUID `json:"screening_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	SeatNumber    int       `json:"seat_number"`
	Status        string    `json:"status" gorm:"type:varchar(255);default:'confirmed'"`
	PaymentStatus string    `json:"payment_status" gorm:"type:varchar(255);default:'confirmed'"`
	PaymentID     uuid.UUID `json:"payment_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	TicketCode    string    `json:"ticket_code" gorm:"uniqueIndex;type:varchar(255);default:'standard'"`
	QRCodeURL     string
	BookedAt      time.Time `json:"booked_at"`
	CanceledAt    time.Time `json:"canceled_at"`
}
