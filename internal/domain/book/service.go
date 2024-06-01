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
