package book

import (
	"mybooks/internal/infrastructure/model"
	"mybooks/internal/infrastructure/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
// It takes a gin.Context parameter which represents the HTTP request context.
// It returns an error if there was an issue binding the request body to a model.Book object.
// It returns an error if there was an issue creating the book in the BookRepository.
// It returns nil if the book was successfully created.
func (s *BookService) CreateBook(c *gin.Context) {
	book := new(model.Book)

	id, err := utils.GenerateRandomID()
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	book.ID = id

	if err := c.Bind(book); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	if err := utils.ValidateStruct(book); err != nil {
		utils.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	if err := s.repo.CreateBook(book); err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": book.ID,
	})
}

// GetAllBooks retrieves all books from the BookService.
//
// It takes a gin.Context parameter which represents the HTTP request context.
// It returns an error if there was an issue retrieving the books from the BookRepository.
// It returns a JSON response with the retrieved books if successful, or an error message if there was an issue.
func (s *BookService) GetAllBooks(c *gin.Context) {
	filters := make(map[string]interface{})

	if title := strings.TrimSpace(c.Query("title")); title != "" {
		filters["title"] = strings.ToLower(title)
	}
	if author := strings.TrimSpace(c.Query("author")); author != "" {
		filters["author"] = strings.ToLower(author)
	}
	if genre := strings.TrimSpace(c.Query("genre")); genre != "" {
		filters["genre"] = strings.ToLower(genre)
	}
	if isbn := strings.TrimSpace(c.Query("isbn")); isbn != "" {
		filters["isbn"] = isbn
	}
	if language := strings.TrimSpace(c.Query("language")); language != "" {
		filters["language"] = strings.ToLower(language)
	}
	if read := strings.TrimSpace(c.Query("read")); read != "" {
		readBool, err := strconv.ParseBool(read)
		if err != nil {
			utils.HandleError(c, err, http.StatusBadRequest)
			return
		}
		filters["read"] = readBool
	}

	books, err := s.repo.GetAllBooks(filters)
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBookById retrieves a book by its ID from the BookService.
//
// It takes a gin.Context parameter which represents the HTTP request context.
// The parameter "bookId" is extracted from the request path.
// It returns an error if there was an issue retrieving the book from the BookRepository.
// If the book is not found, it returns a JSON response with a status code of 404.
// If there was an issue retrieving the book, it returns a JSON response with a status code of 500.
// If the book is found, it returns a JSON response with the retrieved book and a status code of 200.
func (s *BookService) GetBookById(c *gin.Context) {
	id := c.Param("bookId")

	book, err := s.repo.GetBookById(id)
	if err != nil {
		if strings.Contains(err.Error(), "book not found") {
			utils.HandleError(c, err, http.StatusNotFound)
			return
		}
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book by its ID from the BookService.
//
// It takes a gin.Context parameter which represents the HTTP request context.
// The parameter "bookId" is extracted from the request path.
// It returns an error if there was an issue deleting the book from the BookRepository.
// If the book is not found, it returns a JSON response with a status code of 404.
// If there was an issue deleting the book, it returns a JSON response with a status code of 500.
// If the book is deleted successfully, it returns a JSON response with a status code of 200.
func (s *BookService) DeleteBook(c *gin.Context) {
	id := c.Param("bookId")

	if err := s.repo.DeleteBook(id); err != nil {
		if strings.Contains(err.Error(), "book not found") {
			utils.HandleError(c, err, http.StatusNotFound)
			return
		}
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// UpdateBook updates a book in the BookService.
//
// It takes a gin.Context parameter which represents the HTTP request context.
// The parameter "bookId" is extracted from the request path.
// It returns an error if there was an issue updating the book in the BookRepository.
// If the book is not found, it returns a JSON response with a status code of 404.
// If there was an issue updating the book, it returns a JSON response with a status code of 500.
// If the book is updated successfully, it returns a JSON response with a status code of 200.
func (s *BookService) UpdateBook(c *gin.Context) {
	id := c.Param("bookId")

	var book model.Book
	if err := c.Bind(&book); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	bookID, err := uuid.Parse(id)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	book.ID = bookID

	if err := s.repo.UpdateBook(&book); err != nil {
		if strings.Contains(err.Error(), "book not found") {
			utils.HandleError(c, err, http.StatusNotFound)
			return
		}
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}
