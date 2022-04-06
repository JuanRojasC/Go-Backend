package products

import (
	"encoding/json"
	. "fmt"
	"log"
	"os"
	"time"
)

// Persistances Struct
type Persistence struct {
	Path string
}

func (p *Persistence) Open() {
	file, errOpen := os.ReadFile(p.Path)
	if errOpen != nil {
		log.Fatal("persistence file can not be open", errOpen)
	}
	if errJSON := json.Unmarshal(file, &products); errJSON != nil {
		log.Fatal("json data can not be unmarshal", errJSON)
	}
}

func (p *Persistence) Close() {
	file, errJSON := json.Marshal(products)
	if errJSON != nil {
		log.Fatal(errJSON)
	}
	os.WriteFile("../../products.json", file, 0644)
}

// Model Product
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

// Contains all data
var products []Product

// Persistence variable
var persistence Persistence = Persistence{"../../products.json"}

// INTERFACE
type Repository interface {
	GetAll() ([]Product, error)
	GetOne(id int) (Product, error)
	Save(name string, color string, price float64, stock float64, code string, published bool) (Product, error)
	Update(id int, name string, color string, price float64, stock float64, code string, published bool) (Product, error)
	UpdateName(id int, newValue string) (Product, error)
	UpdatePrice(id int, newValue float64) (Product, error)
	Delete(id int) error
	CheckExistence(id int) (int, error)
	LastID() (int, error)
}

// STRUCT APPLYING INTERFACE
type repository struct {
}

// Returns all data saved
func (r *repository) GetAll() ([]Product, error) {
	persistence.Open()
	return products, nil
}

// Return a product with id pass like parameter or an error if it is not found
func (r *repository) GetOne(id int) (Product, error) {
	persistence.Open()
	pIndex, err := r.CheckExistence(id)
	if err != nil {
		return Product{}, err
	}
	return products[pIndex], nil
}

// Save a new product and return this with the its ID
func (r *repository) Save(name string, color string, price float64, stock float64, code string, published bool) (Product, error) {
	persistence.Open()
	id, _ := r.LastID()
	newProduct := Product{
		Id:          id,
		Name:        name,
		Color:       color,
		Price:       price,
		Stock:       stock,
		Code:        code,
		Published:   published,
		CreatedDate: time.Now().Format(time.RFC822),
	}
	products = append(products, newProduct)
	persistence.Close()
	return newProduct, nil
}

// Update completely resource and return itself but updated
func (r *repository) Update(id int, name string, color string, price float64, stock float64, code string, published bool) (Product, error) {
	persistence.Open()
	pIndex, err := r.CheckExistence(id)
	if err != nil {
		return Product{}, err
	}
	products[pIndex] = Product{
		Id:          products[pIndex].Id,
		Name:        name,
		Color:       color,
		Price:       price,
		Stock:       stock,
		Code:        code,
		Published:   published,
		CreatedDate: products[pIndex].CreatedDate,
	}
	persistence.Close()
	return products[pIndex], nil
}

// Update Name field of resource
func (r *repository) UpdateName(id int, newValue string) (Product, error) {
	persistence.Open()
	pIndex, err := r.CheckExistence(id)
	if err != nil {
		return Product{}, err
	}
	products[pIndex].Name = newValue
	persistence.Close()
	return products[pIndex], nil
}

// Update Price field of resource
func (r *repository) UpdatePrice(id int, newValue float64) (Product, error) {
	persistence.Open()
	pIndex, err := r.CheckExistence(id)
	if err != nil {
		return Product{}, err
	}
	products[pIndex].Price = newValue
	persistence.Close()
	return products[pIndex], nil
}

// Delete a resource and return a error if can not do it
func (r *repository) Delete(id int) error {
	persistence.Open()
	pIndex, err := r.CheckExistence(id)
	if err != nil {
		return err
	}
	products = append(products[:pIndex], products[pIndex+1:]...)
	persistence.Close()
	return nil
}

// Generate the next id to be use in the presistence system
func (r *repository) LastID() (int, error) {
	if len(products) == 0 {
		return 1, nil
	}
	return products[len(products)-1].Id + 1, nil
}

// Check if element with this ID exists in persistence and return its index or an error if not exists
func (r *repository) CheckExistence(id int) (int, error) {
	for i := range products {
		if products[i].Id == id {
			return i, nil
		}
	}
	return 0, Errorf("product with id %d not found", id)
}

// Return a Repository Interface
func NewRepository() Repository {
	return &repository{}
}
