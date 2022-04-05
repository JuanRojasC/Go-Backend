package main

import (
	"encoding/json"
	. "fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"unicode"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id          int     "json:id"
	Name        string  "json:name"
	Color       string  "json:color"
	Price       float64 "json:price"
	Stock       float64 "json:stock"
	Code        string  "json:code"
	Published   bool    "json:published"
	CreatedDate string  "json:fecha_creacion"
}

// Get Products from JSON file
func GetProducts() []product {
	file, errFile := os.ReadFile("products.json")
	if errFile != nil {
		return make([]product, 0)
	}
	var products []product
	errJson := json.Unmarshal(file, &products)
	if errJson != nil {
		return make([]product, 0)
	}
	return products
}

// GetAll
func GetAll(c *gin.Context) {
	products := GetProducts()

	if len(products) == 0 {
		c.JSON(500, gin.H{"error": Errorf("error")})
	}

	queries := c.Request.URL.Query()
	productsMatch := make([]product, 0)

	for _, p := range products {
		if matchValues(p, queries) {
			productsMatch = append(productsMatch, p)
		}
	}

	c.JSON(200, productsMatch)
}

func matchValues(p product, queries map[string][]string) bool {
	for k, v := range queries {
		r := []rune(k)
		r[0] = unicode.ToUpper(r[0])
		k = string(r)
		if !(Sprint(getField(&p, k)) == v[0]) {
			return false
		}
	}
	return true
}

// GET FIELD BY NAME FROM STRUCT
func getField(p *product, field string) interface{} {
	r := reflect.ValueOf(p)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}

// GetOne
func GetOne(c *gin.Context) {
	products := GetProducts()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Query Param")
	}

	for _, p := range products {
		if p.Id == id {
			c.JSON(200, p)
		}
	}

	c.JSON(http.StatusBadRequest, "Not Found")
}

func main() {
	r := gin.Default()

	r.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola Juan!!"})
	})

	r.GET("/productos", GetAll)

	r.GET("/productos/:id", GetOne)

	r.Run(":8082")
}
