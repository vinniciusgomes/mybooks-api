package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication(r *gin.Engine) {
	r.POST("/auth", func(c *gin.Context) {
		c.String(http.StatusOK, "POST /auth")
	})
	r.POST("/auth/google", func(c *gin.Context) {
		c.String(http.StatusOK, "POST /auth/google")
	})
}
