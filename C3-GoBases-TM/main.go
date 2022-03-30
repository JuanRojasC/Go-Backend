package main

// PACKAGES
// Characters escape like \n - \\ - \t - \v (vertical tab)

// verbs like %p memory addres, %v standar etc

// Sprint is a method that allow concate strings

// PACKAGE OS
/*
	ENVIROMENT VARIABLES
	err := .Setenv(key, value string)
	value := os.Getenv("NAME") return a empty string if not exists
	value, ok := os.Lookupenv("NAME") return value and boolean if the var does not exist

	FILES (os)
	os.Exit(1) kill the program with code input
	files, err := os.ReadDir(".") path like arg and return a silece of files and error if not exits
	data, errr := os.ReadFile("PATH") return data in bites and error if can read that file
	err := os.WriteFile("PATH", data []byte, 0644 (permisos)) return error if can write

	READER (io)
	C := strings.NewReader("my string")
	alue, err := io.Copy(os.Stdout, C)

	r := string.NewReader("my string")
	b, err := io.ReadAll(r)

	io.WriteString(os.Stdout, "Hello World")

*/

func main() {

}
