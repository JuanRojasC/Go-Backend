package server

import (
	"fmt"
	"log"
	"os"

	"github.com/JuanDRojasC/C18-GoNoSQL-TM/go-web/cmd/server/handler"
	"github.com/JuanDRojasC/C18-GoNoSQL-TM/go-web/internal/products"
	"github.com/JuanDRojasC/C18-GoNoSQL-TM/go-web/pkg/store"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetEngine(db store.Store) *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	region := "us-east-1"
	endpoint := fmt.Sprintf("http://%s:%s", os.Getenv("DYNAMO_HOST"), os.Getenv("DYNAMO_PORT"))
	cred := credentials.NewStaticCredentials(os.Getenv("DYNAMO_USER"), os.Getenv("DYNAMO_PASSWORD"), "")
	sess, err := session.NewSession(aws.NewConfig().WithEndpoint(endpoint).WithRegion(region).WithCredentials(cred))

	if err != nil {
		log.Fatal(err)
	}

	dynamoDB := dynamodb.New(sess)
	table := "Products"

	repo := products.NewRepository(dynamoDB, table)
	service := products.NewService(repo)
	handler := handler.NewProductHandler(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.Use(handler.AuthToken())
	pr.GET("/", handler.GetAll())
	pr.GET("/:id", handler.GetOne())
	pr.POST("/", handler.SaveProduct())
	pr.PUT("/:id", handler.UpdateProduct())
	// pr.PATCH("/:id", handler.PatchProduct())
	pr.DELETE("/:id", handler.Delete())

	return r
}
