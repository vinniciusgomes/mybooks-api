package endpoints

import (
	"mybooks/internal/domain/library"

	"github.com/gin-gonic/gin"
)

func Libraries(r *gin.Engine, libraryService *library.LibraryService) {
	r.GET("/v1/libraries", libraryService.GetAllLibraries)
	r.GET("/v1/libraries/:libraryId", libraryService.GetLibraryByID)
	r.POST("/v1/libraries", libraryService.CreateLibrary)
	r.PUT("/v1/libraries/:libraryId", libraryService.UpdateLibrary)
	r.DELETE("/v1/libraries/:libraryId", libraryService.DeleteLibrary)
	r.POST("/v1/libraries/:libraryId/books", libraryService.AddBookToLibrary)
}
