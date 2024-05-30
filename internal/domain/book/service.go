package book

import (
	"mybooks/internal/infrastructure/model"
	"mybooks/internal/infrastructure/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type BookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

// CreateBook creates a new book in the BookService.
//
// It takes an echo.Context parameter which represents the HTTP request context.
// It returns an error if there was an issue binding the request body to a model.Book object.
// It returns an error if there was an issue creating the book in the BookRepository.
// It returns nil if the book was successfully created.
func (s *BookService) CreateBook(c echo.Context) error {
	book := new(model.Book)

	id, err := utils.GenerateRandomID()
	if err != nil {
		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	book.ID = id

	if err := c.Bind(book); err != nil {
		return utils.HandleError(c, err, http.StatusBadRequest)
	}

	if err := utils.ValidateStruct(book); err != nil {
		return utils.HandleError(c, err, http.StatusUnprocessableEntity)
	}

	if err := s.repo.CreateBook(book); err != nil {
		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"id": book.ID,
	}
	println(data)
	return c.NoContent(http.StatusCreated)
}

// GetAllBooks retrieves all books from the BookService.
//
// It takes an echo.Context parameter which represents the HTTP request context.
// It returns an error if there was an issue retrieving the books from the BookRepository.
// It returns a JSON response with the retrieved books if successful, or an error message if there was an issue.
func (s *BookService) GetAllBooks(c echo.Context) error {
	filters := make(map[string]interface{})

	if title := strings.TrimSpace(c.QueryParam("title")); title != "" {
		filters["title"] = title
	} else if author := strings.TrimSpace(c.QueryParam("author")); author != "" {
		filters["author"] = author
	} else if genre := strings.TrimSpace(c.QueryParam("genre")); genre != "" {
		filters["genre"] = genre
	} else if isbn := strings.TrimSpace(c.QueryParam("isbn")); isbn != "" {
		filters["isbn"] = isbn
	} else if language := strings.TrimSpace(c.QueryParam("language")); language != "" {
		filters["language"] = language
	} else if read := strings.TrimSpace(c.QueryParam("read")); read != "" {
		readBool, err := strconv.ParseBool(read)
		if err != nil {
			return utils.HandleError(c, err, http.StatusBadRequest)
		}
		filters["read"] = readBool
	}

	books, err := s.repo.GetAllBooks(filters)
	if err != nil {
		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, books)
}

// GetBookById retrieves a book by its ID from the BookService.
//
// It takes an echo.Context parameter which represents the HTTP request context.
// The parameter "bookId" is extracted from the request path.
// It returns an error if there was an issue retrieving the book from the BookRepository.
// If the book is not found, it returns a JSON response with a status code of 404.
// If there was an issue retrieving the book, it returns a JSON response with a status code of 500.
// If the book is found, it returns a JSON response with the retrieved book and a status code of 200.
func (s *BookService) GetBookById(c echo.Context) error {
	id := c.Param("bookId")

	book, err := s.repo.GetBookById(id)
	if err != nil {
		if strings.Contains(err.Error(), "book not found") {
			return utils.HandleError(c, err, http.StatusNotFound)
		}

		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book by its ID from the BookService.
//
// It takes an echo.Context parameter which represents the HTTP request context.
// The parameter "bookId" is extracted from the request path.
// It returns an error if there was an issue deleting the book from the BookRepository.
// If the book is not found, it returns a JSON response with a status code of 404.
// If there was an issue deleting the book, it returns a JSON response with a status code of 500.
// If the book is deleted successfully, it returns a JSON response with a status code of 200.
func (s *BookService) DeleteBook(c echo.Context) error {
	id := c.Param("bookId")

	if err := s.repo.DeleteBook(id); err != nil {
		if strings.Contains(err.Error(), "book not found") {
			return utils.HandleError(c, err, http.StatusNotFound)
		}

		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// UpdateBook updates a book in the BookService.
//
// It takes an echo.Context parameter which represents the HTTP request context.
// The parameter "bookId" is extracted from the request path.
// It returns an error if there was an issue updating the book in the BookRepository.
// If the book is not found, it returns a JSON response with a status code of 404.
// If there was an issue updating the book, it returns a JSON response with a status code of 500.
// If the book is updated successfully, it returns a JSON response with a status code of 200.
func (s *BookService) UpdateBook(c echo.Context) error {
	id := c.Param("bookId")

	var book model.Book
	if err := c.Bind(&book); err != nil {
		return utils.HandleError(c, err, http.StatusBadRequest)
	}

	book.ID = id

	if err := s.repo.UpdateBook(&book); err != nil {
		if strings.Contains(err.Error(), "book not found") {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
