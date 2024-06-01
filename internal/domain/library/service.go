package library

import (
	"mybooks/internal/infrastructure/model"
	"mybooks/internal/infrastructure/utils"
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

func NewLibraryService(repo LibraryRepository) *LibraryService {
	return &LibraryService{
		repo: repo,
	}
}

func (s *LibraryService) CreateLibrary(c *gin.Context) {
	library := new(model.Library)

	id, err := utils.GenerateRandomID()
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	if err := c.BindJSON(library); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	library.ID = id

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

func (s *LibraryService) GetAllLibraries(c *gin.Context) {
	libraries, err := s.repo.GetAllLibraries()
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
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

func (s *LibraryService) GetLibraryByID(c *gin.Context) {
	libraryID := c.Param("libraryId")

	library, err := s.repo.GetLibraryByID(libraryID)
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, library)
}

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

func (s *LibraryService) UpdateLibrary(c *gin.Context) {
	id := c.Param("libraryId")

	var library model.Library
	if err := c.BindJSON(&library); err != nil {
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

func (s *LibraryService) AddBookToLibrary(c *gin.Context) {
	libraryID := c.Param("libraryId")

	var req AddBookRequest

	if err := c.BindJSON(&req); err != nil {
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
