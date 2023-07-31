package routes

import (
	AuthorHandler "github.com/fathxn/go-books-restapi/handlers/AuthorHandler"
	"github.com/gin-gonic/gin"
)

func AuthorRoutes(rg *gin.RouterGroup) {
	author := rg.Group("/authors")

	author.GET("/", AuthorHandler.Index)
	author.GET("/:id", AuthorHandler.Detail)
	author.POST("/", AuthorHandler.Create)
	author.PUT("/:id", AuthorHandler.Update)
	author.DELETE("/:id", AuthorHandler.Delete)
}