package library

import (
	"errors"
	"fmt"
	"mybooks/internal/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LibraryRepository interface {
	CreateLibrary(library *model.Library) error
	GetAllLibraries() ([]model.Library, error)
	GetLibraryByID(id string) (*model.Library, error)
	DeleteLibrary(id string) error
	AddBookToLibrary(libraryID string, bookID string) error
}

type libraryRepository struct {
	db *gorm.DB
}

// NewLibraryRepository creates a new instance of LibraryRepository using the provided *gorm.DB.
//
// Parameters:
// - db: The *gorm.DB object representing the database connection.
//
// Returns:
// - LibraryRepository: The newly created instance of LibraryRepository.
func NewLibraryRepository(db *gorm.DB) LibraryRepository {
	return &libraryRepository{
		db: db,
	}
}

// CreateLibrary creates a new library in the libraryRepository.
//
// It takes a pointer to a model.Library object as a parameter, which represents the library to be created.
// It returns an error if there was an issue creating the library in the database.
// It returns nil if the library was successfully created.
func (r *libraryRepository) CreateLibrary(library *model.Library) error {
	if err := r.db.Create(library).Error; err != nil {
		return err
	}

	return nil
}

// GetAllLibraries retrieves all libraries from the libraryRepository.
//
// It returns a slice of model.Library objects representing all the libraries in the repository,
// and an error if there was an issue retrieving the libraries from the database.
// If no error occurs, it returns nil.
func (r *libraryRepository) GetAllLibraries() ([]model.Library, error) {
	var libraries []model.Library

	if err := r.db.Model(&model.Library{}).Select("id", "name", "description", "created_at", "updated_at").Find(&libraries).Error; err != nil {
		return nil, err
	}

	return libraries, nil
}

// GetLibraryByID retrieves a library by its ID from the libraryRepository.
//
// Parameters:
// - id: The ID of the library to retrieve.
//
// Returns:
// - *model.Library: The library with the specified ID, or nil if not found.
// - error: An error if there was a problem retrieving the library from the database.
func (r *libraryRepository) GetLibraryByID(id string) (*model.Library, error) {
	var library model.Library

	if err := r.db.Preload("Books").Where("id = ?", id).First(&library).Error; err != nil {
		return nil, err
	}

	return &library, nil
}

// DeleteLibrary deletes a library from the libraryRepository.
//
// Parameters:
// - id: The ID of the library to delete.
//
// Returns:
// - error: An error if there was a problem deleting the library, or if the library was not found.
func (r *libraryRepository) DeleteLibrary(id string) error {
	var library model.Library
	if err := r.db.Where("id = ?", id).First(&library).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("library not found")
		}

		return err
	}

	return r.db.Delete(&library).Error
}

// AddBookToLibrary adds a book to a library in the libraryRepository.
//
// Parameters:
// - libraryID: The ID of the library to add the book to.
// - bookID: The ID of the book to add.
//
// Returns:
// - error: An error if there was a problem adding the book to the library.
func (r *libraryRepository) AddBookToLibrary(libraryID string, bookID string) error {
	libUUID, err := uuid.Parse(libraryID)
	if err != nil {
		return err
	}
	bookUUID, err := uuid.Parse(bookID)
	if err != nil {
		return err
	}

	var library model.Library
	if err := r.db.Preload("Books").First(&library, "id = ?", libUUID).Error; err != nil {
		return err
	}

	var book model.Book
	if err := r.db.First(&book, "id = ?", bookUUID).Error; err != nil {
		return err
	}

	if err := r.db.Model(&library).Association("Books").Append(&book); err != nil {
		return err
	}

	return nil
}
