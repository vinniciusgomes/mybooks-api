package handlers

import (
	"mybooks/internal/domain/services"
	"mybooks/internal/infrastructure/api/middlewares"

	"github.com/gin-gonic/gin"
)

// AuthHandler registers the auth handler with the provided gin.Engine and services.AuthService.
//
// Parameters:
// - router: a pointer to a gin.Engine object representing the HTTP router.
// - authService: a pointer to a services.AuthService object providing the auth-related operations.
//
// Returns: None.
func AuthHandler(router *gin.Engine, authService *services.AuthService) {
	v1 := router.Group("/v1")
	{
		authRouter := v1.Group("/auth")
		{
			authRouter.POST("/signup/credentials", authService.CreateUserWithCredentials)
			authRouter.POST("/signin/credentials", authService.SignInWithCredentials)
			authRouter.POST("/signout", authService.SignOut)
			authRouter.POST("/forgot-password", authService.ForgotPassword)
			authRouter.POST("/reset-password/:token", authService.ResetPassword)
			authRouter.GET("/validate-token", middlewares.AuthMiddleware(), authService.ValidateAuthToken)
		}
	}
}
