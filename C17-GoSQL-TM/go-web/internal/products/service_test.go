package products

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockService struct {
	CallReadMethod  bool
	CallWriteMethod bool
	ErrorRead       error
	ErrorWrite      error
	Data            []Product
}

func (mock *MockService) Read(data interface{}) error {
	if mock.ErrorRead != nil {
		return mock.ErrorRead
	}
	products := data.(*[]Product)
	*products = mock.Data
	mock.CallReadMethod = true
	return nil
}

func (mock *MockService) Write(data interface{}) error {
	if mock.ErrorWrite != nil {
		return mock.ErrorWrite
	}
	mock.CallWriteMethod = true
	mock.Data = data.([]Product)
	return nil
}

func NewMockService(errRead error, errWrite error) MockService {
	return MockService{
		CallReadMethod:  false,
		CallWriteMethod: false,
		ErrorRead:       errRead,
		ErrorWrite:      errWrite,
		Data: []Product{
			{1, "MacBook Pro", "gray", 989, 134, "8977JNFJKD", true, "2022-09-15"},
			{2, "iPhone", "gray", 1145, 67, "949JJ54LF", true, "2021-10-01"},
			{3, "Magic Mouse", "white", 79.99, 96, "JFOVN405", true, "2018-05-01"},
			{4, "EarPods", "green update", 5456.76, 345, "OFJ30585KF", true, "2021-04-22"},
		},
	}
}

func TestGetAllService(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	service := NewService(repo)

	ps, err := service.GetAll()

	assert.True(t, mock.CallReadMethod)
	assert.Equal(t, mock.Data, ps, "estos productos deben ser iguales")
	assert.Nil(t, err)
}

func TestNegative1GetAllService(t *testing.T) {
	mock := NewMock(errors.New("error to read products"), nil)
	repo := NewRepository(&mock)
	service := NewService(repo)
	ps, err := service.GetAll()

	assert.False(t, mock.CallReadMethod)
	assert.Nil(t, ps)
	assert.Error(t, err)
}

func TestGetOneService(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	service := NewService(repo)
	pExpected := mock.Data[0]
	id := pExpected.Id

	p, err := service.GetOne(id)

	assert.True(t, mock.CallReadMethod)
	assert.Equal(t, pExpected, p, "estos productos deben ser iguales")
	assert.Nil(t, err)
}

func TestNegative1GetOneService(t *testing.T) {
	mock := NewMock(errors.New("error to read products"), nil)
	repo := NewRepository(&mock)
	service := NewService(repo)
	pExpected := Product{}
	id := pExpected.Id

	p, err := service.GetOne(id)

	assert.False(t, mock.CallReadMethod)
	assert.Equal(t, pExpected, p, "estos productos deben ser iguales")
	assert.Error(t, err)
}

func TestSaveService(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	service := NewService(repo)
	newp := Product{1, "Test Service", "Test Service Color", 1.0, 1.0, "KJDF38", true, time.Now().Format(time.RFC822)}

	p, err := service.Save(newp.Name, newp.Color, newp.Price, newp.Stock, newp.Code, newp.Published)

	assert.True(t, mock.CallReadMethod)
	assert.True(t, mock.CallWriteMethod)
	assert.Equal(t, newp.Name, p.Name, "el nombre de los productos deben ser iguales")
	assert.Equal(t, newp.Color, p.Color, "el color de los productos deben ser iguales")
	assert.Equal(t, newp.Price, p.Price, "el precio de los productos deben ser iguales")
	assert.Equal(t, newp.Stock, p.Stock, "el stock de los productos deben ser iguales")
	assert.Equal(t, newp.Code, p.Code, "el codigo de los productos deben ser iguales")
	assert.Equal(t, newp.Published, p.Published, "ambos productos deben tener el mismo status en publicados")
	assert.Equal(t, newp.CreatedDate, p.CreatedDate, "la fecha de creacion de los productos deben ser iguales")
	assert.Nil(t, err, "el error debe ser nulo")
}

func TestNegative1SaveService(t *testing.T) {
	mock := NewMock(nil, errors.New("error to write products"))
	repo := NewRepository(&mock)
	service := NewService(repo)
	newp := Product{}

	p, err := service.Save(newp.Name, newp.Color, newp.Price, newp.Stock, newp.Code, newp.Published)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, newp, p, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestUpdateService(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	service := NewService(repo)

	newp := Product{1, "Test 1 Updated Service", "Color Updated", 1.0, 1.0, "KJDF38", true, "2021-09-15"}

	p, err := service.Update(newp.Id, newp.Name, newp.Color, newp.Price, newp.Stock, newp.Code, newp.Published)

	assert.True(t, mock.CallWriteMethod)
	assert.True(t, mock.CallReadMethod)
	assert.Equal(t, newp.Name, p.Name, "el nombre de los productos deben ser iguales")
	assert.Equal(t, newp.Color, p.Color, "el color de los productos deben ser iguales")
	assert.Equal(t, newp.Price, p.Price, "el precio de los productos deben ser iguales")
	assert.Equal(t, newp.Stock, p.Stock, "el stock de los productos deben ser iguales")
	assert.Equal(t, newp.Code, p.Code, "el codigo de los productos deben ser iguales")
	assert.Equal(t, newp.Published, p.Published, "ambos productos deben tener el mismo status en publicados")
	assert.NotEqual(t, newp.CreatedDate, p.CreatedDate, "la fecha de creacion de los productos deben ser iguales")
	assert.Nil(t, err)
}

func TestNegative1UpdateService(t *testing.T) {
	mock := NewMock(errors.New("error to read products"), nil)
	repo := NewRepository(&mock)
	service := NewService(repo)

	newp := Product{}

	p, err := service.Update(newp.Id, newp.Name, newp.Color, newp.Price, newp.Stock, newp.Code, newp.Published)

	assert.False(t, mock.CallWriteMethod)
	assert.False(t, mock.CallReadMethod)
	assert.Equal(t, newp.Name, p.Name, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestUpdateNameService(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	service := NewService(repo)
	id := mock.Data[0].Id
	newName := "Test Update Name"

	p, err := service.UpdateName(id, newName)

	assert.True(t, mock.CallWriteMethod)
	assert.True(t, mock.CallReadMethod)
	assert.Equal(t, newName, p.Name, "el nombre de los productos deben ser iguales")
	assert.Nil(t, err)
}

func TestNegative1UpdateNameService(t *testing.T) {
	mock := NewMock(nil, errors.New("error to write products"))
	repo := NewRepository(&mock)
	service := NewService(repo)
	id := mock.Data[0].Id
	newName := "Test Update Name"

	resultExpected := Product{}
	p, err := service.UpdateName(id, newName)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, p, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestUpdatePriceService(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	service := NewService(repo)
	id := mock.Data[0].Id
	newPrice := 12345.6789

	p, err := service.UpdatePrice(id, newPrice)

	assert.True(t, mock.CallWriteMethod)
	assert.True(t, mock.CallReadMethod)
	assert.Equal(t, newPrice, p.Price, "el nombre de los productos deben ser iguales")
	assert.Nil(t, err)
}

func TestNegative1UpdatePriceService(t *testing.T) {
	mock := NewMock(nil, errors.New("error to write products"))
	repo := NewRepository(&mock)
	service := NewService(repo)
	id := mock.Data[0].Id
	newPrice := 12345.6789

	resultExpected := Product{}
	p, err := service.UpdatePrice(id, newPrice)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, p, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestDeleteService(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	service := NewService(repo)
	id := mock.Data[0].Id

	resultProductExists := service.Delete(id)
	resultProductNotExists := service.Delete(id)

	assert.True(t, mock.CallWriteMethod)
	assert.True(t, mock.CallReadMethod)
	assert.Nil(t, resultProductExists, "el retorno debe ser nulo")
	assert.NotNil(t, resultProductNotExists, "el retorno debe ser un error")
}
