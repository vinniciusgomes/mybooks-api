package handlers

import (
	"mybooks/internal/domain/services"
	"mybooks/internal/infrastructure/api/middlewares"

	"github.com/gin-gonic/gin"
)

// LoanHandler sets up the loan handler on the given gin router.
//
// Parameters:
// - router: a pointer to a gin.Engine representing the router.
// - loanService: a pointer to a services.LoanService representing the loan service.
//
// Return type: None.
func LoanHandler(router *gin.Engine, loanService *services.LoanService) {
	v1 := router.Group("/v1")
	{
		loansRouter := v1.Group("/loans")
		{
			loansRouter.POST("/", middlewares.AuthMiddleware(), loanService.CreateLoan)
			loansRouter.GET("/", middlewares.AuthMiddleware(), loanService.GetAllLoans)
			loansRouter.PUT("/:loanId/return", middlewares.AuthMiddleware(), loanService.ReturnLoan)
		}
	}
}
