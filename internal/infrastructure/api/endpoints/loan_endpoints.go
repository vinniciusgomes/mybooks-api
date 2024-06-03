package endpoints

import (
	"mybooks/internal/domain/loan"

	"github.com/gin-gonic/gin"
)

// Loan sets up the loan endpoints on the given gin router.
//
// Parameters:
// - r: a pointer to a gin.Engine representing the router.
// - loanService: a pointer to a loan.LoanService representing the loan service.
//
// Return type: None.
func Loan(r *gin.Engine, loanService *loan.LoanService) {
	v1 := r.Group("/api/v1")
	{
		loansRouter := v1.Group("/loans")
		{
			loansRouter.POST("/", loanService.CreateLoan)
			loansRouter.GET("/", loanService.GetAllLoans)
			loansRouter.PUT("/:loanId/return", loanService.ReturnLoan)
		}
	}
}
