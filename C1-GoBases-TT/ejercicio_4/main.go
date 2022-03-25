package main

import "fmt"

// Ejercicio 5 - Qué edad tiene...

func main() {
	
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Printf("El empleado Benjamin tiene %d\n", employees["Benjamin"])


	mayoresDe21 := 0
	for _, v := range employees {
		if v > 21 {
			mayoresDe21 ++
		}
	}

	employees["Federico"] = 25

	delete(employees, "Pedro")

	fmt.Printf("Los empleados mayores de 21 son %d\n", mayoresDe21)
	fmt.Println(employees)

}