package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

type CinemaRequest struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name            string
	Address         string
	City            string
	ContactEmail    string
	ContactPhone    string
	Website         string
	Description     string
	DocumentURLs    pq.StringArray `gorm:"type:text[]"`
	Status          string         `gorm:"default:pending"`
	RejectionReason string
	ReviewedByID    *string
	ReviewedAt      *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
