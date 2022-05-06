package server

import (
	"github.com/JuanDRojasC/C17-GoSQL-TM/go-web/cmd/server/handler"
	"github.com/JuanDRojasC/C17-GoSQL-TM/go-web/internal/products"
	"github.com/JuanDRojasC/C17-GoSQL-TM/go-web/pkg/store"
	"github.com/gin-gonic/gin"
)

func GetEngine(db store.Store) *gin.Engine {
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	handler := handler.NewProductHandler(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.Use(handler.AuthToken())
	pr.GET("/", handler.GetAll())
	pr.GET("/:id", handler.GetOne())
	pr.POST("/", handler.SaveProduct())
	pr.PUT("/:id", handler.UpdateProduct())
	pr.PATCH("/:id", handler.PatchProduct())
	pr.DELETE("/:id", handler.Delete())

	return r
}
