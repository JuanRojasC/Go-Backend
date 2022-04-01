package main

import "fmt"

type animal interface {
	respirar()
	caminar()
}

type leon struct {
	edad int
}

func (l leon) respirar() {
	fmt.Println("leon respira")
}

func (l leon) caminar() {
	fmt.Println("leon camina")
}

type loro struct {
	edad                      int
	color                     string
	cantidadDeVecesQueRespiro int
}

func (l *loro) respirar() {
	fmt.Println("loro respira")
	l.cantidadDeVecesQueRespiro++
}

func (l loro) caminar() {
	fmt.Println("loro camina")
}

func main() {
	var a animal

	a.respirar()
	a.caminar()

	a = &loro{edad: 2}

	// a = &leon{edad: 5}
	a.respirar()
	a.caminar()
}
