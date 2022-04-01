package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id          int
	Name        string
	Color       string
	Price       float64
	Stock       float64
	Code        string
	Published   bool
	CreatedDate string
}

func GetAll(c *gin.Context) {
	file, errFile := os.ReadFile("products.json")
	if errFile != nil {
		log.Fatal(errFile)
	}
	var products []product
	errJson := json.Unmarshal(file, &products)
	if errJson != nil {
		log.Fatal(errJson)
	}
	c.JSON(200, gin.H{"productos": products})
}

func main() {
	r := gin.Default()

	r.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola Juan!!"})
	})

	r.GET("/productos", GetAll)

	r.Run()
}
