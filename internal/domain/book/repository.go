package book

import (
	"errors"
	"fmt"
	"mybooks/internal/infrastructure/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *model.Book) error
	GetAllBooks(filters map[string]interface{}) ([]model.Book, error)
	GetBookById(id string) (model.Book, error)
	DeleteBook(id string) error
	UpdateBook(book *model.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

// NewBookRepository creates a new instance of BookRepository using the provided *gorm.DB.
//
// Parameters:
// - db: The *gorm.DB object representing the database connection.
//
// Returns:
// - BookRepository: The newly created instance of BookRepository.
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

// GetAllBooks retrieves all books from the bookRepository based on the provided filters.
//
// The filters parameter is a map of key-value pairs used to filter the books. The keys represent the fields to filter on,
// and the values represent the values to match against. The supported keys are "read" and any other field of the Book model.
// The "read" key filters the books based on the "read" field, while any other key filters the books based on a LIKE match
// with the corresponding field.
//
// The function returns a slice of model.Book objects representing the retrieved books. If there was an error retrieving the books,
// the function returns nil and the error.
func (r *bookRepository) GetAllBooks(filters map[string]interface{}) ([]model.Book, error) {
	var books []model.Book
	query := r.db.Model(&model.Book{})

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

	return books, nil
}

// GetBookById retrieves a book from the bookRepository by its ID.
//
// Parameters:
// - id: the ID of the book to retrieve.
//
// Returns:
// - model.Book: the retrieved book.
// - error: an error if the book was not found or there was an issue retrieving it.
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

// UpdateBook updates a book in the bookRepository.
//
// It takes a pointer to a model.Book object as a parameter, which represents the book to be updated.
// It returns an error if there was an issue updating the book in the database.
// If the book is not found, it returns a "book not found" error.
func (r *bookRepository) UpdateBook(book *model.Book) error {
	var existingBook model.Book
	if err := r.db.Where("id = ?", book.ID).First(&existingBook).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("book not found")
		}
		return err
	}

	if err := r.db.Model(&existingBook).Omit("ID", "CreatedAt").Updates(book).Error; err != nil {
		return err
	}

	return nil
}
