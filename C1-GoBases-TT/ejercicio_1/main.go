package main

import "fmt"

// Ejercicio 1 - Letras de una palabra

func main() {
	
	word := "bootcamp"
	fmt.Println("Cantidad de Letras:", len(word))

	for l := 0; l < len(word); l++ {
		fmt.Printf("Letra %d: %s\n", l, string(word[l]))
	}
}