package main

import (
	. "fmt"
	"strings"
)

/* Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario. */

func main() {

	Printf("El empleado que trabajo %d minutos y es categoria %s ganará: %0.2f\n", 8000, "B", calculateSalary(8000, "B"))
	Printf("El empleado que trabajo %d minutos y es categoria %s ganará: %0.2f\n", 4000, "B", calculateSalary(4000, "B"))
	Printf("El empleado que trabajo %d minutos y es categoria %s ganará: %0.2f\n", 5000, "A", calculateSalary(5000, "A"))
	Printf("El empleado que trabajo %d minutos y es categoria %s ganará: %0.2f\n", 3400, "C", calculateSalary(3400, "C"))

}

func calculateSalary(mins int, category string) float64 {

	var hours float64 = float64(mins) / 60

	switch strings.ToUpper(category) {
	case "A":
		return hours * 1000
	case "B":
		salary := (hours * 1500)
		return salary + (salary * 0.2)
	case "C":
		salary := (hours * 3000)
		return salary + (salary * 0.5)
	default:
		return 0
	}
}
