package main

import (
	"fmt"
	"time"
)

func procesarConCanal(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")

	c <- i
	time.Sleep(1000 * time.Millisecond)
	c <- 3344
	fmt.Println("termina la go routine")
}

func main() {

	canal := make(chan int)

	go procesarConCanal(1, canal)

	recibido := <-canal // recibimos y lo asignamos a una variable
	fmt.Println("Aun no termino el programa ", recibido)
	// fmt.Println("Que pasa si vuelvo a leer de un canal ya cerrado? ", <-canal)

}
