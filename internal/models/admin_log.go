package models

import (
	"github.com/google/uuid"
	"time"
)

type AdminLog struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	AdminID   uuid.UUID `json:"admin_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Action    string    `json:"action" gorm:"type:varchar(255);default:'update'"`
	MetaData  string    `json:"meta_data" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"check between 1 AND 10"`
}
