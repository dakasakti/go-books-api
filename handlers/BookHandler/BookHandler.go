package bookhandler

import (
	"github.com/fathxn/go-books-restapi/config"
	"github.com/fathxn/go-books-restapi/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Show all books data
func Index(c *gin.Context) {
	var books []models.Book
	var bookResponse []models.BookResponse

	config.DB.Joins("Author").Find(&books).Find(&bookResponse)
	c.JSON(200, gin.H{"books": bookResponse})
}

// Show detail of book by the Id
func Detail(c *gin.Context) {
	// id := c.Param("id")

	// var book models.Book
	// var bookResponse models.BookResponse

	book := models.Book{
		ID: c.Param("id"),
	}

	// config.DB.Joins("Author").First(&book, "id = ?", id).First(&bookResponse)
	// if err := config.DB.First(&book, "id = ?", id).First(&bookResponse).Error; err != nil {
	// 	c.JSON(400, gin.H{"msg": "book not found"})
	// 	return
	// }

	if err := config.DB.Preload("Author").First(&book).Error; err != nil {
		c.JSON(400, gin.H{"msg": "book not found"})
		return
	}

	bookResponse := models.BookResponse{
		ID:          book.ID,
		BookTitle:   book.BookTitle,
		BookCover:   book.BookCover,
		Description: book.Description,
		AuthorID:    book.AuthorID,
		Author:      models.AuthorBookResponse{ID: book.Author.ID, Name: book.Author.Name},
		TotalPages:  book.TotalPages,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}

	c.JSON(200, gin.H{"book": bookResponse})
}

// Create/store a new book data into database
func Create(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid data",
		})
		return
	}

	book.ID = uuid.New().String()

	config.DB.Create(&book)

	c.JSON(200, gin.H{
		"status": "book added",
	})
}
