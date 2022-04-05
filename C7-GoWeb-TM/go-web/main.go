package main

import (
	"encoding/json"
	. "fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"unicode"

	"github.com/gin-gonic/gin"
)

// GLOBAL VARIABLES
var GlobalProducts []product
var Token string = "K8384FKHID"

// UTILS
func GenerateID() int {
	if len(GlobalProducts) == 0 {
		return 1
	}
	return GlobalProducts[len(GlobalProducts)-1].Id + 1
}

func StartData() {
	file, errFile := os.ReadFile("products.json")
	if errFile != nil {
		log.Fatal("file with data could not be access:", errFile)
	}
	var products []product
	errJson := json.Unmarshal(file, &products)
	if errJson != nil {
		log.Fatalf("error in format json: %s", errJson)
	}
	GlobalProducts = products
}

func PersistData() error {
	file, err := json.Marshal(GlobalProducts)
	if err != nil {
		return Errorf("data can not be saved")
	}
	os.WriteFile("products.json", file, 0644)
	return nil
}

func validateFields(p product) error {
	if p.Name == "" {
		return Errorf("El campo %s es requerido", "nombre")
	}
	if p.Color == "" {
		return Errorf("El campo %s es requerido", "color")
	}
	if p.Price == 0 {
		return Errorf("El campo %s es requerido", "precio")
	}
	if p.Stock == 0 {
		return Errorf("El campo %s es requerido", "stock")
	}
	if p.Code == "" {
		return Errorf("El campo %s es requerido", "codigo")
	}
	if p.CreatedDate == "" {
		return Errorf("El campo %s es requerido", "fecha_creacion")
	}
	return nil
}

// MODEL
type product struct {
	Id          int     `json:"id"`
	Name        string  `json:"nombre" binding:"required"`
	Color       string  `json:"color" binding:"required"`
	Price       float64 `json:"precio" binding:"required"`
	Stock       float64 `json:"stock" binding:"required"`
	Code        string  `json:"codigo" binding:"required"`
	Published   bool    `json:"publicado" binding:"required"`
	CreatedDate string  `json:"fecha_creacion" binding:"required"`
}

// GetAll
func GetAll(c *gin.Context) {
	queries := c.Request.URL.Query()
	productsMatch := make([]product, 0)

	for _, p := range GlobalProducts {
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Query Param")
	}

	for _, p := range GlobalProducts {
		if p.Id == id {
			c.JSON(200, p)
		}
	}

	c.JSON(http.StatusBadRequest, "Not Found")
}

// POST METHODS
// Save new product
func Save(c *gin.Context) {
	var p product

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errValidate := validateFields(p); errValidate != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": errValidate.Error(),
		})
		return
	}

	if c.GetHeader("token") != Token {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no tiene permisos para realizar la peticion solicitada",
		})
		return
	}

	p.Id = GenerateID()
	BackUpGlobalProducts := GlobalProducts
	GlobalProducts = append(GlobalProducts, p)

	err := PersistData()
	if err != nil {
		GlobalProducts = BackUpGlobalProducts
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, p)
}

func main() {

	// START DATA
	StartData()

	r := gin.Default()

	r.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola Juan!!"})
	})

	r.GET("/productos", GetAll)

	r.GET("/productos/:id", GetOne)

	r.POST("/nuevo", Save)

	r.Run()
}
