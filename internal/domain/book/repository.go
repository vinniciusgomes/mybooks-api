package book

import (
	"errors"
	"fmt"
	"mybooks/internal/infrastructure/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *model.Book) error
	GetAllBooks(userID string, filters map[string]interface{}) (*[]model.Book, error)
	GetBookById(userID string, id string) (*model.Book, error)
	DeleteBook(userID string, id string) error
	UpdateBook(userID string, book *model.Book) error
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
// The function takes a userID string and a map of key-value pairs representing the filters to be applied.
// The keys represent the fields to be filtered, and the values represent the values to match against.
// The function returns a pointer to a slice of model.Book objects representing the retrieved books,
// and an error if there was an issue retrieving the books.
// If there is an error during the retrieval process, the function returns nil and the error.
// The books are ordered by the creation date in descending order.
func (r *bookRepositoryImp) GetAllBooks(userID string, filters map[string]interface{}) (*[]model.Book, error) {
	var books []model.Book
	query := r.db.Model(&model.Book{}).Where("user_id = ?", userID).Omit("libraries")

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
// It takes a string parameter `userID` representing the ID of the user and a string parameter `id` representing the ID of the book to retrieve.
// The function returns a pointer to a model.Book object representing the retrieved book, and an error if there was an issue retrieving the book.
// If the book is not found, it returns nil and an error with the message "book not found".
// If there is any other error during the retrieval process, it returns nil and the error.
func (r *bookRepositoryImp) GetBookById(userID, id string) (*model.Book, error) {
	var book model.Book
	if err := r.db.First(&book, "id = ? AND user_id = ?", id, userID).Error; err != nil {
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
// - userID: a string representing the ID of the user.
// - id: a string representing the ID of the book to be deleted.
//
// Returns:
// - error: an error object if there was an issue deleting the book.
func (r *bookRepositoryImp) DeleteBook(userID, id string) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // Re-throw panic after Rollback
		}
	}()

	// Delete the relation in the book_library table
	if err := r.db.Exec("DELETE FROM book_library WHERE book_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete the book
	result := tx.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Book{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// Check if no rows were affected (book not found)
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("book not found")
	}

	return tx.Commit().Error
}

// UpdateBook updates a book in the bookRepositoryImp.
//
// It takes a userID string and a book pointer as parameters. The userID represents the ID of the user, and the book pointer represents the book to be updated.
// The function updates the specified book in the database by updating its fields except for the ID and CreatedAt.
// If the book is not found, it returns an error with the message "book not found".
// Otherwise, it returns nil.
//
// Parameters:
// - userID: a string representing the ID of the user.
// - book: a pointer to a model.Book object representing the book to be updated.
//
// Returns:
// - error: an error object if there was an issue updating the book.
func (r *bookRepositoryImp) UpdateBook(userID string, book *model.Book) error {
	result := r.db.Model(&model.Book{}).Omit("ID", "CreatedAt").Where("id = ? AND user_id = ?", book.ID, userID).Updates(book)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("book not found")
	}

	return nil
}
