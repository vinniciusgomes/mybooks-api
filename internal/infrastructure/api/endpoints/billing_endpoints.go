package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Billing(r *gin.Engine) {
	r.GET("/billing", func(c *gin.Context) {
		c.String(http.StatusOK, "GET /billing")
	})
	r.POST("/subscribe", func(c *gin.Context) {
		c.String(http.StatusOK, "POST /subscribe")
	})
}
