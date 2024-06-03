package endpoints

import (
	"mybooks/internal/domain/book"

	"github.com/gin-gonic/gin"
)

// Books registers the book endpoints with the provided gin.Engine and book.BookService.
//
// Parameters:
// - r: a pointer to a gin.Engine object representing the HTTP router.
// - bookService: a pointer to a book.BookService object providing the book-related operations.
//
// Returns: None.
func Books(r *gin.Engine, bookService *book.BookService) {
	r.GET("/v1/books", bookService.GetAllBooks)
	r.GET("/v1/books/:bookId", bookService.GetBookById)
	r.POST("/v1/books", bookService.CreateBook)
	r.PUT("/v1/books/:bookId", bookService.UpdateBook)
	r.DELETE("/v1/books/:bookId", bookService.DeleteBook)
}
