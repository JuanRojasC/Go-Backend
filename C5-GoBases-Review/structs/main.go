package main

import . "fmt"

// STRUCT <Type data that collect fields of diferents types>
type Person struct {
	name string
	age  int
}

func (p Person) Greeting() {
	Printf("Holas! Spy %s\n", p.name)
}

func (p *Person) setAge(age int) {
	p.age = age
}

func goStruct() {
	p := Person{"Juan", 21}

	p.Greeting()
	p.setAge(22)
}

// INHERITANCE
type LivingBeign struct {
	name string
	age  int
}

type Human struct {
	nacionality string
	LivingBeign
}

type Cat struct {
	race string
	LivingBeign
}

func goInheritance() {
	c := Cat{race: "siames", LivingBeign: LivingBeign{name: "Bigotes", age: 5}}
	Printf("%+v\n", c.name)

	h := Human{nacionality: "Argentino", LivingBeign: LivingBeign{name: "Ignacio", age: 26}}
	Printf("%+v\n", h.age)
}

// INTERFACES AND ABSTRACTION
type Company interface {
	sell()
	buy()
}

type Store struct {
	name   string
	addres string
}

type clotheStore struct {
	Store
	qtyTShirts   int
	priceTShirts float64
	qtyPants     int
	pricePants   float64
}

func (cs *clotheStore) sellPant(qty int) {
	cs.qtyPants -= qty
	Printf("Nuevo stock %d pantalones\n", cs.qtyPants)
}

func (cs *clotheStore) sellTShirt() {
	Printf("Estoy vendiendo %d pantalones", cs.qtyPants)
}

func (cs clotheStore) sell() {
	Printf("Estoy vendiendo %d pantalones", cs.qtyPants)
}

func (cs clotheStore) buy() {
	Printf("Estoy comprando %d nuevas camisetas", cs.qtyTShirts)
}

type candyStore struct {
	Store
	qtyChocolates   int
	priceChocolates float64
	qtyMandMs       int
	priceMandMs     float64
}

func (cs candyStore) sell() {
	Printf("Estoy vendiendo %d chocolates", cs.qtyChocolates)
}

func (cs candyStore) buy() {
	Printf("Estoy comprando %d M&Ms", cs.qtyMandMs)
}

func createCompany(t string) Company {
	switch t {
	case "candy":
		return &candyStore{Store{"Candy Stores", "Candy Street"}, 3450, 100, 1450, 250}
	case "clothe":
		return &clotheStore{Store{"Clothes Stores", "Clothes Street"}, 130, 1360, 74, 2500}
	default:
		return &candyStore{}
	}
}

func goAbstraction() {
	candys := createCompany("candy")
	clothes := createCompany("clothe")

	//Printf("%+v %+v", candys, clothes)

	// type Assertion
	switch candys.(type) {
	case *clotheStore:
		clothes.(*clotheStore).sellPant(10)
		Println("Es una estructura de tipo clotheStore")
	case *candyStore:
		Println("Es una estructura de tipo candyStore")
	default:
		Println("No es ni candyStore ni clotheStore")
	}

	switch clothes.(type) {
	case *clotheStore:
		Println("Es una estructura de tipo clotheStore")
		clothes.(*clotheStore).sellPant(10)
	case *candyStore:
		Println("Es una estructura de tipo candyStore")
	default:
		Println("No es ni candyStore ni clotheStore")
	}

}

// MAIN
func main() {
	//goInheritance()
	// goAbstraction()
	goAbstraction()
}
