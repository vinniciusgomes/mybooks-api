package book

import (
	"mybooks/internal/infrastructure/model"
	"mybooks/internal/infrastructure/utils"
	"net/http"

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
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	book.ID = id

	if err := c.Bind(book); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusBadRequest, data)
	}

	if err := utils.ValidateStruct(book); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusBadRequest, data)
	}

	if err := s.repo.CreateBook(book); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
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
	books, err := s.repo.GetAllBooks()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, books)
}
