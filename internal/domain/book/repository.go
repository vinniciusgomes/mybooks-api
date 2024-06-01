package book

import (
	"errors"
	"fmt"
	"mybooks/internal/infrastructure/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *model.Book) error
	GetAllBooks(filters map[string]interface{}) (*[]model.Book, error)
	GetBookById(id string) (*model.Book, error)
	DeleteBook(id string) error
	UpdateBook(book *model.Book) error
}

type bookRepositoryImp struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepositoryImp{
		db: db,
	}
}

func (r *bookRepositoryImp) CreateBook(book *model.Book) error {
	if err := r.db.Create(book).Error; err != nil {
		return err
	}

	return nil
}

func (r *bookRepositoryImp) GetAllBooks(filters map[string]interface{}) (*[]model.Book, error) {
	var books []model.Book
	query := r.db.Model(&model.Book{}).Omit("libraries")

	for key, value := range filters {
		if key == "read" {
			query = query.Where("read = ?", value)
		} else {
			query = query.Where(fmt.Sprintf("%s LIKE ?", key), fmt.Sprintf("%%%s%%", value))
		}
	}

	query = query.Order("created_at DESC")

	if err := query.Find(&books).Error; err != nil {
		return nil, err
	}

	return &books, nil
}

func (r *bookRepositoryImp) GetBookById(id string) (*model.Book, error) {
	var book model.Book
	if err := r.db.First(&book, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	return &book, nil
}

func (r *bookRepositoryImp) DeleteBook(id string) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := r.db.Exec("DELETE FROM book_library WHERE book_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&model.Book{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *bookRepositoryImp) UpdateBook(book *model.Book) error {
	if err := r.db.Model(&model.Book{}).Omit("ID", "CreatedAt").Where("id = ?", book.ID).Updates(book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("book not found")
		}

		return err
	}

	return nil
}
