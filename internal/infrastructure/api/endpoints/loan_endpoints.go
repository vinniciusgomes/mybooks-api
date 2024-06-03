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
	r.POST("/v1/loans", loanService.CreateLoan)
	r.GET("/v1/loans", loanService.GetAllLoans)
	r.PUT("/v1/loans/:loanId/return", loanService.ReturnLoan)
}
