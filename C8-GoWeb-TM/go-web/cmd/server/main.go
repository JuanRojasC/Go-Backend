package main

import (
	"github.com/JuanDRojasC/C8-GoWeb-TM/go-web/cmd/server/handler"
	"github.com/JuanDRojasC/C8-GoWeb-TM/go-web/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	repository := products.NewRepository()
	service := products.NewService(repository)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", productHandler.Save())
	pr.GET("/", productHandler.GetAll())
	pr.GET("/:id", productHandler.GetOne())
	pr.PUT("/:id", productHandler.Update())
	pr.PATCH("/:id", productHandler.Patch())
	r.Run()
}
