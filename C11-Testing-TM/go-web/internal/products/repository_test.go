package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Store interface {
	Write(data interface{}) error
	Read(data interface{}) error
}

type StubStore struct {
	path           string
	callReadMethod bool
}

func (s *StubStore) Read(data interface{}) error {
	products := data.(*[]Product)
	*products = []Product{
		{1, "Test 1", "Color", 0.0, 0.0, "SDJF454", false, "21/04/2022"},
		{2, "Before Update", "Color", 0.0, 0.0, "JKJ29040", true, "21/04/2022"},
	}
	s.callReadMethod = true
	return nil
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func NewStore() Store {
	return &StubStore{"", false}
}

func TestRead(t *testing.T) {
	repo := NewRepository(NewStore())
	p, _ := repo.GetAll()

	assert.Equal(t, products, p, "deben ser iguales")
}

func TestUpdateName(t *testing.T) {
	store := NewStore()
	repo := NewRepository(store)
	resultExpected := "Test Update 1"
	p, _ := repo.UpdateName(1, resultExpected)

	assert.True(t, store.(*StubStore).callReadMethod)
	assert.Equal(t, resultExpected, p.Name, "los nombres deben ser iguales")
}
