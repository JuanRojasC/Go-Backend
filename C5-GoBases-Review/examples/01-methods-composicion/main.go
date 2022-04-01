package main

import "fmt"

type SerVivo struct {
	nombre string
	edad   int
}

type OtraStruct struct {
	nombre string
}

type Humano struct {
	nacionalidad string
	SerVivo
	OtraStruct
}

type Gato struct {
	raza string
	SerVivo
}

func main() {
	g := Gato{raza: "siames", SerVivo: SerVivo{nombre: "Bigotes"}}
	p := Humano{nacionalidad: "Argentino",
		SerVivo:    SerVivo{nombre: "Ignacio"},
		OtraStruct: OtraStruct{"otro nombre"}}

	fmt.Printf("%+v\n", g)
	fmt.Printf("%+v\n", p)

	fmt.Printf("%s\n", p.SerVivo.nombre)

}
