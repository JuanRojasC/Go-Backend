package main

import (
	"errors"
	. "fmt"
	"strings"
)

/* Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado. */

func main() {

	qtyDogs, _ := animal("perro")
	Printf("12 perros necesitan %.2f Kg de alimento\n", qtyDogs(12))

	qtyCats, _ := animal("gato")
	Printf("8 gatos necesitan %.2f Kg de alimento\n", qtyCats(8))

	qtyHamsters, _ := animal("hamster")
	Printf("20 hamsters necesitan %.2f Kg de alimento\n", qtyHamsters(20))

	qtySpiders, _ := animal("tarantula")
	Printf("2 tarantulas necesitan %.2f Kg de alimento\n", qtySpiders(2))

	_, error := animal("esdfs")
	Println("Error", error)
}

func animal(animal string) (func(qtyAnimals int) float32, error) {
	switch strings.ToLower(animal) {
	case "perro":
		return qtyDogs, nil
	case "gato":
		return qtyCats, nil
	case "hamster":
		return qtyHamsters, nil
	case "tarantula":
		return qtySpiders, nil
	default:
		return nil, errors.New("Animal not found")
	}
}

func qtyDogs(qty int) float32 {
	return float32(10 * qty)
}

func qtyCats(qty int) float32 {
	return float32(5 * qty)
}

func qtyHamsters(qty int) float32 {
	return 0.25 * float32(qty)
}

func qtySpiders(qty int) float32 {
	return 0.15 * float32(qty)
}
