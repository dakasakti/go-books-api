package routes

import (
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func Run() {
	getRoutes()
	r.Run()
}

func getRoutes() {
	v1 := r.Group("v1")
	AuthorRoutes(v1)
	BookRoutes(v1)
}