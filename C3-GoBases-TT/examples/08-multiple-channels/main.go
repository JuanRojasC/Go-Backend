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
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go procesarConCanal(i, c)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Termino el programa", <-c)
	}
}
