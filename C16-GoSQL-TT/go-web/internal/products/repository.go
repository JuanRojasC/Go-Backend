package products

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

var (
	ErrMethodNotImplemented = errors.New("This method has not been implemented")
)

const (
	GetByName   = "SELECT * FROM products WHERE name = ?"
	SaveProduct = "INSERT INTO products (name, color, price, stock, code, published, created_date) VALUES (?,?,?,?,?,?,?)"
)

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

// INTERFACE
type Repository interface {
	GetAll() ([]Product, error)
	GetOne(id int) (Product, error)
	GetByName(name string) ([]Product, error)
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
	db *sql.DB
}

// Returns all data saved
func (r *repository) GetAll() ([]Product, error) {
	return []Product{}, ErrMethodNotImplemented
}

// Return a product with id pass like parameter or an error if it is not found
func (r *repository) GetOne(id int) (Product, error) {
	return Product{}, ErrMethodNotImplemented
}

// Return a product with name pass like parameter or an error if it is not found
func (r *repository) GetByName(name string) ([]Product, error) {
	var ps []Product
	var p Product
	rows, err := r.db.Query(GetByName, name)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&p.Id, &p.Name, &p.Color, &p.Price, &p.Stock, &p.Code, &p.Published, &p.CreatedDate); err != nil {
			log.Fatal(err)
		}
		ps = append(ps, p)
	}
	return ps, nil
}

// Save a new product and return this with the its ID
func (r *repository) Save(name string, color string, price float64, stock float64, code string, published bool) (Product, error) {
	stmt, err := r.db.Prepare(SaveProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	date_creation := time.Now().Format("2006-01-02 15:04:05")
	result, err = stmt.Exec(name, color, price, stock, code, published, date_creation)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	product := Product{int(id), name, color, price, stock, code, published, date_creation}
	return product, nil
}

// Update completely resource and return itself but updated
func (r *repository) Update(id int, name string, color string, price float64, stock float64, code string, published bool) (Product, error) {
	return Product{}, ErrMethodNotImplemented
}

// Update Name field of resource
func (r *repository) UpdateName(id int, newValue string) (Product, error) {
	return Product{}, ErrMethodNotImplemented
}

// Update Price field of resource
func (r *repository) UpdatePrice(id int, newValue float64) (Product, error) {
	return Product{}, ErrMethodNotImplemented
}

// Delete a resource and return a error if can not do it
func (r *repository) Delete(id int) error {
	return ErrMethodNotImplemented
}

// Generate the next id to be use in the presistence system
func (r *repository) LastID() (int, error) {
	return 0, ErrMethodNotImplemented
}

// Check if element with this ID exists in persistence and return its index or an error if not exists
func (r *repository) CheckExistence(id int) (int, error) {
	return 0, ErrMethodNotImplemented
}

// Return a Repository Interface
func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}
