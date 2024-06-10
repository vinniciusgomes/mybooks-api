package config

import (
	"fmt"
	"mybooks/internal/infrastructure/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database *gorm.DB

var database *gorm.DB
var e error

// DatabaseInit initializes the database connection by loading the environment variables and establishing a connection to the PostgreSQL database.
//
// This function does not take any parameters.
// It does not return any values.
func DatabaseInit() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}

	database.AutoMigrate(&model.User{}, &model.Book{}, &model.Library{}, &model.Loan{})
}

// DB returns the *gorm.DB object representing the database connection.
//
// No parameters.
// Returns *gorm.DB.
func DB() *gorm.DB {
	return database
}
