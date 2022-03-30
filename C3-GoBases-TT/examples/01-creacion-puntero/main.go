package main

import "fmt"

func main() {
	var p1 *int
	fmt.Printf("%p\n", p1)
	fmt.Printf("%d\n", *p1)

	// Forma 2
	/* var p2 = new(int)
	fmt.Printf("%p\n", p2) */

	//Forma 3 y operador de direccion
	/* var p3 *int
	fmt.Printf("%p\n", p3)
	v := 12
	p3 = &v
	fmt.Printf("p3 %p\n", p3)
	fmt.Printf("&p3 %p\n", &p3)
	fmt.Printf("&v %p\n", &v) */

}
