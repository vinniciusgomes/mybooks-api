package api

import (
	"log"
	"mybooks/internal/domain/book"
	"mybooks/internal/domain/library"
	"mybooks/internal/domain/loan"
	"mybooks/internal/infrastructure/api/endpoints"
	"mybooks/internal/infrastructure/api/middlewares"
	"mybooks/internal/infrastructure/config"
	"net/http"
	"os"

	docs "mybooks/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

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
	loanService := loan.NewLoanService(loan.NewLoanRepository(config.DB()))

	// Routes
	endpoints.Libraries(router, libraryService)
	endpoints.Books(router, bookService)
	endpoints.Loan(router, loanService)

	// Others routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Start server
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	return router.Run(":" + httpPort)
}
