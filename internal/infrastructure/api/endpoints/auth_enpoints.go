package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Authentication registers the authentication-related API endpoints with the provided Echo instance.
//
// Parameters:
// - e: The Echo instance to register the endpoints with.
//
// Return type: None.
func Authentication(e *echo.Echo) {
	e.POST("/auth", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST /auth")
	})
	e.POST("/auth/google", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST /auth/google")
	})
}
