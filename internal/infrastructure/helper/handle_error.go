package helper

import (
	"log"

	"github.com/gin-gonic/gin"
)

// HandleError handles the error by logging it and returning a JSON response with the error message.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request context.
// - err: The error to be handled.
// - statusCode: The HTTP status code to be returned in the JSON response.
//
// Returns:
// - error: None.
func HandleError(c *gin.Context, err error, statusCode int) {
	log.Printf("Error: %s", err.Error())
	data := map[string]interface{}{
		"message": err.Error(),
	}

	c.JSON(statusCode, data)
}
