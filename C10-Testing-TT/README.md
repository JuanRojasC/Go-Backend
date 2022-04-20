# Unit Test

Un test unitario es la forma de probar una parte pequeña de codigo, lo mas atomizada posible, usualmente estos son cada funcion o método.

* Pueden ser ejecutados en cualquier orden
* No requieren de acceso a ningun repositorio de datos
* Deben probar bloques de código atómicos
* Son legibles y fáciles y de comprender
* Los resultados arrojados deben ser claros y legibles

Se considera que el unit testing es parte del proceso de programacion y desarrollo, es decir empezar a programa es empezar a crear test unitarios.

## Package Testing

Libreria nativa de Go, que provee las herramientas necesarias para el diseño e implementacion de test. Tambien se encarga de la ejecución automatizada de los test diseñados.

```txt
mypackage_test.go
```

```bash
go test
go test -v
```

```go
import (
    "testing"
)

func Sumar(a, b int) {
    return a + b
}

func TestSumar(t *testing.T) {
    // Arrange/Given
    a := 5
    b := 3
    resultExpected := 8

    // Action/Then 
    result := Sumar(a, b)

    // Assert/When 
    if result != resultExpected {
        t.Errorf("Funcion sumar() arrojo el resultado = %v, pero el resultado esperado es %v", result, resultExpected)
    }
}
```

## Package Testify

Paqute que facilita el desarrollo y la implementacion de los tests.

```bash
go get -u github.com/stretchr/testify
```

```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
    // Arrange/Given
    a := 5
    b := 3
    resultExpected := 8

    // Action/Then 
    result := Sumar(a, b)

    // Assert/When
    assert.Equal(t, resultExpected, result, "deben ser iguales")
    assert.Nil(t, result) 
}
```
