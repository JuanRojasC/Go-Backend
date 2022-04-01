package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

func (p Persona) Saludar() {
	fmt.Printf("Hola! Soy %s\n", p.nombre)
}

func (p *Persona) SetEdad(edad int) {
	p.edad = edad
}

func main() {

	p := Persona{nombre: "Nacho", edad: 26}
	p1 := Persona{nombre: "Nico", edad: 22}

	p.Saludar()
	p.SetEdad(23)
	fmt.Println(p.edad)
	fmt.Println(p.edad)
	p1.Saludar()
	fmt.Println(p.edad)

}
