package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ignaciofalco/goweb-c2-tt/cmd/server/handler"
	"github.com/ignaciofalco/goweb-c2-tt/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()
}
