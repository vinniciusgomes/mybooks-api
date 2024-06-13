package repositories

import (
	"errors"
	"fmt"
	"mybooks/internal/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	UpdatePassword(userID uuid.UUID, newPassword string) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateToken(token *models.ValidationToken) error
	GetToken(token string) (*models.ValidationToken, error)
	InvalidateToken(token *models.ValidationToken) error
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

// CreateToken creates a new token in the authentication repository.
//
// Parameters:
// - token: a pointer to a ValidationToken object representing the token to be created.
//
// Returns:
// - error: an error object if there was an issue creating the token, otherwise nil.
func (r *authRepositoryImp) CreateToken(token *models.ValidationToken) error {
	if err := r.db.Create(token).Error; err != nil {
		return err
	}

	return nil
}

// GetToken retrieves a token from the database based on the provided token string.
//
// Parameters:
// - token: a string representing the token.
//
// Returns:
// - *models.ValidationToken
func (r *authRepositoryImp) GetToken(token string) (*models.ValidationToken, error) {
	var validationToken models.ValidationToken

	if err := r.db.Where("token = ?", token).First(&validationToken).Error; err != nil {
		return nil, err
	}

	return &validationToken, nil
}

// UpdatePassword updates the password of a user in the database based on the provided user ID.
//
// - userID: a UUID representing the ID of the user.
// - newPassword: a string representing the new password to be set.
// Returns an error object if there was an issue updating the password, otherwise nil.
func (r *authRepositoryImp) UpdatePassword(userID uuid.UUID, newPassword string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("password", newPassword).Error
}

// InvalidateToken updates the validity status of a token in the database.
//
// Parameters:
// - token: a pointer to a ValidationToken object to be invalidated.
// Returns:
// - error: an error object if there was an issue invalidating the token, otherwise nil.
func (r *authRepositoryImp) InvalidateToken(token *models.ValidationToken) error {
	return r.db.Model(token).Where("Token = ?", token.Token).Update("valid", false).Error
}
