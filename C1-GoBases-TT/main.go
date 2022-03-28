package main

import "fmt"

func main() {
	// Arrays & Slices

	// Array <Static Size>
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Printf(a[0], "\n")
	fmt.Println(a)

	// Slices <Dynamic Size>
	// var s []T tipo de dato variable

	var s = []bool{true, false}
	fmt.Println(s[0])

	// Inicia un slice con 4 posiciones en 0 <3 arg opt, capacity of slice)
	// {0,0,0,0,0}
	b := make([]int, 5)
	fmt.Printf("%d", b[1:4]) // incluye el primero excluye el ultimo

	// longitud de un Slice # de elementos que contiene len(s)
	// capacidad de un Slice # de elementos del array subyacente, contando desde el primer elemento del segmento cap(s)

	// add elements to slice
	b = append(b, 6, 7, 8, 9, 10)

	//Maps <Key, Value>
	// myMap := map[string]int{}
	// myMapMake := make(map[string]string)

	// Declare with elements
	var mapStudents = map[string]int{"Benjamin": 20, "Nahuel": 26}
	fmt.Println(mapStudents["Benjamin"])

	// New Key-Value
	mapStudents["Brenda"] = 19

	// Update Value
	mapStudents["Benjamin"] = 22

	// Delete Key-Values
	delete(mapStudents, "Benjamin")

	// For Maps key
	// For Arr or Slice Index
	for key, element := range mapStudents {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	// FOR
	// Standar For
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}

	// For Range
	frutas := []string{"apple", "banana", "peras"}
	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}

	// Bucle Infinito
	sum1 := 0
	for {
		sum1++
		if sum1 == 10 {
			break
		}
	}

	// Loop While
	i := 1
	for i < 10 {
		i += i
	}
	fmt.Println(i)

	// Break <Break a Loop>
	// Continue <Next iterarion>

	// CONDICIONALES IF - ELSE

	sueldo := 4500
	if sueldo > 3000 {
		fmt.Println("Condicion verdadera")
	} else {
		fmt.Println("Condicion falsa")
	}

	if edad := 20; edad > 50 {

	} else if edad >= 18 {

	} else {

	}

	x := "z"
	switch x {
	case "h":
		fmt.Println("es h")
	case "z":
		fmt.Println("es z")
	default:
		fmt.Println("no encontro caso alguno")
	}

	switch x {
	case "h", "z":
		fmt.Println("es h o z")
	default:
		fmt.Println("no cumplio ningun case")
	}

	switch p := 1; p {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		fmt.Println("Es un numero decimal")
	default:
		fmt.Println("No es un numero decimal")
	}

	// fallthrough
	// Ejecuta el siguiente case sin imporatar si aplica
}
