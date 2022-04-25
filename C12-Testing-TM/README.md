# Golangci & Code Coverage

Linter es una herramienta para realizar el analisis automático y estatico del código. Permite detectar de forma temprana errores y posibles malas prácticas.

* Ayuda a sseguir las buenas prácticas del lenguaje
* Permite detectar errores de forma temprana
* Mejora la calida del código

## golangci-lint

```powershell
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s --
-b $(go env GOPATH)/bin v1.41.1
```

```bash
brew install golangci-lint
brew update golangci-lint
```

```bash
docker run --rm -v $(pwd):/app -w /app golangci/golanci-lint:v1.41.1
golangci-lint run -v
```

```bash
golangci-lint run
```

## Coverage

```bash
go test -cover ./...
```

Métrica que nos permite saber cuánto del código fuente de un sofrware fue testeado

### Coverage Report

```bash
go test -cover -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Reporte que además del Code Coverage también brinda información para saber qué partes del software fueron cubiertos por las pruebas.

## Benchmark

Tipo de test que permite examinar y poner a prueba la eficiencia del software, Go provee las herramientas nativas para poder realizar este tipo de pruebas.

* El archivo benchmark test debe tener el sufijo *_test*
* El nombre de cada funcion debe iniciar con la palabra *Benchmark*
* El paramatro de la funcion es *b \*testting.B*

```bash
go test -bench .
```

```go
func BenckmarkSum256(b *testing.B) {
    data := []byte("Digital House")
    for i := 0; i < b.N; i++ {
        sha1.Sum256(data)
    }
}

func BenckmarkSum(b *testing.B) {
    data := []byte("Digital House")
    for i := 0; i < b.N; i++ {
        sha1.Sum(data)
    }
}
```
