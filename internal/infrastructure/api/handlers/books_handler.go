package handlers

import (
	"mybooks/internal/domain/services"
	"mybooks/internal/infrastructure/api/middlewares"

	"github.com/gin-gonic/gin"
)

// BooksHandler registers the book handler with the provided gin.Engine and services.BookService.
//
// Parameters:
// - router: a pointer to a gin.Engine object representing the HTTP router.
// - bookService: a pointer to a services.BookService object providing the book-related operations.
//
// Returns: None.
func BooksHandler(router *gin.Engine, bookService *services.BookService) {
	v1 := router.Group("/v1")
	{
		booksRouter := v1.Group("/books")
		{
			booksRouter.GET("/", middlewares.AuthMiddleware(), bookService.GetAllBooks)
			booksRouter.GET("/:bookId", middlewares.AuthMiddleware(), bookService.GetBookById)
			booksRouter.POST("", middlewares.AuthMiddleware(), bookService.CreateBook)
			booksRouter.PUT("/:bookId", middlewares.AuthMiddleware(), bookService.UpdateBook)
			booksRouter.DELETE("/:bookId", middlewares.AuthMiddleware(), bookService.DeleteBook)
		}
	}
}
