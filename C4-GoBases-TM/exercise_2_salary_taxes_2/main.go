package main

import (
	"errors"
	. "fmt"
)

// Ejercicio 2 - Impuestos de Salario
/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/

func main() {
	salary1 := 134000
	salary2 := 187600

	if salary1 < 150000 {
		Println(errors.New("error: el salario ingresado no alcanza el minimo imponible"))
	} else {
		Println("Debe pagar impuesto")
	}

	if salary2 < 150000 {
		Println(errors.New("error: el salario ingresado no alcanza el minimo imponible"))
	} else {
		Println("Debe pagar impuesto")
	}
}
