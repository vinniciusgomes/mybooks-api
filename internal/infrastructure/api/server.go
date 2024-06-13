package api

import (
	"log"
	"mybooks/internal/domain/repositories"
	"mybooks/internal/domain/services"
	"mybooks/internal/infrastructure/api/handlers"
	"mybooks/internal/infrastructure/api/middlewares"
	"mybooks/internal/infrastructure/config"
	"net/http"
	"os"

	"mybooks/docs"

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
// reading handlers with the Gin instance.
// It adds a health check handler that returns "OK" with a status code of 200.
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
	docs.SwaggerInfo.BasePath = "/v1"

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
	authService := services.NewAuthService(repositories.NewAuthRepository(config.DB()))
	bookService := services.NewBookService(repositories.NewBookRepository(config.DB()))
	libraryService := services.NewLibraryService(repositories.NewLibraryRepository(config.DB()))
	loanService := services.NewLoanService(repositories.NewLoanRepository(config.DB()))

	// Routes
	handlers.AuthHandler(router, authService)
	handlers.LibrariesHandler(router, libraryService)
	handlers.BooksHandler(router, bookService)
	handlers.LoanHandler(router, loanService)

	// Others routes
	router.GET("/v1/health", func(c *gin.Context) {
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
