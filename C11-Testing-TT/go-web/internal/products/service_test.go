package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	CallReadMethod  bool
	CallWriteMethod bool
	Data            []Product
}

func (mock *MockStore) Read(data interface{}) error {
	products := data.(*[]Product)
	*products = mock.Data
	mock.CallReadMethod = true
	return nil
}

func (mock *MockStore) Write(data interface{}) error {
	mock.CallWriteMethod = true
	mock.Data = data.([]Product)
	return nil
}

func TestUpdate(t *testing.T) {
	p := Product{1, "Test 1", "Color", 0.0, 0.0, "SDJF454", false, "21/04/2022"}
	mock := MockStore{false, false, []Product{p}}
	repo := NewRepository(&mock)
	service := NewService(repo)

	newp := Product{1, "Test 1 Updated Service", "Color Updated", 1.0, 1.0, "KJDF38", true, "21/04/2022"}

	result, _ := service.Update(newp.Id, newp.Name, newp.Color, newp.Price, newp.Stock, newp.Code, newp.Published)

	assert.True(t, mock.CallWriteMethod)
	assert.True(t, mock.CallReadMethod)
	assert.Equal(t, result, newp, "estos productos deben ser iguales")
}

func TestDelete(t *testing.T) {
	id := 2
	p := Product{1, "Test 1", "Color", 0.0, 0.0, "SDJF454", false, "21/04/2022"}
	mock := MockStore{false, false, []Product{p}}
	repo := NewRepository(&mock)
	service := NewService(repo)

	resultProductExists := service.Delete(p.Id)
	resultProductNotExists := service.Delete(id)

	assert.True(t, mock.CallWriteMethod)
	assert.True(t, mock.CallReadMethod)
	assert.Nil(t, resultProductExists, "el retorno debe ser nulo")
	assert.NotNil(t, resultProductNotExists, "el retorno debe ser un error")
}
