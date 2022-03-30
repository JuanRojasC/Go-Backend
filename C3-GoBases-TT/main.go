package main

import (
	. "fmt"
	"time"
)

// POINTERS
/*
	Type data that can store a memory addres
	& return memory addres of that variable
	* desreference operator give you access a content of that addres o  memory
*/

func pointers() {
	// POINTERS
	var p *int
	Printf("%p\n", p)

	var p2 = new(int)
	Printf("%p\n", p2)

	v := 12
	p = &v
	Printf("%p\n", p)
	Printf("%p\n", &v)

	// DESREFERENCE
	var vd int = 19
	var pd *int

	pd = &vd

	Printf("El puntero p referencia a la direccion de memoria: %v\n", pd)
	Printf("Al desreferenciar el puntero p obtengo el valor: %d\n", *pd)

	// TRY
	x := 25 // Línea 1
	y := &x // Línea 2
	Println("Línea 3:", x)
	Println("Línea 4:", &x)
	Println("Línea 5:", y)
	Println("Línea 6:", *y)
	Println("Línea 7:", &y)
}

// CONCURRENCE AND PARALELISM
/*
Concurrency: Dos tareas en periodos de tiempo intercalados
Paralelism: Dos tareas que se ejecutan exactamente al mismo tiempo
*/

// GO ROTUTINES
/*
	Solucion para implementar concurrencia en GO <No son Threads> estas son gestionadas por GO Runtime y schedules, no por el OS.
	La ejecucion de aquella funcion que la ejecuta no sera bloqueada por la ejecucion de la Goroutine

	RESERVER WORD go
*/

func procesar(i int) {
	Println(i, "Start Goroutine")
	time.Sleep(1000 * time.Millisecond)
	if i == 2 {
		panic("El numero es un 2")
	}
	Println(i, "Finish Goroutine")
}

func concurrency() {
	for i := 0; i < 5; i++ {
		go procesar(i)
	}

	time.Sleep(5000 * time.Millisecond)
	Println("Termino Concurrencia")
}

// CHANNELS
/*
enviar valores a las Go Routines y esperar hasta recibi a dicho valor <async JS>

RESERVED WORD chan
*/
func processChannel(i int, c chan int) {
	Println(i, "Start")
	time.Sleep(1000 * time.Millisecond)
	Println(i, "End")

	// Write in Channel
	c <- i
	c <- i * 33454534767

	// Read From Channel
	// <-c

	// <You can write more code below but the channel just wait until the first write o read>
}

func channels() {
	channel := make(chan int)
	for i := 1; i <= 5; i++ {
		//  We can create a channel with buffer <default 1> but this work for wait until the buffer is full the channel finish
		go processChannel(i, channel)
		receive := <-channel // Programm stop until a value is returned
		idReceive := <-channel
		// <-channel you the can use the read syntax and represent that you just need wait the end of channel process but not its value
		Printf("Channel %d with ID %d finished\n", receive, idReceive)
	}
}

func main() {

	//pointers()
	//concurrency()
	channels()

}
