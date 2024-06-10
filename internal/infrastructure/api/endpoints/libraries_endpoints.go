package endpoints

import (
	"mybooks/internal/domain/library"
	"mybooks/internal/infrastructure/api/middlewares"

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
	v1 := r.Group("/api/v1")
	{
		librariesRouter := v1.Group("/libraries")
		{
			librariesRouter.GET("/", middlewares.RequireAuth, libraryService.GetAllLibraries)
			librariesRouter.GET("/:libraryId", middlewares.RequireAuth, libraryService.GetLibraryByID)
			librariesRouter.POST("/", middlewares.RequireAuth, libraryService.CreateLibrary)
			librariesRouter.PUT("/:libraryId", middlewares.RequireAuth, libraryService.UpdateLibrary)
			librariesRouter.DELETE("/:libraryId", middlewares.RequireAuth, libraryService.DeleteLibrary)
			librariesRouter.POST("/:libraryId/books", middlewares.RequireAuth, libraryService.AddBookToLibrary)
		}
	}
}
