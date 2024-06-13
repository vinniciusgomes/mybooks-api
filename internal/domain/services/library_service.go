package services

import (
	"mybooks/internal/domain/models"
	"mybooks/internal/domain/repositories"
	"mybooks/internal/infrastructure/helpers"
	"mybooks/pkg"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LibraryService struct {
	repo repositories.LibraryRepository
}

type LibraryResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewLibraryService creates a new instance of LibraryService.
//
// Parameters:
// - repo: The LibraryRepository implementation used by the service.
//
// Returns:
// - *LibraryService: A pointer to the newly created LibraryService instance.
func NewLibraryService(repo repositories.LibraryRepository) *LibraryService {
	return &LibraryService{
		repo: repo,
	}
}

// CreateLibrary creates a new library in the database.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function generates a random ID, binds the JSON request body to a Library struct,
// validates the struct, creates the library in the repository, and returns the ID of the created library.
func (s *LibraryService) CreateLibrary(c *gin.Context) {
	library := new(models.Library)

	user, err := helpers.GetUserFromContext(c)
	if err != nil {
		helpers.HandleError(c, err, http.StatusUnauthorized)
		return
	}

	id, err := pkg.GenerateRandomID()
	if err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	if err := c.BindJSON(library); err != nil {
		helpers.HandleError(c, err, http.StatusBadRequest)
		return
	}

	library.ID = id
	library.UserID = user.ID
	library.User = *user

	if err := pkg.ValidateModelStruct(library); err != nil {
		helpers.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	if err := s.repo.CreateLibrary(library); err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
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
// - c: a pointer to a gin.Context object representing the HTTP request and response.
//
// Returns:
// - None
func (s *LibraryService) GetAllLibraries(c *gin.Context) {
	user, err := helpers.GetUserFromContext(c)
	if err != nil {
		helpers.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	libraries, err := s.repo.GetAllLibraries(userID.String())
	if err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	response := make([]LibraryResponse, 0)
	for _, library := range *libraries {
		response = append(response, LibraryResponse{
			ID:          library.ID,
			Name:        library.Name,
			Description: library.Description,
			CreatedAt:   library.CreatedAt,
			UpdatedAt:   library.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetLibraryByID retrieves a library by its ID.
//
// Parameters:
// - c: a pointer to a gin.Context object representing the HTTP request and response.
//
// Returns:
// - None
func (s *LibraryService) GetLibraryByID(c *gin.Context) {
	libraryID := c.Param("libraryId")

	user, err := helpers.GetUserFromContext(c)
	if err != nil {
		helpers.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	library, err := s.repo.GetLibraryByID(userID.String(), libraryID)
	if err != nil {
		if strings.Contains(err.Error(), "library not found") {
			helpers.HandleError(c, err, http.StatusNotFound)
			return
		}

		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, library)
}

// DeleteLibrary deletes a library by its ID.
//
// Parameters:
// - c: a pointer to a gin.Context object representing the HTTP request and response.
//
// Returns:
// - None
func (s *LibraryService) DeleteLibrary(c *gin.Context) {
	libraryID := c.Param("libraryId")

	user, err := helpers.GetUserFromContext(c)
	if err != nil {
		helpers.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	if err := s.repo.DeleteLibrary(userID.String(), libraryID); err != nil {
		if strings.Contains(err.Error(), "library not found") {
			helpers.HandleError(c, err, http.StatusNotFound)
			return
		}

		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Library deleted successfully"})
}

// UpdateLibrary updates a library in the LibraryService.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function retrieves the library ID from the request parameter,
// binds the JSON request body to a Library struct,
// parses the ID, validates the struct,
// updates the library in the repository,
// and returns an HTTP status code indicating the success of the operation.
func (s *LibraryService) UpdateLibrary(c *gin.Context) {
	id := c.Param("libraryId")

	user, err := helpers.GetUserFromContext(c)
	if err != nil {
		helpers.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	var library models.Library
	if err := c.BindJSON(&library); err != nil {
		helpers.HandleError(c, err, http.StatusBadRequest)
		return
	}

	libraryID, err := uuid.Parse(id)
	if err != nil {
		helpers.HandleError(c, err, http.StatusBadRequest)
		return
	}

	library.ID = libraryID
	library.User = *user

	if err := pkg.ValidateModelStruct(library); err != nil {
		helpers.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	if err := s.repo.UpdateLibrary(userID.String(), &library); err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// AddBookToLibrary adds a book to a library.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function retrieves the library ID and book ID from the request parameters,
// retrieves the user ID from the context,
// adds the book to the library in the repository,
// and returns an HTTP status code indicating the success of the operation.
func (s *LibraryService) AddBookToLibrary(c *gin.Context) {
	libraryID := c.Param("libraryId")
	bookID := c.Param("bookId")

	user, err := helpers.GetUserFromContext(c)
	if err != nil {
		helpers.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	if err := s.repo.AddBookToLibrary(userID.String(), libraryID, bookID); err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// RemoveBookFromLibrary removes a book from a library in the LibraryService.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function retrieves the library ID and book ID from the request parameters,
// retrieves the user ID from the context,
// removes the book from the library in the repository,
// and returns an HTTP status code indicating the success of the operation.
func (s *LibraryService) RemoveBookFromLibrary(c *gin.Context) {
	libraryID := c.Param("libraryId")
	bookID := c.Param("bookId")

	user, err := helpers.GetUserFromContext(c)
	if err != nil {
		helpers.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	if err := s.repo.RemoveBookFromLibrary(userID.String(), libraryID, bookID); err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
