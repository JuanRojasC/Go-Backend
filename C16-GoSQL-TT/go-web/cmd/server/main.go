package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/JuanDRojasC/C16-GoSQL-TT/go-web/cmd/server/handler"
	"github.com/JuanDRojasC/C16-GoSQL-TT/go-web/internal/products"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/JuanDRojasC/C16-GoSQL-TT/go-web/docs"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	StorageDB *sql.DB
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle Products
// @termOfService https://developers.mercadolibre.com.co/es_co/terminos-y-condiciones

// @contac.name API Support
// @contac.url https://developers.mercadolibre.com.co/support

// @license.name Apache 2.0
// @license.url https://apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_TABLE"))
	StorageDB, err := sql.Open("mysql", dataSource)

	if err != nil {
		log.Fatal(err)
	}

	if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}

	// db := store.NewStore(store.FileType, "./products.json")
	repository := products.NewRepository(StorageDB)
	service := products.NewService(repository)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	pr.Use(productHandler.AuthToken())
	pr.POST("/", productHandler.SaveProduct())
	pr.GET("/", productHandler.GetAll())
	pr.GET("/:id", productHandler.GetOne())
	pr.GET("/product", productHandler.GetByName())
	pr.PUT("/:id", productHandler.UpdateProduct())
	pr.PATCH("/:id", productHandler.PatchProduct())
	pr.DELETE("/:id", productHandler.Delete())
	if err := r.Run(); err != nil {
		os.Exit(1)
	}
}
