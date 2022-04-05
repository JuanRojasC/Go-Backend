package main

import (
	. "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTT
/*
	Request
	- Method: GET/POST/PUT/DELETE
	- URL: endpoint
	- Version: HTTP Version (1.1)
	- Header: Metada, info about request, are optionals
	- Body: optinal

	Response
	- Version:
	- Status: Integer that indicate the successfully or unsuccesful of request
	- Header:
	- Body:
*/

// GIN CONTEXT
/*
	Permite pasar variables entre middleware, asi como los headesr, parametros, query string parameters, method entre otros.
*/

// POST REQUEST
func postRequest() {
	r := gin.Default()

	r.GET("/hola", func(c *gin.Context) {
		c.String(http.StatusOk, "Hola!")
	})

	r.POST("/hola", func(c *gin.Context) {
		body := c.Request.Body
		header := c.Request.Header
		method := c.Request.Method

		Println("!He recibido algo¡")
		Printf("El metodo es: %s", &method)
		Printf("El contenido del header es: ")

		for key, value := range header {
			Printf("\t\nt%s -> %s\n", key, value)
		}

		Println("\tEl body es un io.ReadCloser:(%s), y para trabajar con el vamos a tener que leerlo luego", body)
		c.String(200, "!Lo recibi¡")
	})
}

// GROUP ROUTES
func groupRoutes() {
	server := gin.Default()
	gopher := server.Group("/gophers")
	{
		gopher.GET("/", func(c *gin.Context) {})
		gopher.GET("/get", func(c *gin.Context) {})
		gopher.GET("/info", func(c *gin.Context) {})
	}

	server.GET("/about", func(c *gin.Context) {})
	server.Run(":8080")
}

// PARAMS
func paramsQuery() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {})
	server.GET("/empleados/:id", func(c *gin.Context) {
		// find employee by id
		// c.Param("id")
		// c.String(200, empoloyee)
	})
	server.Run(":8085")
}

// QUERY STRING
func queryString() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {})
	server.GET("/empleados", func(c *gin.Context) {
		// localhost:8090/empleados?id=644&type=seller
		// find employee by id
		// c.Query("id")
		// c.String(200, empoloyee)
	})
	server.Run(":8090")
}

// READ BODY
func readBody() {
	var dict interface{}
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&dict); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// use dict
	})
}

func main() {
	Println("Start Program")
}
