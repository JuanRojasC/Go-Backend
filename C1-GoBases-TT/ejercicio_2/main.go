package main

import "fmt"

// Ejercicio 2 - Préstamo

func main() {
	
	customer1 := map[string]interface{}{"edad": 20, "empleado": true, "antiguedad": 1, "sueldo": 174000}
	customer2 := map[string]interface{}{"edad": 25, "empleado": false, "antiguedad": 0, "sueldo": 86970}
	customer3 := map[string]interface{}{"edad": 30, "empleado": true, "antiguedad": 4, "sueldo": 265000}

	customers := []map[string]interface{}{customer1, customer2, customer3}

	for i := 0; i < len(customers); i++{
		c := customers[i]
		posible_prestamo := c["edad"].(int) > 22 && c["empleado"].(bool) && c["antiguedad"].(int) > 1
		if posible_prestamo && c["sueldo"].(int) > 100000 {
			fmt.Println("El cliente es un buen candidato a un prestamo sin intereses")
		} else if posible_prestamo {
			fmt.Println("El cliente es un buen candidatoa a un prestamo")
		} else {
			fmt.Println("Lastimosamente no podemos otorgarte un prestamos")
		}
	}

}