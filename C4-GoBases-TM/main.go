package main

import (
	"errors"
	. "fmt"
)

// ERRORS
/*
	Is an interface that has the method Error() that return a string

	type error interface {
		Error() string
	}

	FUNCTIONS
	Error()
	errors.New()
	fmat.Errorf()

	PACKAGE errors
	New() Crear errores que solo contienen un mensaje
	Is() Compare an error with a direct value
	As() Check is an error is an specific type
	Unwrap() Unpackage
*/

// CUSTOM
type MyError struct {
	message string
	status  int
}

func (e *MyError) Error() string {
	return Sprintf("%d - %v", e.status, e.message)
}

var statusGreaterThan500 = &MyError{"algo salio mal", 500}

func checkIfStatusGraterThan300(status int) error {
	return Errorf("el error es mayor a 300: %d", status)
}

func testCustomErrors(status int) (int, error) {
	if status >= 500 {
		return 500, statusGreaterThan500
	} else if status >= 400 {
		return status, Errorf("%d - este es un error creado con errors.New()", status)
	} else if status >= 300 {
		return status, Errorf("error mayor o igual a 300: %w", checkIfStatusGraterThan300(status))
	}
	return 200, nil
}

func errorsFunc() {
	status, err := testCustomErrors(500)
	if err != nil {
		if errors.Is(err, statusGreaterThan500) {
			Println("El error es de tipo MyError")
		} else {
			Println("El error no es de tipo MyError")
		}
	}
	status, err = testCustomErrors(400)
	var customError error = Errorf("")
	if err != nil {
		if errors.As(err, &customError) {
			Println("El error es de tipo Errorf")
		} else {
			Println("El error no es de tipo Errorf")
		}
	}
	status, err = testCustomErrors(304)
	if err != nil {
		Println("Error directo -", err)
		Println("Error desempacado -", errors.Unwrap(err))
	}
	status, err = testCustomErrors(200)
	Printf("Status %d, Funciona!", status)
}

func main() {

	errorsFunc()

}
