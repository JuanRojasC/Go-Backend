package main

import "fmt"

// Incrementar recibe un puntero de tipo entero
func Incrementar(v *int) {
	// Desreferenciamos la variable v para obtener
	// su valor e incrementarlo en 1
	*v++

}
func main() {
	var v int = 19
	// La función Increase recibe un puntero
	// utilizamos el operador de dirección &
	// para pasar la dirección de memoria de v
	Incrementar(&v)
	fmt.Println("El valor de v ahora vale:", v)
}
