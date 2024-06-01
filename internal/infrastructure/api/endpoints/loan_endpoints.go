package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Loan registers the loan-related API endpoints with the provided Echo instance.
//
// Parameters:
// - e: The Echo instance to register the endpoints with.
//
// Return type: None.
func Loan(r *gin.Engine) {
	r.PUT("/books/:bookId/loan", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT /books/:bookId/loan")
	})
	r.PUT("/books/:bookId/return", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT /books/:bookId/return")
	})
}
