package products

import (
	. "fmt"
	"time"
)

// b
type Product struct {
	Id          int
	Name        string
	Color       string
	Price       float64
	Stock       *float64
	Code        string
	Published   *bool
	CreatedDate string
}

// c
var products []Product

// d
type Repository interface {
	GetAll() ([]Product, error)
	GetOne(id int) (Product, error)
	Save(name string, color string, price float64, stock float64, code string, published bool) (Product, error)
	LastID() (int, error)
}

// e
type repository struct {
}

// g
func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

// g
func (r *repository) GetOne(id int) (Product, error) {
	for _, p := range products {
		if p.Id == id {
			return p, nil
		}
	}
	return Product{}, Errorf("product with id %d nnot found", id)
}

// g
func (r *repository) Save(name string, color string, price float64, stock float64, code string, published bool) (Product, error) {
	id, _ := r.LastID()
	newProduct := Product{
		Id:          id,
		Name:        name,
		Color:       color,
		Price:       price,
		Stock:       &stock,
		Code:        code,
		Published:   &published,
		CreatedDate: time.Now().Format(time.RFC822),
	}
	products = append(products, newProduct)
	return newProduct, nil
}

// g
func (r *repository) LastID() (int, error) {
	if len(products) == 0 {
		return 1, nil
	}
	return products[len(products)-1].Id + 1, nil
}

// f
func NewRepository() Repository {
	return &repository{}
}
