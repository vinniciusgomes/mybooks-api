package model

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey" validate:"required,uuid4"`
	BookID       string    `json:"book_id" gorm:"not null;size:36" validate:"required,min=1,max=36"`
	LoanDate     string    `json:"loan_date" gorm:"not null;size:20" validate:"required,min=1,max=20"`
	BorrowerName string    `json:"borrower_name" gorm:"not null;size:100" validate:"required,min=1,max=100"`
	IsReturned   bool      `json:"is_returned" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
