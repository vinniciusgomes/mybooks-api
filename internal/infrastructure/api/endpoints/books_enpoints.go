package endpoints

import (
	"mybooks/internal/domain/book"

	"github.com/gin-gonic/gin"
)

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
