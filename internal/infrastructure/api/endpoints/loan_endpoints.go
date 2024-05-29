package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Loan registers the loan-related API endpoints with the provided Echo instance.
//
// Parameters:
// - e: The Echo instance to register the endpoints with.
//
// Return type: None.
func Loan(e *echo.Echo) {
	e.PUT("/books/:bookId/loan", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /books/:bookId/loan")
	})
	e.PUT("/books/:bookId/return", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /books/:bookId/return")
	})
}
