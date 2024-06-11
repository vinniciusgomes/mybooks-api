package endpoints

import (
	"mybooks/internal/domain/loan"
	"mybooks/internal/infrastructure/api/middlewares"

	"github.com/gin-gonic/gin"
)

// Loan sets up the loan endpoints on the given gin router.
//
// Parameters:
// - router: a pointer to a gin.Engine representing the router.
// - loanService: a pointer to a loan.LoanService representing the loan service.
//
// Return type: None.
func Loan(router *gin.Engine, loanService *loan.LoanService) {
	v1 := router.Group("/v1")
	{
		loansRouter := v1.Group("/loans")
		{
			loansRouter.POST("/", middlewares.JWTAuthMiddleware(), loanService.CreateLoan)
			loansRouter.GET("/", middlewares.JWTAuthMiddleware(), loanService.GetAllLoans)
			loansRouter.PUT("/:loanId/return", middlewares.JWTAuthMiddleware(), loanService.ReturnLoan)
		}
	}
}
