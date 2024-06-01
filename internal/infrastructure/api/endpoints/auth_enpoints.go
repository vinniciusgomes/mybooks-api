package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authentication registers the authentication-related API endpoints with the provided Gin engine.
//
// Parameters:
// - r: The Gin engine to register the endpoints with.
//
// Return type: None.
func Authentication(r *gin.Engine) {
	r.POST("/auth", func(c *gin.Context) {
		c.String(http.StatusOK, "POST /auth")
	})
	r.POST("/auth/google", func(c *gin.Context) {
		c.String(http.StatusOK, "POST /auth/google")
	})
}
