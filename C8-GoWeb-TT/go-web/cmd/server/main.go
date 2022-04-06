package main

import (
	"log"

	"github.com/JuanDRojasC/C8-GoWeb-TT/go-web/cmd/server/handler"
	"github.com/JuanDRojasC/C8-GoWeb-TT/go-web/internal/products"
	"github.com/JuanDRojasC/C8-GoWeb-TT/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db := store.NewStore(store.FileType, "./products.json")
	repository := products.NewRepository(db)
	service := products.NewService(repository)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", productHandler.Save())
	pr.GET("/", productHandler.GetAll())
	pr.GET("/:id", productHandler.GetOne())
	pr.PUT("/:id", productHandler.Update())
	pr.PATCH("/:id", productHandler.Patch())
	pr.DELETE("/:id", productHandler.Delete())
	r.Run()
}
