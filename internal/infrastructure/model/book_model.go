package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey" validate:"required,uuid4"`
	Title         string    `json:"title" gorm:"not null;size:100" validate:"required,min=1,max=100"`
	Author        string    `json:"author" gorm:"not null;size:100" validate:"required,min=1,max=100"`
	Description   string    `json:"description" gorm:"size:1024" validate:"max=1024"`
	Cover         string    `json:"cover" gorm:"size:1024" validate:"max=1024,url"`
	Genre         string    `json:"genre" gorm:"size:100" validate:"max=100"`
	ISBN          string    `json:"isbn" gorm:"size:20" validate:"max=20"`
	PublishedDate string    `json:"published_date" gorm:"size:20" validate:"max=20"`
	Language      string    `json:"language" gorm:"size:10" validate:"max=10"`
	Pages         int       `json:"pages" gorm:"default:0" validate:"min=0"`
	Read          bool      `json:"read" gorm:"default:false"`
	Libraries     []Library `json:"-" gorm:"many2many:book_library;"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
