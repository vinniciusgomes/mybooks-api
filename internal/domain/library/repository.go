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

func NewLibraryRepository(db *gorm.DB) LibraryRepository {
	return &libraryRepositoryImp{
		db: db,
	}
}

func (r *libraryRepositoryImp) CreateLibrary(library *model.Library) error {
	return r.db.Create(library).Error
}

func (r *libraryRepositoryImp) GetAllLibraries() (*[]model.Library, error) {
	var libraries []model.Library
	if err := r.db.Model(&model.Library{}).Find(&libraries).Error; err != nil {
		return nil, err
	}

	return &libraries, nil
}

func (r *libraryRepositoryImp) GetLibraryByID(id string) (*model.Library, error) {
	var library model.Library
	if err := r.db.Preload("Books").First(&library, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &library, nil
}

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

func (r *libraryRepositoryImp) UpdateLibrary(library *model.Library) error {
	return r.db.Model(&model.Library{}).Where("id = ?", library.ID).Updates(library).Error
}

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
