package main

import . "fmt"

// Ejercicio 4 - Impuestos de Salario #4
/*
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
1. Desarrolla las funciones necesarias para permitir a la empresa calcular:
	a) Salario mensual de un trabajador según la cantidad de horas trabajadas.
	  - La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
	  - Dicha función deberá retornar más de un valor (salario calculado y error).
	  - En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de impuesto.
	  - En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
	b) Calcular el medio aguinaldo correspondiente al trabajador
	  - Fórmula de cálculo de aguinaldo: [mejor salario del semestre] / 12 * [meses trabajados en el semestre].
	  - La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número negativo.

2. Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función “main()”.
*/

func monthSalary(hoursWorked, valueHour float64) (float64, error) {
	if hoursWorked < 80 {
		return 0, Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	salary := hoursWorked * valueHour
	if salary >= 150000 {
		salary *= 0.9
	}
	return salary, nil
}

func bonusJob(bestSalary float64, monthsWorked int) (float64, error) {
	if bestSalary < 0 || monthsWorked < 0 {
		return 0, Errorf("error: el mejor salario recibido y los meses trabajados no pueden ser negativos")
	}
	return bestSalary / 12 * float64(monthsWorked), nil
}

func main() {

}
