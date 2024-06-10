package helper

import (
	"errors"
	"mybooks/internal/infrastructure/model"

	"github.com/gin-gonic/gin"
)

// GetUserFromContext retrieves the user from the gin.Context and returns a pointer to the model.User and an error.
//
// Parameters:
// - c: a pointer to a gin.Context object representing the HTTP request and response.
//
// Returns:
// - *model.User
func GetUserFromContext(c *gin.Context) (*model.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, errors.New("user not found")
	}

	return user.(*model.User), nil
}
