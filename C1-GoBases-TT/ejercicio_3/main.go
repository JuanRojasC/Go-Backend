package main

import "fmt"

// Ejercicio 3 - A qué meses corresponde

func main() {
	
	numMonth := 6

	switch numMonth {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("Septembre")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Printf("Mes con el numero %d no encontrado", numMonth)
	}

	monthsNumber := map[int]string{
		1: "January",
		2: "February",
		3: "March",
		4: "April",
		5: "May",
		6: "June",
		7: "July",
		8: "August",
		9: "September",
		10: "October",
		11: "November",
		12: "December",
	}

	
	if _, found := monthsNumber[numMonth]; found != false {
		fmt.Println(monthsNumber[numMonth])
	} else {
		fmt.Printf("Mes con el numero %d no encontrado\n", numMonth)
	}

	// En lo personas me gusta mas la opcion del diccionario, a mi percepcion es mas sencilla y facil de leer. Sin embargo siento que el switch tambien es una muy buena opcion.

}