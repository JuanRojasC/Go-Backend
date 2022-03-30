package main

import (
	. "fmt"
)

// Ejercicio 3 - Impuestos de Salario
/*
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/

type MinimumSalaryError struct {
	message string
}

func (mse *MinimumSalaryError) Error() string {
	return mse.message
}

func checkSalary(salary int) error {
	if salary < 150000 {
		return Errorf("error: el minimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	} else {
		return nil
	}
}

func main() {
	salary1 := checkSalary(134000)
	salary2 := checkSalary(187600)

	if salary1 != nil {
		Println(salary1)
	} else {
		Println("Debera pagar impuestos")
	}

	if salary2 != nil {
		Println(salary2)
	} else {
		Println("Debera pagar impuestos")
	}
}
