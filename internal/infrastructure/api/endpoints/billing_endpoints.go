package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Billing(r *gin.Engine) {
	r.GET("/v1/billing", func(c *gin.Context) {
		c.String(http.StatusOK, "GET /billing")
	})
	r.POST("/v1/subscribe", func(c *gin.Context) {
		c.String(http.StatusOK, "POST /subscribe")
	})
}
