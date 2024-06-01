package endpoints

import (
	"mybooks/internal/domain/book"

	"github.com/gin-gonic/gin"
)

func Books(r *gin.Engine, bookService *book.BookService) {
	r.GET("/v1/books", bookService.GetAllBooks)
	r.GET("/v1/books/:bookId", bookService.GetBookById)
	r.POST("/v1/books", bookService.CreateBook)
	r.PUT("/v1/books/:bookId", bookService.UpdateBook)
	r.DELETE("/v1/books/:bookId", bookService.DeleteBook)
}
