package middlewares

import (
	"fmt"
	"mybooks/internal/domain/repositories"
	"mybooks/internal/infrastructure/config"
	"mybooks/internal/infrastructure/constants"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// AuthMiddleware is a middleware function that checks if the incoming request is authenticated.
//
// It first retrieves the JWT token from a cookie named "auth" in the request.
// If the token is not found or invalid, it aborts the request with a 401 Unauthorized status.
//
// The token is then parsed and validated. If the token is not valid or has expired,
// the request is aborted with a 401 Unauthorized status.
//
// The user ID is extracted from the token claims. If the user ID is not a valid UUID,
// the request is aborted with a 401 Unauthorized status.
//
// The user is retrieved from the authentication repository using the user ID.
// If the user is not found, the request is aborted with a 401 Unauthorized status.
//
// An anonymous struct containing the user's public information is created.
// This struct is attached to the request context.
//
// The next handler in the chain is called.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token string from the cookie
		tokenString, err := c.Cookie(constants.AuthCookieName)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if token is valid and extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if token has expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Parse the user ID from the token claims
		userID, err := uuid.Parse(claims["sub"].(string))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Use the authentication repository to get the user
		repo := repositories.NewAuthRepository(config.DB())
		user, err := repo.GetUserByID(userID.String())
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if user ID is zero value of uuid.UUID
		if user.ID == uuid.Nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach the public user to the context
		c.Set("user", user)

		// Continue with the next handler
		c.Next()
	}
}
