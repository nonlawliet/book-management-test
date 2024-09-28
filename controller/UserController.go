package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nonlawliet/book-management-test/models"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) LoginHandler(c *gin.Context) {
	// #1 - Create user instance
	user := new(models.User)

	// #2 - Bind request into user instance
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// #3 - Validate
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request - please fill user email",
		})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request - please fill user password",
		})
		return
	}

	// #4 - Check user info with Database
	actionUser := new(models.User)
	result := h.db.Where("username = ?", user.Username).First(actionUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid username - please verify username",
		})
		return
	}

	// #5 - Validate password
	if err := bcrypt.CompareHashAndPassword([]byte(actionUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid password - please verify password",
		})
		return
	}

	// #6 - Generate JWT signing method and claims
	var secretKey = []byte("secret-key") // should be in env file, or un-publish place
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": actionUser.ID,
		"role":    "admin",
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// #7 - Generate a completed signed JWT token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}

	// #8 - Make response
	c.JSON(http.StatusOK, gin.H{
		"message": tokenString,
	})
}
func (h *Handler) RegisterUser(c *gin.Context) {
	// #1 - Create user instance
	user := new(models.User)

	// #2 - Bind request into user instance
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// #3 - Validate
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request - please fill user email",
		})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request - please fill user password",
		})
		return
	}

	// #4 - Hash password from user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// #5 - Convert hashed password (byte) to string and put into original password
	user.Password = string(hashedPassword)

	// #6 - Create user
	result := h.db.Create(user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register user successful",
	})
}
