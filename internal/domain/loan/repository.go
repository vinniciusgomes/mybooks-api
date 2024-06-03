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
	ReturnLoan(loanID string) error
}

type loanRepositoryImp struct {
	db *gorm.DB
}

// NewLoanRepository creates a new instance of the LoanRepository interface.
//
// It takes a *gorm.DB parameter, which represents the database connection.
// It returns a LoanRepository pointer, which is an implementation of the LoanRepository interface.
func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepositoryImp{
		db: db,
	}
}

// CreateLoan creates a new loan in the database.
//
// It takes a pointer to a model.Loan object as a parameter, which represents the loan to be created.
// It returns an error if there was a problem creating the loan, such as a book not found or a book already borrowed.
// If the loan is created successfully, it returns nil.
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

// GetAllLoans retrieves all loans from the loan repository.
//
// It returns a pointer to a slice of model.Loan and an error if any.
func (r *loanRepositoryImp) GetAllLoans() (*[]model.Loan, error) {
	var loans []model.Loan

	if err := r.db.Find(&loans).Error; err != nil {
		return nil, err
	}

	return &loans, nil
}

// ReturnLoan updates the loan with the given ID to mark it as returned.
//
// Parameters:
// - loanID: The ID of the loan to be marked as returned.
//
// Returns:
// - error: An error if the update operation fails.
func (r *loanRepositoryImp) ReturnLoan(loanID string) error {
	if err := r.db.Model(&model.Loan{}).Where("id = ?", loanID).Update("is_returned", true).Error; err != nil {
		return err
	}

	return nil
}
