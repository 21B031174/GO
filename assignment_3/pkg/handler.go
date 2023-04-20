package pkg

import (
	"github.com/21B031174/GolandProjects/assignment_3/database"
	"github.com/21B031174/GolandProjects/assignment_3/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBooks(c *gin.Context) {
	var books []*model.Book

	database.GetDB().Find(&books)

	if len(books) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book store is empty."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func GetBookById(c *gin.Context) {
	var book model.Book

	id := c.Param("id")
	database.GetDB().First(&book, id)

	if book.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func AddBook(c *gin.Context) {
	var input model.Book

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := model.Book{
		ID:          input.ID,
		Title:       input.Title,
		Author:      input.Author,
		Description: input.Description,
		Price:       input.Price,
	}

	database.GetDB().Create(&book)
	c.JSON(http.StatusOK, gin.H{"new book added": book})
}

func DeleteBook(c *gin.Context) {
	var book model.Book

	id := c.Param("id")

	database.GetDB().First(&book, id)
	if book.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no such book to delete"})
		return
	}

	database.GetDB().Delete(&book)
	c.JSON(http.StatusBadRequest, gin.H{"deleted": true})
}

func GetBookByTitle(c *gin.Context) {
	var book model.Book
	title := c.Param("title")
	database.GetDB().First(&book, "title=?", title)

	if book.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such title"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book model.Book

	id := c.Param("id")

	database.GetDB().First(&book, id)

	var input model.BookStruct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.GetDB().Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"book is updated": book})
}

func SortByAsc(c *gin.Context) {
	var books []model.Book

	database.GetDB().Find(&books)

	if len(books) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store is empty"})
		return
	}
	database.GetDB().Order("price").Find(&books)
	c.JSON(http.StatusOK, gin.H{"book is sorted": books})
}

func SortByDesc(c *gin.Context) {
	var books []model.Book

	database.GetDB().Find(&books)

	if len(books) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store is empty"})
		return
	}
	database.GetDB().Order("price desc").Find(&books)
	c.JSON(http.StatusOK, gin.H{"book is sorted": books})
}
