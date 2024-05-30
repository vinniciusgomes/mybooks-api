package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Profile registers the authentication-related API endpoints with the provided Echo instance.
//
// Parameters:
// - e: The Echo instance to register the endpoints with.
//
// Return type: None.
func Profile(e *echo.Echo) {
	e.PUT("/profile/photo", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /profile/photo")
	})
	e.PUT("/profile", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /profile")
	})
	e.DELETE("/profile", func(c echo.Context) error {
		return c.String(http.StatusOK, "DELETE /profile")
	})
}
