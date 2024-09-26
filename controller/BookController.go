package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nonlawliet/book-management-test/models"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) DetailBookHandler(c *gin.Context) {
	// #1 - Create book instance
	var book models.Book

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
	if result := h.db.First(&book, book.ID); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// #5 - Make response
	c.JSON(http.StatusOK, &book)
}
func (h *Handler) ListBooksHandler(c *gin.Context) {
	// #1 - Create books instance (as a slice)
	var books []models.Book

	// #2 - Query data
	if result := h.db.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// #3 - Make response
	c.JSON(http.StatusOK, &books)
}
func (h *Handler) CreateBookHandler(c *gin.Context) {
	// #1 - Create book instance
	var book models.Book

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
	if result := h.db.Create(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// #5 - Make response
	c.JSON(http.StatusOK, &book)
}
func (h *Handler) UpdateBookHandler(c *gin.Context) {
	// #1 - Create book instance
	var book models.Book

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
	if result := h.db.Updates(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// #5 - Make response
	c.JSON(http.StatusOK, &book)
}
func (h *Handler) DeleteBookHandler(c *gin.Context) {
	// #1 - Create book instance
	var book models.Book

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
	if result := h.db.Delete(&book, book.ID); result.Error != nil {
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
