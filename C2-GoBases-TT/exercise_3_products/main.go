package main

import (
	. "fmt"
	"strings"
)

/* Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
 - El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda */

// INTERFACES
type Product interface {
	CalculateCost()
}

type Ecommerce interface {
	Total()
	Add()
}

// PRODUCT
type product struct {
	kind  string
	name  string
	price float64
}

func (p product) CalculateCost() float64 {
	switch strings.ToLower(p.kind) {
	case "small":
		return 0
	case "medium":
		return p.price * 0.03
	case "big":
		return (p.price * 0.06) + 2500
	default:
		return 0
	}
}

func newProduct(name string, price float64, kind string) product {
	return product{kind, name, price}
}

func newStore(products ...product) store {
	return store{products}
}

// STORES
type store struct {
	products []product
}

func (s store) Total() float64 {
	var total float64 = 0
	for _, v := range s.products {
		total += (v.price + v.CalculateCost())
	}
	return total
}

func (s store) Add(p product) {
	s.products = append(s.products, p)
}

func main() {

	p1 := newProduct("MacBook", 300000, "big")
	p2 := newProduct("cable", 9400, "small")
	p3 := newProduct("charger", 56000, "medium")
	p4 := newProduct("iPhone", 189000, "medium")

	store := newStore(p1, p2, p3, p4)

	for _, v := range store.products {
		Printf("%s: %.2f\n", v.name, v.price+v.CalculateCost())
	}

	Printf("Coste total: %.2f", store.Total())
}
