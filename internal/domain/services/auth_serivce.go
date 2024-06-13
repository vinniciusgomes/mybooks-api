package services

import (
	"errors"
	"fmt"
	"mybooks/internal/domain/models"
	"mybooks/internal/domain/repositories"
	"mybooks/internal/infrastructure/constants"
	"mybooks/internal/infrastructure/helpers"
	"mybooks/pkg"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repo repositories.AuthRepository
}

// NewAuthService creates a new instance of the AuthService struct.
//
// It takes an AuthRepository as a parameter and returns a pointer to an
// AuthService.
//
// Parameters:
// - repo: an instance of the AuthRepository interface.
//
// Returns:
// - *AuthService: a pointer to an AuthService struct.
func NewAuthService(repo repositories.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser creates a new user in the AuthService.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function generates a random ID, binds the JSON request body to a models.User struct,
// validates the struct, generates a hashed password, creates the user in the repository,
// and returns the ID of the created user.
//
// Parameters:
// - c: a pointer to a gin.Context.
//
// Returns:
// - None.
func (s *AuthService) CreateUserWithCredentials(c *gin.Context) {
	user := new(models.User)

	id, err := pkg.GenerateRandomID()
	if err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	if err := c.BindJSON(&user); err != nil {
		helpers.HandleError(c, err, http.StatusBadRequest)
		return
	}

	user.ID = id

	if err := pkg.ValidateModelStruct(user); err != nil {
		helpers.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	user.Password = string(hash)

	if err := s.repo.CreateUser(user); err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)" {
			helpers.HandleError(c, errors.New("user with email already exists"), http.StatusConflict)
			return
		}

		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": user.ID,
	})
}

// SignInWithCredentials signs in a user with their credentials.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function binds the JSON request body to a models.User struct,
// validates the struct, retrieves the user from the repository using their email,
// compares the hashed password with the provided password, generates a JWT token,
// sets the token as a cookie in the response, and returns a status code indicating success.
//
// Parameters:
// - c: a pointer to a gin.Context.
//
// Returns:
// - None.
func (s *AuthService) SignInWithCredentials(c *gin.Context) {
	body := new(models.User)

	if err := c.BindJSON(&body); err != nil {
		helpers.HandleError(c, err, http.StatusBadRequest)
		return
	}

	if err := pkg.ValidateModelStruct(body); err != nil {
		helpers.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	user, err := s.repo.GetUserByEmail(body.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.HandleError(c, errors.New("invalid email or password"), http.StatusUnauthorized)
			return
		}

		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		helpers.HandleError(c, errors.New("invalid email or password"), http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(constants.AuthCookieName, tokenString, 3600*24*7, "", "", false, true)

	c.Status(http.StatusOK)
}

// SignOut signs out the user by clearing the authentication cookie in the given gin.Context.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function sets the "access_token" cookie with an empty value and an expiration time of -1,
// effectively removing the authentication cookie from the browser.
// It then sets the status code of the response to 200 OK.
//
// Parameters:
// - c: a pointer to a gin.Context.
//
// Returns:
// - None.
func (s *AuthService) SignOut(c *gin.Context) {
	c.SetCookie(constants.AuthCookieName, "", -1, "", "", false, true)
	c.Status(http.StatusOK)
}

// ForgotPassword handles the request to reset a user's password.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function binds the JSON request body to a struct, validates the struct,
// retrieves the user from the repository using their email, generates a secure token,
// creates a validation token in the repository, sends an email with a reset link,
// and returns a JSON response indicating success.
//
// Parameters:
// - c: a pointer to a gin.Context.
//
// Returns:
// - None.
func (s *AuthService) ForgotPassword(c *gin.Context) {
	var body struct {
		Email string `json:"email" validate:"required,email"`
	}

	if err := c.BindJSON(&body); err != nil {
		helpers.HandleError(c, err, http.StatusBadRequest)
		return
	}

	if err := pkg.ValidateModelStruct(body); err != nil {
		helpers.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	user, err := s.repo.GetUserByEmail(body.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "email sent successfully",
			})
			return
		}

		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	tokenString, err := helpers.GenerateSecureToken()
	if err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	token := models.ValidationToken{
		Token:     tokenString,
		Type:      "password_reset",
		Valid:     true,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	if err := s.repo.CreateToken(&token); err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	resetURL := fmt.Sprintf("%s/reset-password/%s", os.Getenv("APP_URL"), tokenString)
	err = pkg.SendEmail([]string{user.Email}, "Reset Password", fmt.Sprintf("Click the link to reset your password: <a href='%s' target='_blank'>Reset Password</a>", resetURL))
	if err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "email sent",
	})
}

// ResetPassword resets the password for a user using a provided token.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function retrieves the token string from the request parameters,
// binds the JSON body to a struct, validates the struct, retrieves the token
// from the repository, checks if the token is valid and not expired, hashes
// the provided password, updates the user's password in the repository, invalidates
// the token, and returns an HTTP status code indicating the success of the operation.
func (s *AuthService) ResetPassword(c *gin.Context) {
	tokenString := c.Param("token")

	var body struct {
		Password string `json:"password" gorm:"not null;size:100" validate:"required,min=1,max=100"`
	}

	if err := c.BindJSON(&body); err != nil {
		helpers.HandleError(c, err, http.StatusBadRequest)
		return
	}

	if err := pkg.ValidateModelStruct(body); err != nil {
		helpers.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	token, err := s.repo.GetToken(tokenString)
	if err != nil {
		helpers.HandleError(c, errors.New("invalid or expired token"), http.StatusBadRequest)
		return
	}

	if !token.Valid || time.Now().After(token.ExpiresAt) {
		token.Valid = false
		if err := s.repo.InvalidateToken(token); err != nil {
			helpers.HandleError(c, err, http.StatusInternalServerError)
			return
		}

		helpers.HandleError(c, errors.New("invalid or expired token"), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	token.Valid = false
	if err := s.repo.InvalidateToken(token); err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	if err := s.repo.UpdatePassword(token.UserID, string(hashedPassword)); err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "password reset successful",
	})
}

// ValidateAuthToken validates a JWT token from a cookie in the given gin.Context.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function extracts the token from the "access_token" cookie in the context,
// parses it using the JWT secret from the environment variable "JWT_SECRET",
// and checks if the token is still valid.
//
// Parameters:
// - c: a pointer to a gin.Context.
//
// Returns:
// - None.
func (s *AuthService) ValidateAuthToken(c *gin.Context) {
	user, err := helpers.GetUserFromContext(c)
	if err != nil {
		helpers.HandleError(c, err, http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"user": map[string]interface{}{
			"id":         user.ID,
			"email":      user.Email,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	}

	c.JSON(http.StatusOK, response)
}
