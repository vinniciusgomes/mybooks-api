package library

import (
	"mybooks/internal/infrastructure/helper"
	"mybooks/internal/infrastructure/model"
	"mybooks/pkg"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LibraryService struct {
	repo LibraryRepository
}

type AddBookRequest struct {
	BookID string `json:"book_id"`
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
func NewLibraryService(repo LibraryRepository) *LibraryService {
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
	library := new(model.Library)

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

	if err := c.BindJSON(library); err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	library.ID = id
	library.UserID = user.ID
	library.User = *user

	if err := pkg.ValidateModelStruct(library); err != nil {
		helper.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	if err := s.repo.CreateLibrary(library); err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
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
	user, err := helper.GetUserFromContext(c)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	libraries, err := s.repo.GetAllLibraries(userID.String())
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
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

	user, err := helper.GetUserFromContext(c)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	library, err := s.repo.GetLibraryByID(userID.String(), libraryID)
	if err != nil {
		if strings.Contains(err.Error(), "library not found") {
			helper.HandleError(c, err, http.StatusNotFound)
			return
		}

		helper.HandleError(c, err, http.StatusInternalServerError)
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

	user, err := helper.GetUserFromContext(c)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	if err := s.repo.DeleteLibrary(userID.String(), libraryID); err != nil {
		if strings.Contains(err.Error(), "library not found") {
			helper.HandleError(c, err, http.StatusNotFound)
			return
		}

		helper.HandleError(c, err, http.StatusInternalServerError)
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

	user, err := helper.GetUserFromContext(c)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	var library model.Library
	if err := c.BindJSON(&library); err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	libraryID, err := uuid.Parse(id)
	if err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	library.ID = libraryID
	library.User = *user

	if err := pkg.ValidateModelStruct(library); err != nil {
		helper.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	if err := s.repo.UpdateLibrary(userID.String(), &library); err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// AddBookToLibrary adds a book to a library in the LibraryService.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function retrieves the library ID from the request parameter,
// binds the JSON request body to an AddBookRequest struct,
// retrieves the book ID from the request body,
// adds the book to the library in the repository,
// and returns an HTTP status code indicating the success of the operation.
func (s *LibraryService) AddBookToLibrary(c *gin.Context) {
	libraryID := c.Param("libraryId")

	user, err := helper.GetUserFromContext(c)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
		return
	}
	userID := user.ID

	var req AddBookRequest

	if err := c.BindJSON(&req); err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	bookID := req.BookID

	if err := s.repo.AddBookToLibrary(userID.String(), libraryID, bookID); err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
