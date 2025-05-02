package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func getAllBooks(c *gin.Context) {

	password := url.QueryEscape("yourStrong(!)Password")
	dsn := "sqlserver://sa:" + password + "@localhost:1433?database=library&encrypt=true&trustservercertificate=true"

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	var books []Book

	db.Raw("SELECT * FROM books").Scan(&books)

	c.JSON(http.StatusOK, books)
}

func getById(c *gin.Context) {

	password := url.QueryEscape("yourStrong(!)Password")
	dsn := "sqlserver://sa:" + password + "@localhost:1433?database=library&encrypt=true&trustservercertificate=true"

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

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
