package main

import "fmt"

// Ejercicio 2 - Clima

func main(){

	var temperature int = 50
	var humidity float32 = 65
	var pressure float32 = 29.90

	var output string = `
	Bogota - Colombia
	Temperatura: %d
	Humedad: %0.2f
	Presion: %0.2f
	`

	fmt.Printf(output, temperature, humidity, pressure)
	
}