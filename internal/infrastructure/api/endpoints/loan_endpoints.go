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
			loansRouter.POST("/v1/loans", loanService.CreateLoan)
			loansRouter.GET("/v1/loans", loanService.GetAllLoans)
			loansRouter.PUT("/v1/loans/:loanId/return", loanService.ReturnLoan)
		}
	}
}
