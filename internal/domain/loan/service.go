package loan

import (
	"mybooks/internal/infrastructure/model"
	"mybooks/internal/infrastructure/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoanService struct {
	repo LoanRepository
}

func NewLoanService(repo LoanRepository) *LoanService {
	return &LoanService{
		repo: repo,
	}
}

func (s *LoanService) CreateLoan(c *gin.Context) {
	loan := new(model.Loan)

	id, err := utils.GenerateRandomID()
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	if err := c.BindJSON(loan); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	loan.ID = id

	if err := utils.ValidateStruct(loan); err != nil {
		utils.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	if err := s.repo.CreateLoan(loan); err != nil {
		if strings.Contains(err.Error(), "book not found") {
			utils.HandleError(c, err, http.StatusNotFound)
			return
		}

		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": loan.ID,
	})
}

func (s *LoanService) GetAllLoans(c *gin.Context) {
	loans, err := s.repo.GetAllLoans()
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, loans)
}

func (s *LoanService) ReturnLoan(c *gin.Context) {
	loanID := c.Param("loanId")

	if err := s.repo.ReturnLoan(loanID); err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
