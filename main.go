package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)
type book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"` 
	Quantity int `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "1984", Author: "Eric Arthur Blair", Quantity: 3},
	{ID: "2", Title: "Fountainhead", Author: "Ayn Rand", Quantity: 2},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id") 
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}
	
	c.IndentedJSON(http.StatusOK, book)
}
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil 
		}
	}
	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
	router.GET("/books/:id", bookById)
}
   

