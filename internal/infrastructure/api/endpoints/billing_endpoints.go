package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Billing registers the billing-related API endpoints with the provided Gin engine.
//
// Parameters:
// - r: The Gin engine to register the endpoints with.
//
// Return type: None.
func Billing(r *gin.Engine) {
	r.GET("/billing", func(c *gin.Context) {
		c.String(http.StatusOK, "GET /billing")
	})
	r.POST("/subscribe", func(c *gin.Context) {
		c.String(http.StatusOK, "POST /subscribe")
	})
}
