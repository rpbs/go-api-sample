package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllBooks(c *gin.Context) {

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	var books []Book

	db.Raw("SELECT * FROM books").Scan(&books)

	c.JSON(http.StatusOK, books)
}

func getById(c *gin.Context) {

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	book := Book{}

	id := c.Param("id")

	db.Raw("SELECT * FROM books WHERE book_id = ?", id).Scan(&book)

	if book.BookID == 0 {
		c.Status(404)
		return
	}

	c.JSON(http.StatusOK, book)
}
