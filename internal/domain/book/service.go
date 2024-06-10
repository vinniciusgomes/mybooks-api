package book

import (
	"mybooks/internal/infrastructure/helper"
	"mybooks/internal/infrastructure/model"
	"mybooks/pkg"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookService struct {
	repo BookRepository
}

// NewBookService creates a new instance of the BookService struct.
//
// It takes a BookRepository as a parameter and returns a pointer to a BookService.
func NewBookService(repo BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

// CreateBook creates a new book in the BookService.
//
// It takes a gin.Context as a parameter and returns nothing.
// The function generates a random ID, binds the JSON from the request to a model.Book struct,
// validates the struct, creates the book in the repository, and returns the ID of the created book.
// If any error occurs during the process, it handles the error and returns an appropriate HTTP status code.
func (s *BookService) CreateBook(c *gin.Context) {
	book := new(model.Book)

	user, err := helper.GetUserFromContext(c)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
		return
	}

	id, err := pkg.GenerateRandomID()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	if err := c.BindJSON(book); err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	book.ID = id
	book.UserID = user.ID
	book.User = *user

	if err := pkg.ValidateModelStruct(book); err != nil {
		helper.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	if err := s.repo.CreateBook(book); err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": book.ID,
	})
}

// GetAllBooks retrieves all books from the BookService that match the provided filters.
//
// Parameters:
// - c: a pointer to a gin.Context object representing the HTTP request and response.
//
// Returns:
// - None.
func (s *BookService) GetAllBooks(c *gin.Context) {
	user, err := helper.GetUserFromContext(c)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

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
			helper.HandleError(c, err, http.StatusBadRequest)
			return
		}
		filters["read"] = readBool
	}

	books, err := s.repo.GetAllBooks(userID.String(), filters)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBookById retrieves a book by its ID from the BookService.
//
// Parameters:
// - c: a pointer to a gin.Context object representing the HTTP request and response.
//
// Returns:
// - None.
func (s *BookService) GetBookById(c *gin.Context) {
	id := c.Param("bookId")

	user, err := helper.GetUserFromContext(c)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	book, err := s.repo.GetBookById(userID.String(), id)
	if err != nil {
		if strings.Contains(err.Error(), "book not found") {
			helper.HandleError(c, err, http.StatusNotFound)
			return
		}
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book by its ID.
//
// Parameters:
// - c: a pointer to a gin.Context object representing the HTTP request and response.
//
// Returns:
// - None.
func (s *BookService) DeleteBook(c *gin.Context) {
	id := c.Param("bookId")

	user, err := helper.GetUserFromContext(c)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	if err := s.repo.DeleteBook(userID.String(), id); err != nil {
		if strings.Contains(err.Error(), "book not found") {
			helper.HandleError(c, err, http.StatusNotFound)
			return
		}

		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// UpdateBook updates a book in the BookService.
//
// It takes a gin.Context as a parameter and returns nothing.
// The function retrieves the book ID from the request parameters, binds the JSON from the request to a model.Book struct,
// parses the ID, sets the ID on the book, updates the book in the repository, and returns a success message.
// If any error occurs during the process, it handles the error and returns an appropriate HTTP status code.
func (s *BookService) UpdateBook(c *gin.Context) {
	id := c.Param("bookId")

	user, err := helper.GetUserFromContext(c)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
		return
	}

	userID := user.ID

	var book model.Book
	if err := c.BindJSON(&book); err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	bookID, err := uuid.Parse(id)
	if err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	book.ID = bookID

	if err := s.repo.UpdateBook(userID.String(), &book); err != nil {
		if strings.Contains(err.Error(), "book not found") {
			helper.HandleError(c, err, http.StatusNotFound)
			return
		}
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}
