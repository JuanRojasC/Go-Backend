package main

import . "fmt"

// Ejercicio 1 - Red Social
/*
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:
cambiar nombre: me permite cambiar el nombre y apellido.
cambiar edad: me permite cambiar la edad.
cambiar correo: me permite cambiar el correo.
cambiar contraseña: me permite cambiar la contraseña.
*/

type Person struct {
	Name     string
	LastName string
	Age      int
	Email    string
	Password string
}

func changeNameAndLastName(person *Person, name string, lastName string) {
	Printf("\nChanging Name and LastName in: %p\n", person)
	person.Name = name
	person.LastName = lastName
}

func changeAge(person *Person, age int) {
	Printf("\nChanging Age in: %p\n", person)
	person.Age = age
}

func changeEmail(person *Person, email string) {
	Printf("\nChanging Email in: %p\n", person)
	person.Email = email
}

func changePassword(person *Person, password string) {
	Printf("\nChanging Password in: %p\n", person)
	person.Password = password
}

func main() {

	var personPointer *Person
	person := Person{"Juan", "Rojas", 21, "juanrojas@email.com", "12345"}

	personPointer = &person
	Printf("\nMEMORY ADDRESSES\nSave in Pointer: %p\nOriginal of Var %p\n\n", personPointer, &person)

	Printf("Original struct: %v\n", person)

	changeNameAndLastName(personPointer, "Juan D", "Rojas C")
	Printf("New Name and LastName in: %p %s %s\n", personPointer, person.Name, person.LastName)

	changeAge(&person, 22)
	Printf("New Age in: %p %d\n", &person, person.Age)

	changeEmail(personPointer, "juanrojaspointer@email.com")
	Printf("New Email in: %p %s\n", personPointer, person.Email)

	changePassword(&person, "newpassword")
	Printf("New Password in: %p %s\n", personPointer, person.Password)

	Printf("\nFinal struct: %v", person)
}
