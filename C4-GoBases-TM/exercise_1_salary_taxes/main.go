package main

import . "fmt"

// Ejercicio 1 - Impuestos de Salario
/*
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

type MinimumSalaryError struct {
	message string
}

func (mse *MinimumSalaryError) Error() string {
	return mse.message
}

func main() {
	salary1 := 134000
	salary2 := 187600

	if salary1 < 150000 {
		Println(MinimumSalaryError{"error: el salario ingresado no alcanza el minimo imponible"})
	} else {
		Println("Debe pagar impuesto")
	}

	if salary2 < 150000 {
		Println(MinimumSalaryError{"error: el salario ingresado no alcanza el minimo imponible"})
	} else {
		Println("Debe pagar impuesto")
	}
}
