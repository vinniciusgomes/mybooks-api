package loan

import (
	"errors"
	"fmt"
	"mybooks/internal/infrastructure/model"

	"gorm.io/gorm"
)

type LoanRepository interface {
	CreateLoan(loan *model.Loan) error
	GetAllLoans() (*[]model.Loan, error)
	GetLoanByBookID(bookID string) (*model.Loan, error)
	ReturnLoan(loanID string) error
}

type loanRepositoryImp struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepositoryImp{
		db: db,
	}
}

func (r *loanRepositoryImp) CreateLoan(loan *model.Loan) error {
	if err := r.db.First(&model.Book{}, "id = ?", loan.BookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("book not found")
		}

		return err
	}

	var existingLoan model.Loan
	if err := r.db.Where("book_id = ? AND is_returned = false", loan.BookID).First(&existingLoan).Error; err == nil {
		return errors.New("book already borrowed")
	}

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

func (r *loanRepositoryImp) ReturnLoan(loanID string) error {
	if err := r.db.Model(&model.Loan{}).Where("id = ?", loanID).Update("is_returned", true).Error; err != nil {
		return err
	}

	return nil
}
