package loan

import (
	"mybooks/internal/infrastructure/helper"
	"mybooks/internal/infrastructure/model"
	"mybooks/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoanService struct {
	repo LoanRepository
}

// NewLoanService creates a new instance of the LoanService struct.
//
// It takes a LoanRepository parameter, which is used to interact with the database.
// It returns a pointer to the newly created LoanService instance.
func NewLoanService(repo LoanRepository) *LoanService {
	return &LoanService{
		repo: repo,
	}
}

// CreateLoan creates a new loan in the LoanService.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function generates a random ID, binds the JSON request body to a model.Loan struct,
// validates the struct, creates the loan in the repository, and returns the ID of the created loan.
// If any error occurs during the process, it handles the error and returns an appropriate HTTP status code.
func (s *LoanService) CreateLoan(c *gin.Context) {
	loan := new(model.Loan)

	id, err := pkg.GenerateRandomID()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	if err := c.BindJSON(loan); err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	loan.ID = id

	if err := pkg.ValidateModelStruct(loan); err != nil {
		helper.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	if err := s.repo.CreateLoan(loan); err != nil {
		if strings.Contains(err.Error(), "book not found") {
			helper.HandleError(c, err, http.StatusNotFound)
			return
		}

		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": loan.ID,
	})
}

// GetAllLoans retrieves all loans from the loan service.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function retrieves all loans from the loan repository and returns them as JSON in the response body.
// If an error occurs during the process, it handles the error and returns an appropriate HTTP status code.
func (s *LoanService) GetAllLoans(c *gin.Context) {
	loans, err := s.repo.GetAllLoans()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, loans)
}

// ReturnLoan handles the return of a loan by its ID.
//
// Parameters:
// - c: a pointer to a gin.Context object representing the HTTP request context.
//
// Returns: None.
func (s *LoanService) ReturnLoan(c *gin.Context) {
	loanID := c.Param("loanId")

	if err := s.repo.ReturnLoan(loanID); err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
