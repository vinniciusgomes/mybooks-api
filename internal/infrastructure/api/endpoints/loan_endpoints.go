package endpoints

import (
	"mybooks/internal/domain/loan"

	"github.com/gin-gonic/gin"
)

func Loan(r *gin.Engine, loanService *loan.LoanService) {
	r.POST("/v1/loans", loanService.CreateLoan)
	r.GET("/v1/loans", loanService.GetAllLoans)
	r.PUT("/v1/loans/:loanId/return", loanService.ReturnLoan)
}
