package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication(r *gin.Engine) {
	r.POST("/v1/auth", func(c *gin.Context) {
		c.String(http.StatusOK, "POST /auth")
	})
	r.POST("/v1/auth/google", func(c *gin.Context) {
		c.String(http.StatusOK, "POST /auth/google")
	})
}
