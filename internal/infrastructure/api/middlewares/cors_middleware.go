package middlewares

import "github.com/gin-gonic/gin"

// CORSMiddleware is a middleware function that enables Cross-Origin Resource Sharing (CORS) for a Gin application.
//
// It sets the necessary headers to allow cross-origin requests. It allows all origins (*), allows credentials,
// sets the allowed headers, and sets the allowed methods (POST, OPTIONS, GET, PUT). If the request method is OPTIONS,
// it aborts the request with a 204 status code. Otherwise, it calls the next middleware or handler.
//
// Returns a Gin handler function.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
