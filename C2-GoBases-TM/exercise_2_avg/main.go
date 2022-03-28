package main

import (
	"errors"
	. "fmt"
)

/* Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo */

func main() {

	students := make(map[string][]float32)
	students["Andrew"] = []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	students["Jenn"] = []float32{3.4, 5.0, 4.6, 3.9, 2.6, 0}
	students["Arthur"] = []float32{-4.0, 3.5, 5.0, 4.1, 3.6}

	for k, v := range students {
		avg, err := calculateAverage(v...)
		if avg != 0 {
			Printf("El estudiante %s obtuvo un promedio de %0.2f\n", k, avg)
		} else {
			Printf("Error con el estudiante %s %s\n", k, err)
		}
	}
}

func calculateAverage(grades ...float32) (float32, error) {
	var total float32 = 0
	for _, v := range grades {
		if v < 0 {
			return 0, errors.New(Sprintf("%0.1f a grade cannot be less than Zero", v))
		}
		total += v
	}
	return total / float32(len(grades)), nil
}
