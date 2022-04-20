package gotesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	s := []int{3, 7, 4, 9, 8, 24, 0, 5, -6, 46, -35}
	ss := []int{-35, -6, 0, 3, 4, 5, 7, 8, 9, 24, 46}
	SortSlice(s)
	assert.IsIncreasing(t, s, "el slice no se encuentra ordenado")
	assert.Equal(t, ss, s, "el slice no se encuentra ordenado")
}
