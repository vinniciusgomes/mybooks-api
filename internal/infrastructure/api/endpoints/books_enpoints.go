package endpoints

import (
	"mybooks/internal/domain/book"
	"mybooks/internal/infrastructure/api/middlewares"

	"github.com/gin-gonic/gin"
)

// Books registers the book endpoints with the provided gin.Engine and book.BookService.
//
// Parameters:
// - router: a pointer to a gin.Engine object representing the HTTP router.
// - bookService: a pointer to a book.BookService object providing the book-related operations.
//
// Returns: None.
func Books(router *gin.Engine, bookService *book.BookService) {
	v1 := router.Group("/api/v1")
	{
		booksRouter := v1.Group("/books")
		{
			booksRouter.GET("/", middlewares.RequireAuth, bookService.GetAllBooks)
			booksRouter.GET("/:bookId", middlewares.RequireAuth, bookService.GetBookById)
			booksRouter.POST("", middlewares.RequireAuth, bookService.CreateBook)
			booksRouter.PUT("/:bookId", middlewares.RequireAuth, bookService.UpdateBook)
			booksRouter.DELETE("/:bookId", middlewares.RequireAuth, bookService.DeleteBook)
		}
	}
}
