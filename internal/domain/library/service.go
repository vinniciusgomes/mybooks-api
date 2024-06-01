package library

import (
	"mybooks/internal/infrastructure/model"
	"mybooks/internal/infrastructure/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LibraryService struct {
	repo LibraryRepository
}

type AddBookRequest struct {
	BookID string `json:"book_id"`
}

func NewLibraryService(repo LibraryRepository) *LibraryService {
	return &LibraryService{
		repo: repo,
	}
}

// CreateLibrary creates a new library in the system.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
func (s *LibraryService) CreateLibrary(c *gin.Context) {
	library := new(model.Library)

	id, err := utils.GenerateRandomID()
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	library.ID = id

	if err := c.Bind(library); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	if err := utils.ValidateStruct(library); err != nil {
		utils.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	if err := s.repo.CreateLibrary(library); err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"id": library.ID,
	}

	c.JSON(http.StatusCreated, data)
}

// GetAllLibraries retrieves all libraries from the library service and returns them as a JSON response.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
func (s *LibraryService) GetAllLibraries(c *gin.Context) {
	libraries, err := s.repo.GetAllLibraries()
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, libraries)
}

// GetLibraryByID retrieves a library by its ID from the LibraryService.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
func (s *LibraryService) GetLibraryByID(c *gin.Context) {
	libraryID := c.Param("libraryId")

	library, err := s.repo.GetLibraryByID(libraryID)
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, library)
}

// DeleteLibrary deletes a library from the LibraryService.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
func (s *LibraryService) DeleteLibrary(c *gin.Context) {
	libraryID := c.Param("libraryId")

	if err := s.repo.DeleteLibrary(libraryID); err != nil {
		if strings.Contains(err.Error(), "library not found") {
			utils.HandleError(c, err, http.StatusNotFound)
			return
		}

		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// UpdateLibrary updates a library in the LibraryService.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
func (s *LibraryService) UpdateLibrary(c *gin.Context) {
	id := c.Param("libraryId")

	var library model.Library
	if err := c.Bind(&library); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	libraryID, err := uuid.Parse(id)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	library.ID = libraryID

	if err := utils.ValidateStruct(library); err != nil {
		utils.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	if err := s.repo.UpdateLibrary(&library); err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// AddBookToLibrary adds a book to a library in the LibraryService.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
func (s *LibraryService) AddBookToLibrary(c *gin.Context) {
	libraryID := c.Param("libraryId")

	var req AddBookRequest

	if err := c.Bind(&req); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	bookID := req.BookID

	if err := s.repo.AddBookToLibrary(libraryID, bookID); err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
