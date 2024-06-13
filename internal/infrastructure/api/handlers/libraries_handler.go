package handlers

import (
	"mybooks/internal/domain/services"
	"mybooks/internal/infrastructure/api/middlewares"

	"github.com/gin-gonic/gin"
)

// LibrariesHandler sets up the routes for the library handler in the provided gin.Engine.
//
// Parameters:
// - router: a pointer to a gin.Engine object representing the HTTP router.
// - libraryService: a pointer to a services.LibraryService object providing the library-related operations.
//
// Returns: None.
func LibrariesHandler(router *gin.Engine, libraryService *services.LibraryService) {
	v1 := router.Group("/v1")
	{
		librariesRouter := v1.Group("/libraries")
		{
			librariesRouter.GET("/", middlewares.AuthMiddleware(), libraryService.GetAllLibraries)
			librariesRouter.GET("/:libraryId", middlewares.AuthMiddleware(), libraryService.GetLibraryByID)
			librariesRouter.POST("/", middlewares.AuthMiddleware(), libraryService.CreateLibrary)
			librariesRouter.PUT("/:libraryId", middlewares.AuthMiddleware(), libraryService.UpdateLibrary)
			librariesRouter.DELETE("/:libraryId", middlewares.AuthMiddleware(), libraryService.DeleteLibrary)
			librariesRouter.POST("/:libraryId/books/:bookId", middlewares.AuthMiddleware(), libraryService.AddBookToLibrary)
			librariesRouter.DELETE("/:libraryId/books/:bookId", middlewares.AuthMiddleware(), libraryService.RemoveBookFromLibrary)
		}
	}
}
