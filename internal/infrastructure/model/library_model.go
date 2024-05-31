package model

import (
	"time"

	"github.com/google/uuid"
)

type Library struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey" validate:"required,uuid4"`
	Name        string    `json:"name" gorm:"not null;size:100" validate:"required,min=1,max=100"`
	Description string    `json:"description" gorm:"size:1024" validate:"max=1024"`
	Books       []Book    `json:"books" gorm:"many2many:book_library;"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
