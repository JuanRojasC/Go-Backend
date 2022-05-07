package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ignaciofalco/storage-implementation/cmd/server/handler"
	"github.com/ignaciofalco/storage-implementation/internal/products"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("%serror al cargar archivo .env %s\n", "\033[31m", "\033[0m")
	}

	dataSource := fmt.Sprintf("%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DATABASE_NAME"))
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.GET("full/:id", p.GetFullData())
	pr.GET("/:id", p.GetOne())
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}

//go run cmd/server/main.go
