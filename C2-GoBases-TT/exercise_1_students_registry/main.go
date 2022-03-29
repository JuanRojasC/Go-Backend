package main

import (
	. "fmt"
	"time"
)

// STUDENTS REGISTRY
/* Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

type Student struct {
	Name      string
	LastName  string
	DNI       int
	admission time.Time
}

func (s Student) Details() {
	Printf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha: %s\n\n", s.Name, s.LastName, s.DNI, s.admission.Format("2/01/2006"))
}

func main() {

	s1 := Student{"Juan", "Rojas", 18693669, time.Date(2022, time.March, 14, 0, 0, 0, 0, time.UTC)}
	s1.Details()

	s2 := Student{"Brunela", "Cerutti", 173647405, time.Date(2018, time.October, 05, 0, 0, 0, 0, time.UTC)}
	s2.Details()

}
