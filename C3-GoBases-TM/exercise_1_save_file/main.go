package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

/* Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable. */

type Product struct {
	Id       int
	Price    float64
	Quantity float64
}

func main() {
	p1 := Product{1, 3456.67, 124}
	p2 := Product{2, 456.6, 2400}
	p3 := Product{2, 34670, 30}

	products := []Product{p1, p2, p3}

	csvFile, err := os.Create("./C3-GoBases-TM/exercise_1_save_file/fileCSV.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(csvFile)
	var data [][]string
	for _, p := range products {
		row := []string{strconv.Itoa(p.Id), strconv.FormatFloat(p.Price, 'f', -1, 64), strconv.FormatFloat(p.Quantity, 'f', -1, 64)}
		data = append(data, row)
	}
	w.WriteAll(data)
	defer csvFile.Close()
}
