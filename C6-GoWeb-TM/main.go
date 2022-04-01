package main

import (
	"encoding/json"
	. "fmt"
	"log"
	"net/http"
)

// PACKAGE JSON
/*
	import . "encode/json"
*/

// MARSHALL FUNCTION
/*
	return a slice of bytes and a error, that contains a repsentation in JSON format
	the fields to map at JSON must start with a Capital
*/

// UNMARSHALL FUNCTION
/*
	return a strcut, that contains a repsentation in GO format
*/

type product struct {
	Name      string
	Price     int
	Published bool
}

func goToJSON() {
	p := product{"MacBook Pro", 1500, true}
	jsonData, err := json.Marshal(p)

	if err != nil {
		log.Fatal(err)
	}

	// Write a file JSON
	//os.WriteFile("./JuanDRojasC/C1-GoWeb-TM/myJSON.json", jsonData, 0644)

	// Byte's Slice
	Println(jsonData)

	// String
	Println(string(jsonData))
}

func goJSONToStruct() {
	jsonData := `{"Name":"MacBook Pro","Price":1500,"Published":true}`
	var p interface{}

	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}

	Printf("%+v\n", p)
}

// PACKAGE NET/HTTP

func goNetHTTP() {
	http.HandleFunc("/hola", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	Fprintf(w, "hola\n")
}

// MAIN

func main() {
	goToJSON()
	goJSONToStruct()
}
