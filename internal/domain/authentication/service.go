package authentication

import (
	"errors"
	"mybooks/internal/infrastructure/helper"
	"mybooks/internal/infrastructure/model"
	"mybooks/pkg"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticationService struct {
	repo AuthenticationRepository
}

// NewAuthenticationService creates a new instance of the AuthenticationService struct.
//
// It takes an AuthenticationRepository as a parameter and returns a pointer to an
// AuthenticationService.
//
// Parameters:
// - repo: an instance of the AuthenticationRepository interface.
//
// Returns:
// - *AuthenticationService: a pointer to an AuthenticationService struct.
func NewAuthenticationService(repo AuthenticationRepository) *AuthenticationService {
	return &AuthenticationService{repo: repo}
}

// CreateUser creates a new user in the AuthenticationService.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function generates a random ID, binds the JSON request body to a model.User struct,
// validates the struct, generates a hashed password, creates the user in the repository,
// and returns the ID of the created user.
//
// Parameters:
// - c: a pointer to a gin.Context.
//
// Returns:
// - None.
func (s *AuthenticationService) CreateUserWithCredentials(c *gin.Context) {
	user := new(model.User)

	id, err := pkg.GenerateRandomID()
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	if err := c.BindJSON(&user); err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	user.ID = id

	if err := pkg.ValidateModelStruct(user); err != nil {
		helper.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	user.Password = string(hash)

	if err := s.repo.CreateUser(user); err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)" {
			helper.HandleError(c, errors.New("user with email already exists"), http.StatusConflict)
			return
		}

		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": user.ID,
	})
}

// SignInWithCredentials signs in a user with their credentials.
//
// It takes a pointer to a gin.Context as a parameter and returns nothing.
// The function binds the JSON request body to a model.User struct,
// validates the struct, retrieves the user from the repository using their email,
// compares the hashed password with the provided password, generates a JWT token,
// sets the token as a cookie in the response, and returns a status code indicating success.
//
// Parameters:
// - c: a pointer to a gin.Context.
//
// Returns:
// - None.
func (s *AuthenticationService) SignInWithCredentials(c *gin.Context) {
	body := new(model.User)

	if err := c.BindJSON(&body); err != nil {
		helper.HandleError(c, err, http.StatusBadRequest)
		return
	}

	if err := pkg.ValidateModelStruct(body); err != nil {
		helper.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	user, err := s.repo.GetUserByEmail(body.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.HandleError(c, errors.New("invalid email or password"), http.StatusUnauthorized)
			return
		}

		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		helper.HandleError(c, errors.New("invalid email or password"), http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		helper.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(helper.AuthCookieName, tokenString, 3600*24*7, "", "", false, true)

	c.Status(http.StatusOK)
}

// ValidateToken validates a JWT token from a cookie in the given gin.Context.
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
func (s *AuthenticationService) ValidateToken(c *gin.Context) {
	user, err := helper.GetUserFromContext(c)
	if err != nil {
		helper.HandleError(c, err, http.StatusUnauthorized)
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
