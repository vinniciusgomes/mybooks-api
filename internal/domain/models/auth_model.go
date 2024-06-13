package models

import (
	"time"

	"github.com/google/uuid"
)

type ValidationToken struct {
	Token     string    `json:"token" gorm:"not null;size:100;index;unique" validate:"required,min=1,max=100"`
	Type      string    `json:"type" gorm:"not null;size:100" validate:"required,min=1,max=100"`
	Valid     bool      `json:"valid" gorm:"not null;default:true"`
	UserID    uuid.UUID `json:"user_id" gorm:"not null;type:uuid;index"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
}
