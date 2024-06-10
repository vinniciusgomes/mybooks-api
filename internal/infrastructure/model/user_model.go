package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Email     string         `json:"email" gorm:"unique;not null;size:100" validate:"required,min=1,max=100"`
	Password  string         `json:"password" gorm:"not null;size:100" validate:"required,min=1,max=100"`
	Books     []Book         `json:"books" gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
