package main

import (
	"encoding/csv"
	. "fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50
*/

type Product struct {
	Id       int     "json:id"
	Price    float64 "json:price"
	Quantity float64 "json:qty"
}

func main() {
	file, err := os.Open("./C3-GoBases-TM/exercise_1_save_file/fileCSV.csv")
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	products, _ := reader.ReadAll()

	printTable(products)

	defer file.Close()
}

func printTable(items [][]string) {
	idSpaces := 18
	priceSpaces := 12
	var totalPrice float64 = 0
	Printf("ID%sPrecio%sCantidad\n", strings.Repeat(" ", idSpaces-len("ID")-len("Precio")), strings.Repeat(" ", priceSpaces-len("Cantidad")))
	for _, r := range items {
		Printf("%s%s%s%s%s\n", r[0], strings.Repeat(" ", idSpaces-len(r[0])-len(r[1])), r[1], strings.Repeat(" ", priceSpaces-len(r[2])), r[2])
		price, _ := strconv.ParseFloat(r[1], 64)
		qty, _ := strconv.ParseFloat(r[2], 64)
		totalPrice += (price * qty)
	}
	Printf("%s%.2f", strings.Repeat(" ", idSpaces-len(strconv.FormatFloat(totalPrice, 'f', -1, 64))), totalPrice)
}
