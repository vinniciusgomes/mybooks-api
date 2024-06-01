package loan

import (
	"mybooks/internal/infrastructure/model"

	"gorm.io/gorm"
)

type LoanRepository interface {
	CreateLoan(loan *model.Loan) error
	GetAllLoans() ([]model.Loan, error)
	GetLoanByBookID(bookID string) (*model.Loan, error)
	UpdateLoan(loan *model.Loan) error
	DeleteLoan(bookID string) error
}

type loanRepositoryImp struct {
	db *gorm.DB
}

func (r *loanRepositoryImp) CreateLoan(loan *model.Loan) error {
	if err := r.db.Create(loan).Error; err != nil {
		return err
	}

	return nil
}

func (r *loanRepositoryImp) GetAllLoans() (*[]model.Loan, error) {
	var loans []model.Loan

	if err := r.db.Find(&loans).Error; err != nil {
		return nil, err
	}

	return &loans, nil
}

func (r *loanRepositoryImp) GetLoanByBookID(bookID string) (*model.Loan, error) {
	var loan model.Loan

	if err := r.db.Where("book_id = ?", bookID).First(&loan).Error; err != nil {
		return nil, err
	}

	return &loan, nil
}
