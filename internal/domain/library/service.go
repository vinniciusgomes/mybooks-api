package library

import (
	"mybooks/internal/infrastructure/model"
	"mybooks/internal/infrastructure/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LibraryService struct {
	repo LibraryRepository
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
