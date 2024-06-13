package repositories

import (
	"errors"
	"fmt"
	"mybooks/internal/domain/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
}

type authRepositoryImp struct {
	db *gorm.DB
}

// NewAuthRepository creates a new instance of the AuthRepository interface.
//
// It takes a *gorm.DB parameter, which represents the database connection.
// It returns an AuthRepository pointer, which is an implementation of the AuthRepository interface.
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImp{
		db: db,
	}
}

// CreateUser creates a new user in the database.
//
// It takes a pointer to a User struct as a parameter and returns an error.
// The function creates a new record in the database using the provided User object.
// If there is an error during the creation process, it returns the error.
// Otherwise, it returns nil.
func (r *authRepositoryImp) CreateUser(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("user with email %s already exists", user.Email)
		}

		return err
	}

	return nil
}

// GetUserByEmail retrieves a user from the database based on the provided email.
//
// Parameters:
// - email: a string representing the email of the user.
//
// Returns:
// - *models.User
func (r *authRepositoryImp) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID retrieves a user from the database based on the provided ID.
//
// Parameters:
// - id: a string representing the ID of the user.
//
// Returns:
// - *models.User
func (r *authRepositoryImp) GetUserByID(id string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
