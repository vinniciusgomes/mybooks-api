package book

import (
	"mybooks/internal/infrastructure/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *model.Book) error
	GetAllBooks() ([]model.Book, error)
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
