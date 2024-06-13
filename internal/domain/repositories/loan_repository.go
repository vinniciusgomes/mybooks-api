package repositories

import (
	"errors"
	"fmt"
	"mybooks/internal/domain/models"

	"gorm.io/gorm"
)

type LoanRepository interface {
	CreateLoan(loan *models.Loan) error
	GetAllLoans(userID string) (*[]models.Loan, error)
	ReturnLoan(userID, loanID string) error
}

type loanRepositoryImp struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepositoryImp{
		db: db,
	}
}

// CreateLoan creates a new loan in the database.
//
// It takes a pointer to a models.Loan object as a parameter, which represents the loan to be created.
// It returns an error if there was a problem creating the loan, such as a book not found or a book already borrowed.
// If the loan is created successfully, it returns nil.
func (r *loanRepositoryImp) CreateLoan(loan *models.Loan) error {
	if err := r.db.First(&models.Book{}, "id = ?", loan.BookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("book not found")
		}

		return err
	}

	var existingLoan models.Loan
	if err := r.db.Where("book_id = ? AND is_returned = false", loan.BookID).First(&existingLoan).Error; err == nil {
		return errors.New("book already borrowed")
	}

	if err := r.db.Create(loan).Error; err != nil {
		return err
	}

	return nil
}

// GetAllLoans retrieves all loans for a given user from the loan repository.
//
// Parameters:
// - userID: the ID of the user whose loans are being retrieved.
//
// Returns:
// - *[]models.Loan: a pointer to a slice of models.Loan representing the loans for the user, or nil if there are no loans.
// - error: an error if there was a problem retrieving the loans.
func (r *loanRepositoryImp) GetAllLoans(userID string) (*[]models.Loan, error) {
	var loans []models.Loan

	if err := r.db.Where("user_id = ?", userID).Find(&loans).Error; err != nil {
		return nil, err
	}

	return &loans, nil
}

// ReturnLoan updates the "is_returned" field of a loan record in the database.
//
// Parameters:
// - userID: the ID of the user who is returning the loan.
// - loanID: the ID of the loan being returned.
//
// Returns:
// - error: an error if there was a problem updating the loan record.
func (r *loanRepositoryImp) ReturnLoan(userID, loanID string) error {
	if err := r.db.Model(&models.Loan{}).Where("id = ? AND user_id = ?", loanID, userID).Update("is_returned", true).Error; err != nil {
		return err
	}

	return nil
}
