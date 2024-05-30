package library

import (
	"mybooks/internal/infrastructure/model"

	"gorm.io/gorm"
)

type LibraryRepository interface {
	CreateLibrary(library *model.Library) error
	GetAllLibraries() ([]model.Library, error)
	GetLibraryByID(id string) (*model.Library, error)
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
	query := r.db.Model(&model.Library{})

	if err := query.Find(&libraries).Error; err != nil {
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

	if err := r.db.Where("id = ?", id).First(&library).Error; err != nil {
		return nil, err
	}

	return &library, nil
}
