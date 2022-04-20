package gotesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	a := 10
	b := 4
	resultExpected := 6

	resultFunction := Restar(a, b)

	assert.Equal(t, resultExpected, resultFunction, "Resultado esperado debe ser igual al resultado de la funcion")
}
