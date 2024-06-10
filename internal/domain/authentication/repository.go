package authentication

import (
	"errors"
	"fmt"
	"mybooks/internal/infrastructure/model"

	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
}

type authenticationRepositoryImp struct {
	db *gorm.DB
}

// NewAuthenticationRepository creates a new instance of the AuthenticationRepository interface.
//
// It takes a *gorm.DB parameter, which represents the database connection.
// It returns an AuthenticationRepository pointer, which is an implementation of the AuthenticationRepository interface.
func NewAuthenticationRepository(db *gorm.DB) AuthenticationRepository {
	return &authenticationRepositoryImp{
		db: db,
	}
}

// CreateUser creates a new user in the database.
//
// It takes a pointer to a User struct as a parameter and returns an error.
// The function creates a new record in the database using the provided User object.
// If there is an error during the creation process, it returns the error.
// Otherwise, it returns nil.
func (r *authenticationRepositoryImp) CreateUser(user *model.User) error {
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
// - *model.User
func (r *authenticationRepositoryImp) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

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
// - *model.User
func (r *authenticationRepositoryImp) GetUserByID(id string) (*model.User, error) {
	var user model.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
