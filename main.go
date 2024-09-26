package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nonlawliet/book-management-test/controller"
	"github.com/nonlawliet/book-management-test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

// Validate and Flow
// # Connect PostgreSQL Database & Automigrate
// #1 - Config connection string
// #2 - Config logger
// #3 - Connect PostgreSQL database
// #4 - Migrate struct
// #5 - Add db instance to handler struct instead

// # Add middleware : login & check permission
// #6 - Init app gin framework

// # Business logic : Book management (CRUD + List)

func main() {
	// #1 - Config connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// #2 - Config logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// #3 - Connect PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to database")
	}

	// #4 - Migrate struct
	db.AutoMigrate(&models.Book{}, &models.User{})

	// #5 - Add db instance to handler struct instead
	handler := controller.NewHandler(db)

	// #6 - Init app gin framework
	app := gin.New()

	// test api
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// #7 - Add login middleware

	// #8 - Add handler function
	app.GET("/books/detail", handler.DetailBookHandler)
	app.GET("/books/list", handler.ListBooksHandler)
	app.POST("/books/create", handler.CreateBookHandler)
	app.PUT("/books/update", handler.UpdateBookHandler)
	app.DELETE("/books/delete", handler.DeleteBookHandler)

	// #9 - Run gin engine
	app.Run()
}
