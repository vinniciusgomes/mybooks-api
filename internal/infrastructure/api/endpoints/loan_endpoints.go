package endpoints

import (
	"mybooks/internal/domain/loan"
	"mybooks/internal/infrastructure/api/middlewares"

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
			loansRouter.POST("/", middlewares.RequireAuth, loanService.CreateLoan)
			loansRouter.GET("/", middlewares.RequireAuth, loanService.GetAllLoans)
			loansRouter.PUT("/:loanId/return", middlewares.RequireAuth, loanService.ReturnLoan)
		}
	}
}
