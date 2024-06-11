package endpoints

import (
	"mybooks/internal/domain/library"
	"mybooks/internal/infrastructure/api/middlewares"

	"github.com/gin-gonic/gin"
)

// Libraries sets up the routes for the library endpoints in the provided gin.Engine.
//
// Parameters:
// - router: a pointer to a gin.Engine object representing the HTTP router.
// - libraryService: a pointer to a library.LibraryService object providing the library-related operations.
//
// Returns: None.
func Libraries(router *gin.Engine, libraryService *library.LibraryService) {
	v1 := router.Group("/v1")
	{
		librariesRouter := v1.Group("/libraries")
		{
			librariesRouter.GET("/", middlewares.JWTAuthMiddleware(), libraryService.GetAllLibraries)
			librariesRouter.GET("/:libraryId", middlewares.JWTAuthMiddleware(), libraryService.GetLibraryByID)
			librariesRouter.POST("/", middlewares.JWTAuthMiddleware(), libraryService.CreateLibrary)
			librariesRouter.PUT("/:libraryId", middlewares.JWTAuthMiddleware(), libraryService.UpdateLibrary)
			librariesRouter.DELETE("/:libraryId", middlewares.JWTAuthMiddleware(), libraryService.DeleteLibrary)
			librariesRouter.POST("/:libraryId/books", middlewares.JWTAuthMiddleware(), libraryService.AddBookToLibrary)
		}
	}
}
