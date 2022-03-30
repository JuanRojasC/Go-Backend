package main

import (
	. "fmt"
	"time"
)

// Ejercicio 3 - Calcular Precio
/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).
*/

type Product struct {
	Name     string
	Price    float64
	Quantity float64
}

type Service struct {
	Name    string
	Price   float64
	Minutes int
}

type Maintenance struct {
	Name  string
	Price float64
}

func totalValueProducts(products []Product) {
	var total float64 = 0
	for _, p := range products {
		total += (p.Price * p.Quantity)
	}
	Printf("Precio total de los productos: %.2f\n", total)
}

func totalValueServices(services []Service) {
	var total float64 = 0
	for _, s := range services {
		middleHours := float64(89) / float64(30)
		existsResidueMinutes := middleHours - float64(int(middleHours))
		if existsResidueMinutes > 0 && middleHours > 0 {
			total += (s.Price * ((middleHours - existsResidueMinutes) + 1))
		} else {
			total += (s.Price * middleHours)
		}
	}
	Printf("Precio total de los servicios: %.2f\n", total)
}

func totalValueMaintenances(maintenances []Maintenance) {
	var total float64 = 0
	for _, m := range maintenances {
		total += m.Price
	}
	Printf("Precio total de los matenimientos: %.2f\n", total)
}

func main() {

	products := []Product{{"p1", 9944.8, 334}, {"p2", 678.65, 1345}, {"p3", 3465.98, 456}}
	services := []Service{{"s1", 2347, 45}, {"s2", 1450, 60}, {"s3", 1000, 177}}
	maintenances := []Maintenance{{"m1", 23474.78}, {"m2", 4785}, {"m3", 7584.7}}

	go totalValueProducts(products)
	go totalValueServices(services)
	go totalValueMaintenances(maintenances)

	time.Sleep(2000 * time.Millisecond)

}
