package routes

import (
	BookHandler "github.com/fathxn/go-books-restapi/handlers/BookHandler"
	"github.com/gin-gonic/gin"
)

func BookRoutes(rg *gin.RouterGroup) {
	author := rg.Group("/books")

	author.GET("/", BookHandler.Index)
	author.GET("/:id", BookHandler.Detail)
	author.POST("/", BookHandler.Create)
	// author.PUT("/:id", BookHandler.Update)
	// author.DELETE("/:id", BookHandler.Delete)
}