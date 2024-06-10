package endpoints

import (
	"mybooks/internal/domain/authentication"
	"mybooks/internal/infrastructure/api/middlewares"

	"github.com/gin-gonic/gin"
)

// Authentication registers the auth endpoints with the provided gin.Engine and authentication.AuthenticationService.
//
// Parameters:
// - router: a pointer to a gin.Engine object representing the HTTP router.
// - authService: a pointer to a authentication.AuthenticationService object providing the auth-related operations.
//
// Returns: None.
func Authentication(router *gin.Engine, authService *authentication.AuthenticationService) {
	v1 := router.Group("/api/v1")
	{
		authRouter := v1.Group("/auth")
		{
			authRouter.POST("/signup/credentials", authService.CreateUserWithCredentials)
			authRouter.POST("/signin/credentials", authService.SignInWithCredentials)
			authRouter.GET("/validate", middlewares.RequireAuth, authService.ValidateToken)
		}
	}
}
