package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JuanDRojasC/C18-GoNoSQL-TM/go-web/cmd/server/handler"
	"github.com/JuanDRojasC/C18-GoNoSQL-TM/go-web/internal/products"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/JuanDRojasC/C18-GoNoSQL-TM/go-web/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	dynamoDB *dynamodb.DynamoDB
	table    = "Products"
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

	region := "us-east-1"
	endpoint := fmt.Sprintf("http://%s:%s", os.Getenv("DYNAMO_HOST"), os.Getenv("DYNAMO_PORT"))
	cred := credentials.NewStaticCredentials(os.Getenv("DYNAMO_USER"), os.Getenv("DYNAMO_PASSWORD"), "")
	sess, err := session.NewSession(aws.NewConfig().WithEndpoint(endpoint).WithRegion(region).WithCredentials(cred))

	if err != nil {
		log.Fatal(err)
	}

	dynamoDB = dynamodb.New(sess)

	repository := products.NewRepository(dynamoDB, table)
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
	// pr.PATCH("/:id", productHandler.PatchProduct())
	pr.DELETE("/:id", productHandler.Delete())
	if err := r.Run(); err != nil {
		os.Exit(1)
	}
}
