package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
