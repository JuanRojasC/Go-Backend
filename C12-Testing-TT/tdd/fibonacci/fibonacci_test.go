package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciRecursive(t *testing.T) {
	value := 6
	resultExpected := 8
	result := CalculateFibonacciRecursive(value)

	assert.Equal(t, resultExpected, result, "los resultados deben ser iguales")
}

func TestFibonnaciRecursiveSequence(t *testing.T) {
	value := 14
	resultExpected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}
	result := GenerateFibonacciRecursive(value)

	assert.Equal(t, resultExpected, result, "las sequencias deben ser iguales")
}

func TestFibonacciIterative(t *testing.T) {
	value := 5
	resultExpected := 5
	result := CalculateFibonacciIterative(value)

	assert.Equal(t, resultExpected, result, "los resultados deben ser iguales")
}

func TestFibonnaciIterativeSequence(t *testing.T) {
	value := 14
	resultExpected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}
	result := GenerateFibonacciIterative(value)

	assert.Equal(t, resultExpected, result, "las sequencias deben ser iguales")
}

// ~ SAME TIME
func BenchmarkFibonacciRecursive(b *testing.B) {
	value := 47
	CalculateFibonacciRecursive(value)
}

func BenchmarkFibonacciIterative(b *testing.B) {
	value := 50000000000
	CalculateFibonacciIterative(value)
}

// ~ SAME TIME
func BenchmarkFibonnaciSequenceRecursive(b *testing.B) {
	value := 45
	GenerateFibonacciRecursive(value)
}

func BenchmarkFibonacciSequenceIterative(b *testing.B) {
	value := 420000000
	GenerateFibonacciIterative(value)
}
