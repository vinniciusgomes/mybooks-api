package endpoints

import (
	"mybooks/internal/domain/book"

	"github.com/gin-gonic/gin"
)

// Books registers the book-related API endpoints with the provided Gin engine and book service.
//
// Parameters:
// - r: The Gin engine to register the endpoints with.
// - bookService: The book service to handle the book-related requests.
//
// Return type: None.
func Books(r *gin.Engine, bookService *book.BookService) {
	r.GET("/v1/books", bookService.GetAllBooks)
	r.GET("/v1/books/:bookId", func(c *gin.Context) {
		bookService.GetBookById(c)
	})
	r.POST("/v1/books", func(c *gin.Context) {
		bookService.CreateBook(c)
	})
	r.PUT("/v1/books/:bookId", func(c *gin.Context) {
		bookService.UpdateBook(c)
	})
	r.DELETE("/v1/books/:bookId", func(c *gin.Context) {
		bookService.DeleteBook(c)
	})
}
