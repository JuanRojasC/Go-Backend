package products

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type Product struct {
	ID              int     `json:"id"`
	Name            string  `json:"nombre"`
	Type            string  `json:"tipo"`
	Count           int     `json:"cantidad"`
	Price           float64 `json:"precio"`
	Warehouse       string  `json:"warehouse,omitempty"`
	WarehouseAdress string  `json:"warehouse_address,omitempty"`
}

type Repository interface {
	GetOne(id int) (Product, error)
	GetOneWithContext(ctx context.Context, id int) (Product, error)

	GetAll() ([]Product, error)
	GetFullData(id int) (Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

var ErrNotImplemented = fmt.Errorf("not implemented")

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetOne(id int) (Product, error) {
	var product Product
	rows, err := r.db.Query("select id, name, type, count, price from products where id = ?", id)
	if err != nil {
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			return product, err
		}
	}
	return product, nil
}

func (r *repository) GetOneWithContext(ctx context.Context, id int) (Product, error) {
	var product Product
	rows, err := r.db.QueryContext(ctx, "select id, name, type, count, price from products where id = ?", id)

	if err != nil {
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			return product, err
		}
	}
	return product, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product

	rows, err := r.db.Query("select id, name, type, count, price from products")
	if err != nil {
		return ps, err
	}
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			return []Product{}, err
		}
		ps = append(ps, product)
	}
	return ps, nil

}

func (r *repository) GetFullData(id int) (Product, error) {
	var product Product
	innerJoin := "SELECT products.id, products.name, products.type, products.count, products.price, warehouses.name, warehouses.adress " +
		"FROM products INNER JOIN warehouses ON products.id_warehouse = warehouses.id " +
		"WHERE products.id = ?"
	rows, err := r.db.Query(innerJoin, id)
	if err != nil {
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.Warehouse,
			&product.WarehouseAdress); err != nil {
			return product, err
		}
	}
	return product, nil
}

func (r *repository) LastID() (int, error) {

	return 0, ErrNotImplemented

}

func (r *repository) Store(nombre, tipo string, cantidad int, precio float64) (Product, error) {
	stmt, err := r.db.Prepare("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )") // se prepara el SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	p := Product{Name: nombre, Type: tipo, Count: cantidad, Price: precio}
	result, err := stmt.Exec(nombre, tipo, cantidad, precio) // retorna un sql.Result y un error
	if err != nil {
		return Product{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Result devuelto en la ejecución obtenemos el Id insertado
	p.ID = int(insertedId)

	return p, nil

}

func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	stmt, err := r.db.Prepare("UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?") // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	product := Product{ID: id, Name: name, Type: productType, Count: count, Price: price}
	_, err = stmt.Exec(name, productType, count, price, id)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	return Product{}, ErrNotImplemented
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM products WHERE id = ?") // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria

	_, err = stmt.Exec(id) // retorna un sql.Result y un error

	if err != nil {
		return err
	}

	return nil
}
