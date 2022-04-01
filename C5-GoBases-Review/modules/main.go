package main

import (
	. "fmt"

	"github.com/JuanDRojasC/C5-GoBases-Review/modules/person"
)

func main() {
	Println("Hola")
	p := person.NewPerson("juan", 21)
	p.Greeting()
}
