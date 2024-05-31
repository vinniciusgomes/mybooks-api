package endpoints

import (
	"mybooks/internal/domain/library"

	"github.com/labstack/echo/v4"
)

// Libraries registers the library-related API endpoints with the provided Echo instance.
//
// Parameters:
// - e: The Echo instance to register the endpoints with.
//
// Return type: None.
func Libraries(e *echo.Echo, libraryService *library.LibraryService) {
	e.GET("/v1/libraries", libraryService.GetAllLibraries)
	e.GET("/v1/libraries/:libraryId", libraryService.GetLibraryByID)
	e.POST("/v1/libraries", libraryService.CreateLibrary)
	e.PUT("/v1/libraries/:libraryId", libraryService.UpdateLibrary)
	e.DELETE("/v1/libraries/:libraryId", libraryService.DeleteLibrary)
	e.POST("/v1/libraries/:libraryId/books", libraryService.AddBookToLibrary)
}
