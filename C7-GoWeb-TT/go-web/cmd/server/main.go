package main

import (
	"github.com/JuanDRojasC/C7-GoWeb-TT/go-web/cmd/server/handler"
	"github.com/JuanDRojasC/C7-GoWeb-TT/go-web/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	repository := products.NewRepository()
	service := products.NewService(repository)
	product := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", product.Save())
	pr.GET("/", product.GetAll())
	pr.GET("/:id", product.GetOne())
	r.Run()
}
