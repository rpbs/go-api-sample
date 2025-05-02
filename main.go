package main

import (
	"github.com/gin-gonic/gin"
)

type Costumers struct {
	Name    string
	Address string
}

func main() {

	r := gin.Default()

	r.GET("/books", getAllBooks)
	r.GET("/books/:id", getById)

	r.Run(":8080")

}
