package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"errors"
)
type book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"` 
	Quantity int `json:"quantity"`
<<<<<<< HEAD
}

var books = []book{
	{ID: "1", Title: "1984", Author: "Eric Arthur Blair", Quantity: 3},
	{ID: "2", Title: "Fountainhead", Author: "Ayn Rand", Quantity: 2},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
   

