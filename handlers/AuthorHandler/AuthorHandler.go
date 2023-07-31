package authorhandler

import (
	"github.com/fathxn/go-books-restapi/config"
	"github.com/fathxn/go-books-restapi/models"
	"github.com/gin-gonic/gin"
)

// Show all data of authors
func Index(c *gin.Context) {
	var authors []models.Author
	
	config.DB.Find(&authors)
	c.JSON(200, gin.H{"authors": authors})
}

// Show detail of author by the Id
// func Detail(c *gin.Context) {
// 	var author []models.Author
// 	id := c.Param("id")

// 	if err := config.DB.First(&author, id).Error; err != nil {
// 		switch err {
// 		case gorm.ErrRecordNotFound:
// 			c.AbortWithStatusJSON(404, gin.H{
// 				"message": "Data not available",
// 			})
// 			return
// 		default:
// 			c.AbortWithStatusJSON(500, gin.H{
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}
// 	c.JSON(200, gin.H{
// 		"author": author,
// 	})
// }

func Detail(c *gin.Context) {
	var book models.Book
	// var bookResponse models.BookResponse

	id := c.Param("id")

	config.DB.First(&book, id)

	c.JSON(200, gin.H{"authors": book})	
}

// Create/store a new author data into database
func Create(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid data",
		})
		return
	}

	config.DB.Create(&author)

	c.JSON(200, gin.H{
		"status": "author added",
	})
}

// Update author data by Id in database
func Update(c *gin.Context) {
	var author models.Author
	id := c.Param("id")

	if err := c.ShouldBindJSON(&author); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if config.DB.Model(&author).Where("id = ?", id).Updates(&author).RowsAffected == 0 {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Data not available",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Data updated successfully",
	})
}

// Delete/remove data from database by Id
func Delete(c *gin.Context) {
	var author models.Author
	
	if err := config.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "Data not available",
		})
		return
	}

	config.DB.Delete(&author)

	c.JSON(200, gin.H{
		"message": "Data deleted",
	})
}