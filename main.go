package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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

// # Add middleware : login & check permission
// #5 - Init app gin framework

// # Business logic : Book management (CRUD + List)

var db *gorm.DB

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
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to database")
	}

	// #4 - Migrate struct
	db.AutoMigrate(&Book{}, &User{})

	// #5 - Init app gin framework
	app := gin.New()

	// test api
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// #6 - Add login middleware

	// #7 - Add handler function
	app.GET("/books/detail", detailBookHandler)
	app.GET("/books/list", listBooksHandler)
	app.POST("/books/create", createBookHandler)
	app.PUT("/books/update", updateBookHandler)
	app.DELETE("/books/delete", deleteBookHandler)

	// #8 - Run gin engine
	app.Run()
}

func detailBookHandler(c *gin.Context) {
	// #1 - Create book instance
	var book Book

	// #2 - Bind request into book instance
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// #3 - Validate
	if book.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request - please verify book id",
		})
		return
	}

	// #4 - Query data
	if result := db.First(&book, book.ID); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// #5 - Make response
	c.JSON(http.StatusOK, &book)
}
func listBooksHandler(c *gin.Context) {
	// #1 - Create books instance (as a slice)
	var books []Book

	// #2 - Query data
	if result := db.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// #3 - Make response
	c.JSON(http.StatusOK, &books)
}
func createBookHandler(c *gin.Context) {
	// #1 - Create book instance
	var book Book

	// #2 - Bind request into book instance
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// #3 - Validate
	if book.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request - please fill book name",
		})
		return
	}
	if book.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request - please fill book author",
		})
		return
	}

	// #4 - Create data
	if result := db.Create(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// #5 - Make response
	c.JSON(http.StatusOK, &book)
}
func updateBookHandler(c *gin.Context) {
	// #1 - Create book instance
	var book Book

	// #2 - Bind request into book instance
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// #3 - Validate
	if book.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request - please verify book id",
		})
		return
	}
	if book.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request - please fill book name",
		})
		return
	}
	if book.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request - please fill book author",
		})
		return
	}

	// #4 - Updates data
	if result := db.Updates(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// #5 - Make response
	c.JSON(http.StatusOK, &book)
}
func deleteBookHandler(c *gin.Context) {
	// #1 - Create book instance
	var book Book

	// #2 - Bind request into book instance
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// #3 - Validate
	if book.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request - please verify book id",
		})
		return
	}

	// #4 - Delete book
	if result := db.Delete(&book, book.ID); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// #5 - Make response
	c.JSON(http.StatusOK, gin.H{
		"message": "delete book successful",
	})
}
