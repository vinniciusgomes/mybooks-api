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
	GetAllLibraries(userID string) (*[]model.Library, error)
	GetLibraryByID(userID, id string) (*model.Library, error)
	UpdateLibrary(userID string, library *model.Library) error
	DeleteLibrary(userID, id string) error
	AddBookToLibrary(userID, libraryID, bookID string) error
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
// It takes a userID as a parameter and returns a pointer to a slice of model.Library objects representing the retrieved libraries.
// If there is an error during the retrieval process, the function returns nil and the error.
//
// Parameters:
// - userID: a string representing the user ID.
//
// Returns:
// - *[]model.Library: a pointer to a slice of model.Library objects representing the retrieved libraries.
// - error: an error object if there was an issue retrieving the libraries.
func (r *libraryRepositoryImp) GetAllLibraries(userID string) (*[]model.Library, error) {
	var libraries []model.Library

	err := r.db.Where("user_id = ?", userID).Find(&libraries).Error
	if err != nil {
		return nil, err
	}

	return &libraries, nil
}

// GetLibraryByID retrieves a library from the repository by its ID.
//
// It takes a userID and an id as parameters. The function returns a pointer to a model.Library object
// representing the retrieved library, and an error if there was an issue retrieving the library.
//
// Parameters:
// - userID: a string representing the user ID.
// - id: a string representing the ID of the library to retrieve.
//
// Returns:
// - *model.Library: a pointer to the retrieved library.
// - error: an error object if there was an issue retrieving the library.
func (r *libraryRepositoryImp) GetLibraryByID(userID, id string) (*model.Library, error) {
	var library model.Library

	if err := r.db.Preload("Books").First(&library, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("library not found")
		}
		return nil, err
	}

	return &library, nil
}

// DeleteLibrary deletes a library from the libraryRepositoryImp by its ID.
//
// It takes a userID and an id as parameters. The function deletes the library from the repository
// based on the provided ID. It uses the GORM library to perform the delete operation.
// The function returns an error if there was an issue deleting the library.
//
// Parameters:
// - userID: a string representing the user ID.
// - id: a string representing the ID of the library to delete.
//
// Returns:
// - error: an error object if there was an issue deleting the library.

func (r *libraryRepositoryImp) DeleteLibrary(userID, id string) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // Re-throw panic after Rollback
		}
	}()

	// Delete the relation in the book_library table
	if err := r.db.Exec("DELETE FROM book_library WHERE library_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete the library
	result := tx.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Library{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// Check if no rows were affected (library not found)
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("library not found")
	}

	return tx.Commit().Error
}

// UpdateLibrary updates a library in the repository.
//
// It takes a userID and a library as parameters and returns an error.
// The function updates the library in the repository based on the provided ID.
// It uses the GORM library to perform the update operation.
// The function returns an error if there was an issue updating the library.
func (r *libraryRepositoryImp) UpdateLibrary(userID string, library *model.Library) error {
	result := r.db.Model(&model.Library{}).Omit("ID", "CreatedAt").Where("id = ? AND user_id = ?", library.ID, userID).Updates(library)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("library not found")
	}

	return nil
}

// AddBookToLibrary adds a book to a library in the library repository.
//
// Parameters:
// - userID: a string representing the ID of the user.
// - libraryID: a string representing the ID of the library.
// - bookID: a string representing the ID of the book.
//
// Returns:
// - error: an error object if there was an issue adding the book to the library.
func (r *libraryRepositoryImp) AddBookToLibrary(userID, libraryID, bookID string) error {
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
		err := tx.First(&library, "id = ? AND user_id = ?", libUUID, userID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("library not found")
			}

			return err
		}

		var book model.Book
		err = tx.First(&book, "id = ? AND user_id = ?", bookUUID, userID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("book not found")
			}
			return err
		}

		if err := tx.Model(&library).Association("Books").Append(&book); err != nil {
			return err
		}

		return nil
	})
}
