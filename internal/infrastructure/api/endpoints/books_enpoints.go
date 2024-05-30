package endpoints

import (
	"mybooks/internal/domain/book"

	"github.com/labstack/echo/v4"
)

// Books registers the book-related API endpoints with the provided Echo instance and book service.
//
// Parameters:
// - e: The Echo instance to register the endpoints with.
// - bookService: The book service to handle the book-related requests.
//
// Return type: None.
func Books(e *echo.Echo, bookService *book.BookService) {
	e.GET("/v1/books", bookService.GetAllBooks)
	e.GET("/v1/books/:bookId", bookService.GetBookById)
	e.POST("/v1/books", bookService.CreateBook)
	e.PUT("/v1/books/:bookId", bookService.UpdateBook)
	e.DELETE("/v1/books/:bookId", bookService.DeleteBook)
}
