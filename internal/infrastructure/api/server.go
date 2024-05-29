package api

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// authHandler := NewAuthHandler()                   // Initialize auth handler
	// libraryHandler := NewLibraryHandler()             // Initialize library handler
	// bookHandler := NewBookHandler()                   // Initialize book handler
	// profileHandler := NewProfileHandler()             // Initialize profile handler
	// billingHandler := NewBillingHandler()             // Initialize billing handler
	// loanHandler := NewLoanHandler()                   // Initialize loan handler
	// readingStatusHandler := NewReadingStatusHandler() // Initialize reading status handler

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Authentication routes
	e.POST("/auth", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST /auth")
	})
	e.POST("/auth/google", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST /auth/google")
	})

	// Library routes
	e.GET("/libraries", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET /libraries")
	})
	e.POST("/libraries", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST /libraries")
	})
	e.PUT("/libraries/:libraryId", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /libraries/:libraryId")
	})
	e.DELETE("/libraries/:libraryId", func(c echo.Context) error {
		return c.String(http.StatusOK, "DELETE /libraries/:libraryId")
	})
	e.POST("/libraries/:libraryId/books", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST /libraries/:libraryId/books")
	})

	// Book routes
	e.GET("/books", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET /books")
	})
	e.POST("/books", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST /books")
	})
	e.PUT("/books/:bookId", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /books/:bookId")
	})
	e.DELETE("/books/:bookId", func(c echo.Context) error {
		return c.String(http.StatusOK, "DELETE /books/:bookId")
	})

	// Profile routes
	e.PUT("/profile/photo", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /profile/photo")
	})
	e.PUT("/profile", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /profile")
	})
	e.DELETE("/profile", func(c echo.Context) error {
		return c.String(http.StatusOK, "DELETE /profile")
	})

	// Billing routes
	e.GET("/billing", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET /billing")
	})
	e.POST("/subscribe", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST /subscribe")
	})

	// Loan routes
	e.PUT("/books/:bookId/loan", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /books/:bookId/loan")
	})
	e.PUT("/books/:bookId/return", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /books/:bookId/return")
	})

	// Reading status routes
	e.PUT("/books/:bookId/read", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT /books/:bookId/read")
	})
	e.DELETE("/books/:bookId/read", func(c echo.Context) error {
		return c.String(http.StatusOK, "DELETE /books/:bookId/read")
	})

	// Start server
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	return e.Start(":" + httpPort)
}
