package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reading(r *gin.Engine) {
	r.PUT("/v1/books/:bookId/read", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT /books/:bookId/read")
	})
	r.DELETE("/v1/books/:bookId/read", func(c *gin.Context) {
		c.String(http.StatusOK, "DELETE /books/:bookId/read")
	})
}
