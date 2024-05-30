package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Billing registers the billing-related API endpoints with the provided Echo instance.
//
// Parameters:
// - e: The Echo instance to register the endpoints with.
//
// Return type: None.
func Billing(e *echo.Echo) {
	e.GET("/billing", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET /billing")
	})
	e.POST("/subscribe", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST /subscribe")
	})
}
