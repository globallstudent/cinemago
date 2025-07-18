package models

import (
	"github.com/google/uuid"
	"time"
)

type Movie struct {
	ID              uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Genres          []string  `json:"genres"`
	Directors       []string  `json:"directors"`
	DurationMinutes int       `json:"duration_minutes"`
	Language        string    `json:"language"`
	ReleaseDate     time.Time `json:"release_date"`
	PosterURL       string    `json:"poster_url"`
	TrailerURL      string    `json:"trailer_url"`
	CreatedAt       time.Time `json:"created_at"`
}
