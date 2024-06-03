package endpoints

import (
	"mybooks/internal/domain/library"

	"github.com/gin-gonic/gin"
)

// Libraries sets up the routes for the library endpoints in the provided gin.Engine.
//
// Parameters:
// - r: a pointer to a gin.Engine object representing the HTTP router.
// - libraryService: a pointer to a library.LibraryService object providing the library-related operations.
//
// Returns: None.
func Libraries(r *gin.Engine, libraryService *library.LibraryService) {
	r.GET("/v1/libraries", libraryService.GetAllLibraries)
	r.GET("/v1/libraries/:libraryId", libraryService.GetLibraryByID)
	r.POST("/v1/libraries", libraryService.CreateLibrary)
	r.PUT("/v1/libraries/:libraryId", libraryService.UpdateLibrary)
	r.DELETE("/v1/libraries/:libraryId", libraryService.DeleteLibrary)
	r.POST("/v1/libraries/:libraryId/books", libraryService.AddBookToLibrary)
}
