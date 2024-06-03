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

// NewBookRepository creates a new instance of the BookRepository interface.
//
// It takes a *gorm.DB parameter, which represents the database connection.
// It returns a BookRepository pointer, which is an implementation of the BookRepository interface.
func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepositoryImp{
		db: db,
	}
}

// CreateBook creates a new book in the bookRepositoryImp.
//
// It takes a book pointer as a parameter and returns an error if there was an issue creating the book.
// The function creates a new record in the database using the provided book object.
// If there is an error during the creation process, it returns the error.
// Otherwise, it returns nil.
func (r *bookRepositoryImp) CreateBook(book *model.Book) error {
	if err := r.db.Create(book).Error; err != nil {
		return err
	}

	return nil
}

// GetAllBooks retrieves all books from the bookRepositoryImp that match the provided filters.
//
// The filters parameter is a map of key-value pairs representing the filters to be applied.
// The keys represent the fields to be filtered, and the values represent the values to match against.
// The function returns a pointer to a slice of model.Book objects representing the retrieved books.
// If there is an error during the retrieval process, the function returns nil and the error.
//
// Parameters:
// - filters: a map[string]interface{} representing the filters to be applied.
//
// Returns:
// - *[]model.Book: a pointer to a slice of model.Book objects representing the retrieved books.
// - error: an error object if there was an issue retrieving the books.
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

// GetBookById retrieves a book from the bookRepositoryImp by its ID.
//
// It takes a string parameter `id` representing the ID of the book to retrieve.
// The function returns a pointer to a model.Book object representing the retrieved book,
// and an error if there was an issue retrieving the book.
// If the book is not found, it returns nil and an error with the message "book not found".
// If there is any other error during the retrieval process, it returns nil and the error.
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

// DeleteBook deletes a book from the bookRepositoryImp by its ID.
//
// Parameters:
// - id: a string representing the ID of the book to be deleted.
//
// Returns:
// - error: an error object if there was an issue deleting the book.
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

// UpdateBook updates a book in the bookRepositoryImp.
//
// It takes a book pointer as a parameter and returns an error if there was an issue updating the book.
// The function updates the specified book in the database by updating its fields except for the ID and CreatedAt.
// If the book is not found, it returns an error with the message "book not found".
// Otherwise, it returns nil.
func (r *bookRepositoryImp) UpdateBook(book *model.Book) error {
	if err := r.db.Model(&model.Book{}).Omit("ID", "CreatedAt").Where("id = ?", book.ID).Updates(book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("book not found")
		}

		return err
	}

	return nil
}
