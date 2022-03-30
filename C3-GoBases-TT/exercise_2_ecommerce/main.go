package main

import . "fmt"

// Ejercicio 2 - Ecommerce
/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/

type User struct {
	Name     string
	LastName string
	Email    string
	Products []product
}

type product struct {
	Name     string
	Price    float64
	Quantity float64
}

func newProduct(name string, price float64) product {
	return product{name, price, 0}
}

func addProduct(user *User, product *product, quantity float64) {
	product.Quantity = quantity
	user.Products = append(user.Products, *product)
}

func deleteProduct(user *User) {
	user.Products = make([]product, 5)
}

func main() {

	user := User{"Juan", "Rojas", "juanrojas@email.com", make([]product, 0)}
	product1 := newProduct("p1", 1345.76)
	product2 := newProduct("p2", 4673)

	Printf("Original User: %v\n", user)

	addProduct(&user, &product1, 15)
	Printf("First Product Added: %v\n", user)

	addProduct(&user, &product2, 27)
	Printf("Second Product Added: %v\n", user)

	Printf("Final User: %v\n", user)
}
