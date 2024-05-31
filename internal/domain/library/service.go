package library

import (
	"mybooks/internal/infrastructure/model"
	"mybooks/internal/infrastructure/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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
// - c: The echo.Context object representing the HTTP request and response.
//
// Returns:
// - error: An error if there was a problem creating the library, otherwise nil.
func (s *LibraryService) CreateLibrary(c echo.Context) error {
	library := new(model.Library)

	id, err := utils.GenerateRandomID()
	if err != nil {
		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	library.ID = id

	if err := c.Bind(library); err != nil {
		return utils.HandleError(c, err, http.StatusBadRequest)
	}

	if err := utils.ValidateStruct(library); err != nil {
		return utils.HandleError(c, err, http.StatusUnprocessableEntity)
	}

	if err := s.repo.CreateLibrary(library); err != nil {
		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"id": library.ID,
	}

	return c.JSON(http.StatusCreated, data)
}

// GetAllLibraries retrieves all libraries from the library service and returns them as a JSON response.
//
// Parameters:
// - c: The echo.Context object representing the HTTP request and response.
//
// Returns:
// - error: An error if there was a problem retrieving the libraries, otherwise nil.
func (s *LibraryService) GetAllLibraries(c echo.Context) error {
	libraries, err := s.repo.GetAllLibraries()
	if err != nil {
		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, libraries)
}

// GetLibraryByID retrieves a library by its ID from the LibraryService.
//
// Parameters:
// - c: The echo.Context object representing the HTTP request and response.
//
// Returns:
// - error: An error if there was a problem retrieving the library, otherwise nil.
func (s *LibraryService) GetLibraryByID(c echo.Context) error {
	libraryID := c.Param("libraryId")

	library, err := s.repo.GetLibraryByID(libraryID)
	if err != nil {
		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, library)
}

// DeleteLibrary deletes a library from the LibraryService.
//
// Parameters:
// - c: The echo.Context object representing the HTTP request and response.
//
// Returns:
// - error: An error if there was a problem deleting the library, otherwise nil.
func (s *LibraryService) DeleteLibrary(c echo.Context) error {
	libraryID := c.Param("libraryId")

	if err := s.repo.DeleteLibrary(libraryID); err != nil {
		if strings.Contains(err.Error(), "library not found") {
			return utils.HandleError(c, err, http.StatusNotFound)
		}

		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// UpdateLibrary updates a library in the LibraryService.
//
// Parameters:
// - c: The echo.Context object representing the HTTP request and response.
//
// Returns:
// - error: An error if there was a problem updating the library, otherwise nil.
func (s *LibraryService) UpdateLibrary(c echo.Context) error {
	id := c.Param("libraryId")

	var library model.Library
	if err := c.Bind(&library); err != nil {
		return utils.HandleError(c, err, http.StatusBadRequest)
	}

	libraryID, err := uuid.Parse(id)
	if err != nil {
		return utils.HandleError(c, err, http.StatusBadRequest)
	}

	library.ID = libraryID

	if err := utils.ValidateStruct(library); err != nil {
		return utils.HandleError(c, err, http.StatusUnprocessableEntity)
	}

	if err := s.repo.UpdateLibrary(&library); err != nil {
		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// AddBookToLibrary adds a book to a library in the LibraryService.
//
// Parameters:
// - c: The echo.Context object representing the HTTP request and response.
//
// Returns:
// - error: An error if there was a problem adding the book to the library, otherwise nil.
func (s *LibraryService) AddBookToLibrary(c echo.Context) error {
	libraryID := c.Param("libraryId")

	var req AddBookRequest

	if err := c.Bind(&req); err != nil {
		return utils.HandleError(c, err, http.StatusBadRequest)
	}

	bookID := req.BookID

	if err := s.repo.AddBookToLibrary(libraryID, bookID); err != nil {
		return utils.HandleError(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
