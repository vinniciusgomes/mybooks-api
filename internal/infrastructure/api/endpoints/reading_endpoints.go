package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Reading registers the reading-related API endpoints with the provided Echo instance.
//
// Parameters:
// - e: The Echo instance to register the endpoints with.
//
// Return type: None.
func Reading(r *gin.Engine) {
	r.PUT("/books/:bookId/read", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT /books/:bookId/read")
	})
	r.DELETE("/books/:bookId/read", func(c *gin.Context) {
		c.String(http.StatusOK, "DELETE /books/:bookId/read")
	})
}
