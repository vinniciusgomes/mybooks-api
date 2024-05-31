package api

import (
	"log"
	"mybooks/internal/domain/book"
	"mybooks/internal/domain/library"
	"mybooks/internal/infrastructure/api/endpoints"
	"mybooks/internal/infrastructure/config"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "mybooks/docs"
)

// StartServer initializes the server and starts it.
//
// It loads the environment variables from the .env file. If there is an error
// loading the file, it logs a fatal error.
// It creates a new Echo instance and sets up the middleware for logging and
// recovering from panics.
// It initializes the database connection and pings the database to check
// its availability.
// It creates a new book service using the book repository and the database
// connection.
// It registers the authentication, libraries, books, profile, billing, loan, and
// reading endpoints with the Echo instance.
// It adds a health check endpoint that returns "OK" with a status code of 200.
// It gets the HTTP port from the environment variable or sets it to "8080" if
// it is not set.
// It starts the server and returns any error that occurs.
//
// Returns:
// - error: An error if there was a problem starting the server.
func StartServer() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Database
	config.DatabaseInit()
	gorm := config.DB()
	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	// Services
	bookService := book.NewBookService(book.NewBookRepository(config.DB()))
	libraryService := library.NewLibraryService(library.NewLibraryRepository(config.DB()))

	// Routes
	endpoints.Authentication(e)
	endpoints.Libraries(e, libraryService)
	endpoints.Books(e, bookService)
	endpoints.Profile(e)
	endpoints.Billing(e)
	endpoints.Loan(e)
	endpoints.Reading(e)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	return e.Start(":" + httpPort)
}
