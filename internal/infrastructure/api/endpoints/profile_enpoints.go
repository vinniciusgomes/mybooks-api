package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Profile registers the authentication-related API endpoints with the provided Echo instance.
//
// Parameters:
// - e: The Echo instance to register the endpoints with.
//
// Return type: None.
func Profile(r *gin.Engine) {
	r.PUT("/profile/photo", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT /profile/photo")
	})
	r.PUT("/profile", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT /profile")
	})
	r.DELETE("/profile", func(c *gin.Context) {
		c.String(http.StatusOK, "DELETE /profile")
	})
}
