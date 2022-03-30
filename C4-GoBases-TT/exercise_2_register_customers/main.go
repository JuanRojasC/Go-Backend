package main

import (
	. "fmt"
	"log"
	"math/rand"
	"os"
)

// Ejercicio 2 - Registrando Clientes
/*
El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos para registrar a un cliente son:
 - Legajo
 - Nombre y Apellido
 - DNI
 - Número de teléfono
 - Domicilio

Tarea 1: El número de legajo debe ser asignado o generado por separado y en forma previa a la carga de los restantes gastos. Desarrolla e implementa una función para generar un ID que luego utilizarás para asignarlo como valor a “Legajo”. Si por algún motivo esta función retorna valor “nil”, debe generar un panic que interrumpa la ejecución y aborte.

Tarea 2: Antes de registrar a un cliente, debes verificar si el mismo ya existe. Para ello, necesitas leer los datos de un archivo .txt. En algún lugar de tu código, implementa la función para leer un archivo llamado “customers.txt” (como en el ejercicio anterior, este archivo no existe, por lo que la función que intente leerlo devolverá un error).
Debes manipular adecuadamente ese error como hemos visto hasta aquí. Ese error deberá:
1.- generar un panic;
2.- lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y continuar con la ejecución del programa normalmente.

Tarea 3: Luego de intentar verificar si el cliente a registrar ya existe, desarrolla una función para validar que todos los datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de los valores retornados deberá ser de tipo error para el caso de que se ingrese por parámetro algún valor cero (recuerda los valores cero de cada tipo de dato, ej: 0, “”, nil).

Tarea 4: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes: “Fin de la ejecución”, “Se detectaron varios errores en tiempo de ejecución” y “No han quedado archivos abiertos” (en ese orden). Utiliza defer para cumplir con este requerimiento.

Requerimientos generales:
 - Utiliza recover para recuperar el valor de los panics que puedan surgir (excepto en la tarea 1).
 - Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error (por ejemplo las que intenten leer archivos).
Genera algún error, personalizandolo a tu gusto, utilizando alguna de las funciones que GO provee para ello (realiza también la validación pertinente para el caso de error retornado).

*/

type Customer struct {
	License  int
	Name     string
	LastName string
	DNI      int
	Phone    int
	Addres   string
}

func generateID() int {
	return rand.Intn(84583745937)
}

func readFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o esta dañado")
	}
	return file
}

func checkDataCustomer(c Customer) (bool, error) {
	if c.License == 0 || c.Name == "" || c.LastName == "" || c.DNI == 0 || c.Phone == 0 || c.Addres == "" {
		return false, Errorf("All fields must be different of ZeroValue")
	}
	return true, nil
}

func main() {
	defer Println("No han quedado archivos abiertos")
	defer Println("Se detectaron varios errores en tiempo de ejecución")
	defer Println("Fin de la Ejecucuion")
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	id := generateID()

	if id == 0 {
		panic("ID cannot be nil")
	}

	readFile("customers.txt")

	customer := Customer{id, "Juan", "Rojas", 28469405, 3445975323, "Calle falsa"}
	_, err := checkDataCustomer(customer)
	if err != nil {
		log.Println(err)
	}

}
