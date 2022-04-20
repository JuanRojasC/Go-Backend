package gotesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositivoDividir(t *testing.T) {
	a := 10
	b := 2
	resultExpected := 5

	resultFunction, _ := Dividir(a, b)

	assert.Equal(t, resultExpected, resultFunction, "error en la divison")
}

func TestNegativoDividir(t *testing.T) {
	a := 10
	b := 0
	resultExpected := 0
	resultFunction, err := Dividir(a, b)

	assert.Equal(t, resultExpected, resultFunction, "el resultado debe ser 0")
	assert.NotNil(t, err)
}
