package products

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"nombre"`
	Color       string  `json:"color"`
	Price       float64 `json:"precio"`
	Stock       float64 `json:"stock"`
	Code        string  `json:"codigo"`
	Published   bool    `json:"publicado"`
	CreatedDate string  `json:"fecha_creacion"`
}

func itemToProduct(input map[string]*dynamodb.AttributeValue) (Product, error) {
	var item Product
	err := dynamodbattribute.UnmarshalMap(input, &item)
	if err != nil {
		return Product{}, err
	}
	return item, nil
}
