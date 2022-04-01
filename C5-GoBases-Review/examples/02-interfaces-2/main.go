package main

import "fmt"

type Persona interface {
	Saludar()
	setEdad(int)
}

type adulto struct {
	nombre               string
	edad                 int
	categoriaMonotributo int
}

func (p adulto) Saludar() {
	fmt.Printf("Hola! Soy %s\n", p.nombre)
}

func (p *adulto) setEdad(edad int) {
	p.edad = edad
}

type ninio struct {
	nombre       string
	edad         int
	salaDeJardin string
}

func (p ninio) Saludar() {
	fmt.Printf("Hola! Soy %s\n", p.nombre)
}

func (p *ninio) setEdad(edad int) {
	p.edad = edad
}

func (p ninio) getSalaDeJardin() string {
	return p.salaDeJardin
}

func newPersona(nombre string, edad int) Persona {
	if edad > 12 {
		return &adulto{nombre: nombre, edad: edad}
	}
	return &ninio{nombre: nombre, edad: edad, salaDeJardin: "verde"}
}

func main() {

	p := newPersona("Nacho", 10)

	// type assertion
	switch p.(type) {
	case *ninio:
		fmt.Println("la persona es un niño")
		n, _ := p.(*ninio)

		fmt.Println(n.salaDeJardin)
	case *adulto:
		//
		fmt.Println("la persona es un adulto")
	}

	p.Saludar()
	p.setEdad(23)

}
