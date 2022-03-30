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
}

func main() {

	canal := make(chan int)

	go procesarConCanal(1, canal)
	fmt.Println("Termino el programa")

}
