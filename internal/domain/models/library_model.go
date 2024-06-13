package models

import (
	"time"

	"github.com/google/uuid"
)

type Library struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey" validate:"required,uuid4"`
	Name        string    `json:"name" gorm:"not null;size:100" validate:"required,min=1,max=100"`
	Description string    `json:"description" gorm:"size:1024" validate:"max=1024"`
	UserID      uuid.UUID `json:"-" gorm:"type:uuid;not null;index"`
	User        User      `json:"-" gorm:"foreignKey:UserID"`
	Books       []Book    `json:"books" gorm:"many2many:book_library;"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
