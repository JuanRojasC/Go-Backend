package main

import (
	"encoding/json"
	. "fmt"
	"math"
	"reflect"
)

// STRUCTURES AND INTERFACES

// STRUCTURES
// Is a colection of data fields with several data types
type Person struct {
	Name   string
	Genre  string
	Age    int
	Job    string
	Weigth float64
	Tastes Tastes
}

type Tastes struct {
	Food   string
	Movies string
	Series string
	Animes string
	Sports string
}

// STRUCTURES TAGS
type myStrcture struct {
	field1 string `myTag:"value"`
	field2 string `myTag:"value"`
	field3 string `myTag:"value"`
}

type People struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Addres    string `json:"addres"`
}

// METHODS
type Circle struct {
	radio float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radio
}

// POINTERS
func (c *Circle) setRadio(r float64) {
	c.radio = r
}

// INHERITANCE ~
type Vehicle struct {
	Time float64
	Km   float64
}

func (v Vehicle) detail() {
	Printf("%+v", v)
}

type Auto struct {
	v Vehicle
}

func (a *Auto) Details() {
	a.v.detail()
}

func (a *Auto) run(minutes int) {
	a.v.Time = float64(minutes) / 60
	a.v.Km = a.v.Time * 100
}

// INTERFACES
type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return r.width*2 + r.height*2
}

func main() {
	p1 := Person{"Celeste", "Female", 34, "Engineer", 65.5, Tastes{"pollo", "titanic", "", "", ""}}
	p2 := Person{
		Name:  "Nahuel",
		Genre: "Male",
		Age:   30,
		Job:   "Engineer",
		Tastes: Tastes{
			Food:   "Pollo",
			Movies: "Coco",
			Animes: "Shingeki no kyokin",
		},
	}

	Printf("%+v\n", p1)
	Printf("%+v\n", p2)

	// Dot property access
	Printf("%s", p2.Name)

	var p3 Person
	p3.Name = "Ulises"
	p3.Age = 32
	p3.Tastes.Food = "Pescado"

	// Tags
	pjson := People{"Celeste", "Rodriguez", "3204755", "Calle false 123"}
	miJSON, _ := json.Marshal(pjson)

	Println(string(miJSON))
	// Println(err)

	// REFLECT
	person := People{}
	p := reflect.TypeOf(person)

	Println("Type:", p.Name())
	Println("Kind:", p.Kind())
	Println("QtyAttributes:", p.NumField())
	Println("AttributeForIndex:", p.Field(0))
	Println("AttributeTag:", p.Field(0).Tag.Get("json"))
}
