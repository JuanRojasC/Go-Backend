package main

import (
	. "fmt"
)

/* Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones. Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior */

func main() {
	grades := []float64{2.5, 4.6, 5.0, 3.6, 3.8, 4.5, 2.6, 4.4}

	max, msg := requireOperation("max")
	maxGrade := max(grades...)
	Printf("%s\nLa maxima calificacion es %0.2f\n", msg, maxGrade)

	min, msg := requireOperation("min")
	minGrade := min(grades...)
	Printf("%s\nLa minima calificacion es %0.2f\n", msg, minGrade)

	average, msg := requireOperation("avg")
	avgGrades := average(grades...)
	Printf("%s\nEl promedio de las notas es %.2f\n", msg, avgGrades)

	_, msg = requireOperation("asf")
	Printf("Error %s", msg)
}

func requireOperation(operation string) (func(values ...float64) float64, string) {
	switch operation {
	case "max":
		return maxNumber, "Máximum number"
	case "min":
		return minNumber, "Minimum number"
	case "avg":
		return avg, "Average"
	default:
		return nil, "Operation not found"
	}
}

func maxNumber(values ...float64) float64 {
	var max float64 = 0
	for _, n := range values {
		if n > max {
			max = n
		}
	}
	return float64(max)
}

func minNumber(values ...float64) float64 {
	var min float64 = values[0]
	for _, n := range values {
		if n < min {
			min = n
		}
	}
	return float64(min)
}

func avg(values ...float64) float64 {
	var total float64 = 0
	for _, v := range values {
		total += v
	}
	return total / float64(len(values))
}
