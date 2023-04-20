package main

import (
	"github.com/21B031174/GolandProjects/assignment_3/database"
	"github.com/21B031174/GolandProjects/assignment_3/pkg"
	"github.com/gin-gonic/gin"
)
import _ "github.com/lib/pq"

func main() {
	database.Init()
	r := gin.Default()
	r.GET("/books", pkg.GetBooks)
	r.GET("/books/:id", pkg.GetBookById)

	r.POST("/books", pkg.AddBook)

	r.DELETE("/books/:id", pkg.DeleteBook)
	r.PATCH("/books/:id", pkg.UpdateBook)

	r.GET("/books/title/:title", pkg.GetBookByTitle)

	r.GET("/books/sortByAsc", pkg.SortByAsc)
	r.GET("/books/sortByDesc", pkg.SortByDesc)

	r.Run(":8080")
}
