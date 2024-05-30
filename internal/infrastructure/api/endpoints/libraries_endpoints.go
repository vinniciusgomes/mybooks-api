package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Libraries registers the library-related API endpoints with the provided Echo instance.
//
// Parameters:
// - e: The Echo instance to register the endpoints with.
//
// Return type: None.
func Libraries(e *echo.Echo) {
	e.GET("/libraries", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET /libraries")
	})
	e.POST("/libraries", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST /libraries")
	})
	e.PUT("/libraries/:libraryId", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /libraries/:libraryId")
	})
	e.DELETE("/libraries/:libraryId", func(c echo.Context) error {
		return c.String(http.StatusOK, "DELETE /libraries/:libraryId")
	})
	e.POST("/libraries/:libraryId/books", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST /libraries/:libraryId/books")
	})
}
