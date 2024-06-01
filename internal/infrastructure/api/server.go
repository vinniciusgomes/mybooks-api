package api

import (
	"log"
	"mybooks/internal/domain/book"
	"mybooks/internal/domain/library"
	"mybooks/internal/infrastructure/api/endpoints"
	"mybooks/internal/infrastructure/api/middlewares"
	"mybooks/internal/infrastructure/config"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// StartServer initializes the server and starts it.
//
// It loads the environment variables from the .env file. If there is an error
// loading the file, it logs a fatal error.
// It creates a new Gin instance and sets up the middleware for logging and
// recovering from panics.
// It initializes the database connection and pings the database to check
// its availability.
// It creates a new book service using the book repository and the database
// connection.
// It registers the authentication, libraries, books, profile, billing, loan, and
// reading endpoints with the Gin instance.
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

	r := gin.Default()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())

	// Database
	config.DatabaseInit()
	gorm := config.DB()
	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	err = dbGorm.Ping()
	if err != nil {
		panic(err)
	}

	// Services
	bookService := book.NewBookService(book.NewBookRepository(config.DB()))
	libraryService := library.NewLibraryService(library.NewLibraryRepository(config.DB()))

	// Routes
	endpoints.Authentication(r)
	endpoints.Libraries(r, libraryService)
	endpoints.Books(r, bookService)
	endpoints.Profile(r)
	endpoints.Billing(r)
	endpoints.Loan(r)
	endpoints.Reading(r)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	return r.Run(":" + httpPort)
}
