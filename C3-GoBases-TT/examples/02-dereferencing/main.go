package main

import "fmt"

func main() {
	var v int = 19
	fmt.Printf("%p\n", &v)
	// Hacemos que el puntero p, referencie la dirección de memoria de la
	// variable v.

	p := &v
	fmt.Printf("El puntero p referencia a la dirección de memoria: %v \n", p)
	fmt.Printf("Al desreferenciar el puntero p obtengo el valor: %d \n", *p)

}
