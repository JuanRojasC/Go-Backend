package products

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

var (
	ErrMethodNotImplemented = errors.New("method has not been implemented")
	ErrInternalServer       = errors.New("internal server error")
	ErrMalFormatted         = errors.New("entity or struct malformatted")
)

type Field string

const (
	ID          Field = "id"
	NAME        Field = "name"
	COLOR       Field = "color"
	PRICE       Field = "price"
	STOCK       Field = "stock"
	CODE        Field = "code"
	PUBLISHED   Field = "published"
	CREATEDDATE Field = "created_date"

	getByID       string = "SELECT " + string(ID) + "," + string(NAME) + "," + string(COLOR) + "," + string(PRICE) + "," + string(STOCK) + "," + string(CODE) + "," + string(PUBLISHED) + "," + string(CREATEDDATE) + " FROM products WHERE id = ?"
	getByName     string = "SELECT " + string(ID) + "," + string(NAME) + "," + string(COLOR) + "," + string(PRICE) + "," + string(STOCK) + "," + string(CODE) + "," + string(PUBLISHED) + "," + string(CREATEDDATE) + " FROM products WHERE name = ?"
	getAll        string = "SELECT " + string(ID) + "," + string(NAME) + "," + string(COLOR) + "," + string(PRICE) + "," + string(STOCK) + "," + string(CODE) + "," + string(PUBLISHED) + "," + string(CREATEDDATE) + " FROM products"
	saveProduct   string = "INSERT INTO products (" + string(NAME) + "," + string(COLOR) + "," + string(PRICE) + "," + string(STOCK) + "," + string(CODE) + "," + string(PUBLISHED) + "," + string(CREATEDDATE) + ") VALUES (?,?,?,?,?,?,?)"
	updateField   string = "UPDATE products SET %s WHERE id = ?"
	updateProduct string = "UPDATE products SET " + string(NAME) + " = ? ," + string(COLOR) + " = ? ," + string(PRICE) + " = ? ," + string(STOCK) + " = ? ," + string(CODE) + " = ? ," + string(PUBLISHED) + " = ?" + " WHERE id = ?"
)

var FIELDS = map[string]Field{
	"Name":      NAME,
	"Color":     COLOR,
	"Price":     PRICE,
	"Stock":     STOCK,
	"Code":      CODE,
	"Published": PUBLISHED,
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

type ProductField struct {
	Name  Field
	Value interface{}
}

// INTERFACE
type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id int) (Product, error)
	GetByName(name string) ([]Product, error)
	Save(name string, color string, price float64, stock float64, code string, published bool) (Product, error)
	Update(ctx context.Context, id int, name string, color string, price float64, stock float64, code string, published bool) (Product, error)
	Patch(id int, fields ...ProductField) (Product, error)
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
	var ps []Product
	var p Product
	rows, err := r.db.Query(getAll)
	if err != nil {
		return ps, ErrInternalServer
	}
	for rows.Next() {
		if err = rows.Scan(&p.Id, &p.Name, &p.Color, &p.Price, &p.Stock, &p.Code, &p.Published, &p.CreatedDate); err != nil {
			return ps, ErrMalFormatted
		}
		ps = append(ps, p)
	}
	return ps, nil
}

// Return a product with id pass like parameter or an error if it is not found
func (r *repository) GetByID(id int) (Product, error) {
	var p Product
	rows, err := r.db.Query(getByID, id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&p.Id, &p.Name, &p.Color, &p.Price, &p.Stock, &p.Code, &p.Published, &p.CreatedDate); err != nil {
			log.Fatal(err)
		}
	}
	return p, nil
}

// Return a product with name pass like parameter or an error if it is not found
func (r *repository) GetByName(name string) ([]Product, error) {
	var ps []Product
	var p Product
	rows, err := r.db.Query(getByName, name)
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
	stmt, err := r.db.Prepare(saveProduct)
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
func (r *repository) Update(ctx context.Context, id int, name string, color string, price float64, stock float64, code string, published bool) (Product, error) {
	stmt, err := r.db.Prepare(updateProduct)
	if err != nil {
		return Product{}, ErrInternalServer
	}
	_, err = stmt.ExecContext(ctx, name, color, price, stock, code, published, id)
	if err != nil {
		return Product{}, ErrInternalServer
	}
	return r.GetByID(id)
}

// Update fields of resource and return itself but updated
func (r *repository) Patch(id int, fields ...ProductField) (Product, error) {
	var sets string
	var values []interface{}
	for i, f := range fields {
		if len(fields) > 1 && i != (len(fields)-1) {
			sets += " " + string(f.Name) + " = ?, "
		} else {
			sets += " " + string(f.Name) + " = ? "
		}
		values = append(values, f.Value)
	}
	values = append(values, id)
	fmt.Printf(updateField, sets)
	stmt, err := r.db.Prepare(fmt.Sprintf(updateField, sets))
	if err != nil {
		return Product{}, err
	}
	_, err = stmt.Exec(values...)
	if err != nil {
		return Product{}, err
	}
	return r.GetByID(id)
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
