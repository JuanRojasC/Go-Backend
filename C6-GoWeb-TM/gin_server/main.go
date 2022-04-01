package main

import "github.com/gin-gonic/gin"

// SERVER WRB WITH GIN

func main() {

	// Create a router with gin
	router := gin.Default()

	// Capture a request GET "/hello-world"
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// Run our server
	router.Run(":8080")

}
