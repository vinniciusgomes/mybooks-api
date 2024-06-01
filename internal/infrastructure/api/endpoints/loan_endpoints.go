package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Loan(r *gin.Engine) {
	r.PUT("/books/:bookId/loan", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT /books/:bookId/loan")
	})
	r.PUT("/books/:bookId/return", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT /books/:bookId/return")
	})
}
