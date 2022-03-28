package main

import (
	"errors"
	. "fmt"
)

// FUNCIONTS

func main() {
	// 5 is an argument
	checkNumber(5)
	checkNumber(-5)
	checkNumber(0)

	// Handle errors
	_, err := divideBetweenZero(2, 0)

	if err != nil {
		Println("Division entre 0 no posible")
	} else {
		Println("Division exitosa")
	}

	// FUNCTION RETURNING A FUNCTION
	fctReturned := functionReturnFunction("")
	result := fctReturned(2)
	Println(result)
}

// If the first letter is Mayus this function migth be used outside
// number is a parameter
// I can make number1, number 2 int <number1 and number2 both are type int>
func checkNumber(number int) {
	if number < 0 {
		Println("EL numero es negativo")
	} else if number > 0 {
		Println("El numero es positivo")
	} else {
		Println("El numero es 0")
	}
}

// RETURN VALUE
func checkNumberReturn(number int) int {
	if number < 0 {
		return -1
	} else if number > 0 {
		return 1
	}
	return 0
}

// ELIPSIS <Infinite Parameters>
// values is in reality an array
// elipsis always must stay at end of parameters
func elipsisFunction(values ...float32) float32 {
	var result float32
	for _, value := range values {
		result += value
	}
	return result
}

// CALLBACK
func functionCallback(callback func(n int) int, number int) int {
	return callback(number)
}

// RETURN A FUNCTION
func functionReturnFunction(value string) func(number int) int {
	return checkNumberReturn
}

// MULTIPLE RETURN
func multipleReturn(v1, v2 float32) (float32, float32, float32, float32) {
	addition := v1 + v2
	sustraction := v1 - v2
	factor := v1 * v2
	var division float32

	if v2 != 0 {
		division = v1 / v2
	}

	return addition, sustraction, factor, division
}

// MULTIPLE RETURN WITH NAME
func multipleReturnWithName(v1, v2 float32) (addition, sustraction, factor, division float32) {
	addition = v1 + v2
	sustraction = v1 - v2
	factor = v1 * v2

	if v2 != 0 {
		division = v1 / v2
	}

	return
}

// ERRORS
func divideBetweenZero(n1 float32, n2 float32) (float32, error) {

	if n2 == 0 {
		return 0, errors.New("El divisor no puede ser cero")
	}

	return n1 / n2, nil
}
