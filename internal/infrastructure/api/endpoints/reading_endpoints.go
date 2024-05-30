package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Reading registers the reading-related API endpoints with the provided Echo instance.
//
// Parameters:
// - e: The Echo instance to register the endpoints with.
//
// Return type: None.
func Reading(e *echo.Echo) {
	e.PUT("/books/:bookId/read", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /books/:bookId/read")
	})
	e.DELETE("/books/:bookId/read", func(c echo.Context) error {
		return c.String(http.StatusOK, "DELETE /books/:bookId/read")
	})
}
