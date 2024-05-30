package utils

import (
	"log"

	"github.com/labstack/echo/v4"
)

// HandleError handles the error by logging it and returning a JSON response with the error message.
//
// Parameters:
// - c: The echo.Context object representing the HTTP request context.
// - err: The error to be handled.
// - statusCode: The HTTP status code to be returned in the JSON response.
//
// Returns:
// - error: An error if there was a problem generating the JSON response.
func HandleError(c echo.Context, err error, statusCode int) error {
	log.Printf("Error: %s", err.Error())
	data := map[string]interface{}{
		"message": err.Error(),
	}

	return c.JSON(statusCode, data)
}
