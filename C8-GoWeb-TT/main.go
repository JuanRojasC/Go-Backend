package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GODOTENV
/*
	Sirve para poder cargar variables de environment desde un archivo .env

	go get -u github.com/joho/gofotenv

	always any var recover is a string

	file .env must be in the same level at i am running the program
	go run main. go -> same level to main
	go run cmd/server/main.go -> same level to cmd
*/

func goGodotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	user := os.Getenv("MY_USER")
	password := os.Getenv("MY_PASS")
}

func main() {

}
