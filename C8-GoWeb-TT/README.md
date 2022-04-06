# Environment Variables

## Godotenv Package

Sirve para poder cargar variables de environment desde un archivo .env

```bash
go get -u github.com/joho/godotenv
```

```go
import "github.com/joho/godotenv"
```

Siempre retorna el valor de la variable como una string y al momento de ejecutar el archivo *main.go* asegurar de que en el mismo nivel del pwd este el archivo *.env*

```txt
go run main. go -> same level to main file
go run cmd/server/main.go -> same level to cmd dir
```
