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
	GetAllLibraries() (*[]model.Library, error)
	GetLibraryByID(id string) (*model.Library, error)
	UpdateLibrary(library *model.Library) error
	DeleteLibrary(id string) error
	AddBookToLibrary(libraryID string, bookID string) error
}

type libraryRepositoryImp struct {
	db *gorm.DB
}

// NewLibraryRepository creates a new instance of the LibraryRepository interface.
//
// It takes a *gorm.DB parameter, which represents the database connection.
// It returns a LibraryRepository pointer, which is an implementation of the LibraryRepository interface.
func NewLibraryRepository(db *gorm.DB) LibraryRepository {
	return &libraryRepositoryImp{
		db: db,
	}
}

// CreateLibrary creates a new library in the database.
//
// It takes a pointer to a Library struct as a parameter and returns an error.
func (r *libraryRepositoryImp) CreateLibrary(library *model.Library) error {
	return r.db.Create(library).Error
}

// GetAllLibraries retrieves all libraries from the library repository.
//
// It returns a pointer to a slice of model.Library objects representing the retrieved libraries.
// If there is an error during the retrieval process, the function returns nil and the error.
//
// Returns:
// - *[]model.Library: a pointer to a slice of model.Library objects representing the retrieved libraries.
// - error: an error object if there was an issue retrieving the libraries.
func (r *libraryRepositoryImp) GetAllLibraries() (*[]model.Library, error) {
	var libraries []model.Library
	if err := r.db.Model(&model.Library{}).Find(&libraries).Error; err != nil {
		return nil, err
	}

	return &libraries, nil
}

// GetLibraryByID retrieves a library from the repository by its ID.
//
// It takes a string parameter `id` representing the ID of the library to retrieve.
// The function returns a pointer to a model.Library object representing the retrieved library,
// and an error if there was an issue retrieving the library.
// If the library is not found, it returns nil and an error with the message "record not found".
// If there is any other error during the retrieval process, it returns nil and the error.
//
// Parameters:
// - id: the ID of the library to retrieve.
//
// Returns:
// - *model.Library: a pointer to the retrieved library.
// - error: an error object if there was an issue retrieving the library.
func (r *libraryRepositoryImp) GetLibraryByID(id string) (*model.Library, error) {
	var library model.Library
	if err := r.db.Preload("Books").First(&library, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &library, nil
}

// DeleteLibrary deletes a library from the libraryRepositoryImp by its ID.
//
// Parameters:
// - id: a string representing the ID of the library to be deleted.
//
// Returns:
// - error: an error object if there was an issue deleting the library.
func (r *libraryRepositoryImp) DeleteLibrary(id string) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Exec("DELETE FROM book_library WHERE library_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&model.Library{}).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("library not found")
		}
		return err
	}

	return tx.Commit().Error
}

// UpdateLibrary updates a library in the repository.
//
// It takes a pointer to a Library struct as a parameter and returns an error.
// The function updates the library in the repository based on the provided ID.
// It uses the GORM library to perform the update operation.
// The function returns an error if there was an issue updating the library.
func (r *libraryRepositoryImp) UpdateLibrary(library *model.Library) error {
	return r.db.Model(&model.Library{}).Where("id = ?", library.ID).Updates(library).Error
}

// AddBookToLibrary adds a book to a library in the library repository.
//
// Parameters:
// - libraryID: a string representing the ID of the library.
// - bookID: a string representing the ID of the book.
//
// Returns:
// - error: an error object if there was an issue adding the book to the library.
func (r *libraryRepositoryImp) AddBookToLibrary(libraryID string, bookID string) error {
	libUUID, err := uuid.Parse(libraryID)
	if err != nil {
		return err
	}
	bookUUID, err := uuid.Parse(bookID)
	if err != nil {
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		var library model.Library
		if err := tx.First(&library, "id = ?", libUUID).Error; err != nil {
			return err
		}

		var book model.Book
		if err := tx.First(&book, "id = ?", bookUUID).Error; err != nil {
			return err
		}

		if err := tx.Model(&library).Association("Books").Append(&book); err != nil {
			return err
		}

		return nil
	})
}
