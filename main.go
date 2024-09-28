package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
// #7 - Add login middleware
// #8 - Add handler function
// #9 - Run gin engine

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
	protected := app.Group("/", AuthRequired)

	// #8 - Add handler function
	app.POST("/register", handler.RegisterUser)
	app.POST("/login", handler.LoginHandler)

	protected.GET("/books/detail", handler.DetailBookHandler)
	protected.GET("/books/list", handler.ListBooksHandler)
	protected.POST("/books/create", handler.CreateBookHandler)
	protected.PUT("/books/update", handler.UpdateBookHandler)
	protected.DELETE("/books/delete", handler.DeleteBookHandler)

	// #9 - Run gin engine
	app.Run()
}

func AuthRequired(c *gin.Context) {
	// secretKey data
	var secretKey = []byte("secret-key") // should be in env file, or un-publish place

	// Check request info
	fmt.Printf("URL = %s, Method = %s, Time = %s\n", c.Request.URL, c.Request.Method, time.Now())

	// #1 - Get token from header
	s := c.Request.Header.Get("Authorization")
	tokenHeader := strings.TrimPrefix(s, "Bearer ")

	// #2 - Get token data
	token, err := jwt.ParseWithClaims(tokenHeader, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token or expired",
		})
		c.Abort()
		return
	}

	// #3 - Check role
	if token.Claims.(jwt.MapClaims)["role"] != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid role",
		})
		c.Abort()
		return
	}

	c.Next()
}
