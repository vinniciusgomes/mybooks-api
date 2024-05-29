package book

import (
	"errors"
	"fmt"
	"mybooks/internal/infrastructure/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *model.Book) error
	GetAllBooks() ([]model.Book, error)
	GetBookById(id string) (model.Book, error)
	DeleteBook(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

// CreateBook creates a new book in the bookRepository.
//
// It takes a pointer to a model.Book object as a parameter, which represents the book to be created.
// It returns an error if there was an issue creating the book in the database.
// It returns nil if the book was successfully created.
func (r *bookRepository) CreateBook(book *model.Book) error {
	if err := r.db.Create(book).Error; err != nil {
		return err
	}

	return nil
}

// GetAllBooks retrieves all books from the bookRepository.
//
// It returns a slice of model.Book objects representing the retrieved books, and an error if there was an issue retrieving the books from the database.
func (r *bookRepository) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

// GetBookById retrieves a book by its ID from the bookRepository.
//
// It takes a string parameter, which represents the ID of the book to retrieve.
// It returns a model.Book object representing the retrieved book, and an error if there was an issue retrieving the book from the database.
func (r *bookRepository) GetBookById(id string) (model.Book, error) {
	var book model.Book
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Book{}, fmt.Errorf("book not found")
		}

		return model.Book{}, err
	}

	return book, nil
}

// DeleteBook deletes a book by its ID from the bookRepository.
//
// It takes a string parameter, which represents the ID of the book to delete.
// It returns an error if there was an issue deleting the book from the database.
// If the book is not found, it returns a "book not found" error.
func (r *bookRepository) DeleteBook(id string) error {
	var book model.Book
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("book not found")
		}

		return err
	}

	return r.db.Delete(&book).Error
}

// TODO: update book
