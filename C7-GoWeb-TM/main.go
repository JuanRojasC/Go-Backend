package main

import (
	"encoding/json"
	. "fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type request struct {
	ID       int
	Nombre   string
	Tipo     string
	Cantidad int
	Precio   float64
}

// SINGLE POST METHOD
func postRequest() {
	r := gin.Default()

	r.POST("/productos", func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastID++
		req.ID = lastID
		products = append(products, req)
		file, err := json.Marshal(products)
		if err != nil {
			c.Abort(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		os.WriteFile("./C7-GoWeb-TM/productos.json", file, 0644)
		c.JSON(http.StatusOK, req)
	})
	r.Run()
}

// GROUP POST METHOD
func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastID++
		req.ID = lastID
		products = append(products, req)
		file, err := json.Marshal(products)
		if err != nil {
			c.Abort(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		os.WriteFile("./C7-GoWeb-TM/productos.json", file, 0644)
		c.JSON(http.StatusOK, req)
	}
}

func postGroupMethod() {
	r := gin.Default()
	pr := r.group("/productos")
	pr.POST("/", Guardar())
	r.Run()
}

// PRESISNTENCE HARDCODED
var products []request
var lastID int

func GetProducts() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})
}

// HEADERS
func goHeaders() {
	r := gin.Default()
	r.POST("/headers", func(c *gin.Context) {
		// GET HEADERS FROM ANY REQUEST
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "hola!",
		})

	})
}

func main() {
	//postRequest()
	//postGroupMethod()
}
