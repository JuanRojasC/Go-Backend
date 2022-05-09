package products

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

var (
	ErrMethodNotImplemented = errors.New("method has not been implemented")
	ErrInternalServer       = errors.New("internal server error")
	ErrMalFormatted         = errors.New("entity or struct malformatted")
	ErrNotFound             = errors.New("product not found")
)

type Repository interface {
	GetAll(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id int) (Product, error)
	GetByName(name string) ([]Product, error)
	Save(ctx context.Context, name string, color string, price float64, stock float64, code string, published bool) (Product, error)
	Update(ctx context.Context, id int, name string, color string, price float64, stock float64, code string, published bool) (Product, error)
	Delete(ctx context.Context, id int) error
	CheckExistence(ctx context.Context, id int) error
}

// STRUCT APPLYING INTERFACE
type repository struct {
	db    *dynamodb.DynamoDB
	table string
}

// Returns all data saved
func (r *repository) GetAll(ctx context.Context) ([]Product, error) {
	var products []Product
	result, err := r.db.ScanWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(r.table),
	})
	if err != nil {
		return nil, ErrInternalServer
	}
	for _, p := range result.Items {
		pr, err := itemToProduct(p)
		if err != nil {
			continue
		}
		products = append(products, pr)
	}
	return products, nil
}

// Return a product with id pass like parameter or an error if it is not found
func (r *repository) GetByID(ctx context.Context, id int) (Product, error) {
	result, err := r.db.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(fmt.Sprintf("%d", id)),
			},
		},
	})

	if err != nil {
		return Product{}, ErrInternalServer
	}

	if result.Item == nil {
		return Product{}, ErrNotFound
	}

	return itemToProduct(result.Item)
}

// Return a product with name pass like parameter or an error if it is not found
func (r *repository) GetByName(name string) ([]Product, error) {
	return nil, ErrMethodNotImplemented
}

// Save a new product and return this with the its ID
func (r *repository) Save(ctx context.Context, name string, color string, price float64, stock float64, code string, published bool) (Product, error) {
	id := uuid.New().ClockSequence()
	p := Product{id, name, color, price, stock, code, published, time.Now().Format("2006-01-02 15:04:05")}
	av, err := dynamodbattribute.MarshalMap(&p)
	if err != nil {
		return Product{}, err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.table),
	}
	_, err = r.db.PutItemWithContext(ctx, input)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

// Update completely resource and return itself but updated
func (r *repository) Update(ctx context.Context, id int, name string, color string, price float64, stock float64, code string, published bool) (Product, error) {
	p := Product{id, name, color, price, stock, code, published, time.Now().Format("2006-01-02 15:04:05")}
	av, err := dynamodbattribute.MarshalMap(&p)
	if err != nil {
		return Product{}, err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.table),
	}
	_, err = r.db.PutItemWithContext(ctx, input)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

// Update fields of resource and return itself but updated

// Delete a resource and return a error if can not do it
func (r *repository) Delete(ctx context.Context, id int) error {
	if exists := r.CheckExistence(ctx, id); exists != nil {
		return ErrNotFound
	}
	_, err := r.db.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(fmt.Sprintf("%d", id)),
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

// Check if element with this ID exists in persistence and return its index or an error if not exists
func (r *repository) CheckExistence(ctx aws.Context, id int) error {
	result, err := r.db.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(fmt.Sprintf("%d", id)),
			},
		},
	})

	if err != nil {
		return err
	}

	if result.Item == nil {
		return ErrNotFound
	}

	return nil
}

// Return a Repository Interface
func NewRepository(db *dynamodb.DynamoDB, table string) Repository {
	return &repository{db, table}
}
