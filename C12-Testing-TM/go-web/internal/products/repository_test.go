package products

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockRepository struct {
	CallReadMethod  bool
	CallWriteMethod bool
	ErrorRead       error
	ErrorWrite      error
	Data            []Product
}

func (mock *MockRepository) Read(data interface{}) error {
	if mock.ErrorRead != nil {
		return mock.ErrorRead
	}
	products := data.(*[]Product)
	*products = mock.Data
	mock.CallReadMethod = true
	return nil
}

func (mock *MockRepository) Write(data interface{}) error {
	if mock.ErrorWrite != nil {
		return mock.ErrorWrite
	}
	mock.CallWriteMethod = true
	mock.Data = data.([]Product)
	return nil
}

func NewMock(errRead error, errWrite error) MockRepository {
	return MockRepository{
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

func TestGetAllRepository(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	p, err := repo.GetAll()

	assert.Equal(t, products, p, "deben ser iguales")
	assert.Nil(t, err, "el error debe ser nulo")
}

func TestNegativeGetAllRepository(t *testing.T) {
	mock := NewMock(errors.New("error to read products"), nil)
	repo := NewRepository(&mock)

	resultExpected := []Product{}
	ps, err := repo.GetAll()

	assert.False(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, ps, "debe ser un slice vacio de productos")
	assert.Error(t, err)
}

func TestGetOneRepository(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	p, err := repo.GetOne(mock.Data[0].Id)

	assert.True(t, mock.CallReadMethod)
	assert.Equal(t, mock.Data[0], p, "los productos deben ser iguales")
	assert.Nil(t, err, "el error debe ser nulo")
}

func TestNegative1GetOneRepository(t *testing.T) {
	mock := NewMock(errors.New("error to read product"), nil)
	repo := NewRepository(&mock)

	resultExpected := Product{}
	p, err := repo.GetOne(mock.Data[0].Id)

	assert.False(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, p, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestNegative2GetOneRepository(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)

	resultExpected := Product{}
	p, err := repo.GetOne(786231)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, p, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestNegative3GetOneRepository(t *testing.T) {
	mock := NewMock(errors.New("error to read product"), nil)
	repo := NewRepository(&mock)

	resultExpected := Product{}
	p, err := repo.GetOne(1000)

	assert.False(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, p, "debe ser un producto vacio porque no existe")
	assert.Error(t, err)
}

func TestSaveRepository(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	newp := Product{0, "Test Save", "color test", 3243.76, 345, "JSD839KD", false, time.Now().Format(time.RFC822)}

	p, err := repo.Save(newp.Name, newp.Color, newp.Price, newp.Stock, newp.Code, newp.Published)

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

func TestNegative1SaveRepository(t *testing.T) {
	mock := NewMock(errors.New("error to read product"), nil)
	repo := NewRepository(&mock)
	newp := Product{0, "Test Save", "color test", 3243.76, 345, "JSD839KD", false, time.Now().Format(time.RFC822)}

	resultExpected := Product{}
	p, err := repo.Save(newp.Name, newp.Color, newp.Price, newp.Stock, newp.Code, newp.Published)

	assert.False(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, p, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestNegative2SaveRepository(t *testing.T) {
	mock := NewMock(nil, errors.New("error to write product"))
	repo := NewRepository(&mock)
	newp := Product{0, "Test Save", "color test", 3243.76, 345, "JSD839KD", false, time.Now().Format(time.RFC822)}

	resultExpected := Product{}
	p, err := repo.Save(newp.Name, newp.Color, newp.Price, newp.Stock, newp.Code, newp.Published)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, p, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestUpdateRepository(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	product := mock.Data[1]
	newColor := "Test Update Color"
	newCode := "Test Update Code"
	newCreatedDate := time.Now().Format(time.RFC822)
	product.Color = newColor
	product.Code = newCode
	product.CreatedDate = newCreatedDate

	p, err := repo.Update(product.Id, product.Name, product.Color, product.Price, product.Stock, product.Code, product.Published)

	assert.True(t, mock.CallReadMethod)
	assert.True(t, mock.CallWriteMethod)
	assert.Equal(t, product.Name, p.Name, "el nombre de los productos deben ser iguales")
	assert.Equal(t, newColor, p.Color, "el color de los productos deben ser iguales")
	assert.Equal(t, product.Price, p.Price, "el precio de los productos deben ser iguales")
	assert.Equal(t, product.Stock, p.Stock, "el stock de los productos deben ser iguales")
	assert.Equal(t, newCode, p.Code, "el codigo de los productos deben ser iguales")
	assert.Equal(t, product.Published, p.Published, "ambos productos deben tener el mismo status en publicados")
	assert.NotEqual(t, newCreatedDate, p.CreatedDate, "la fecha de creacion de los productos deben ser iguales")
	assert.Nil(t, err, "el error debe ser nulo")
}

func TestNegative1UpdateRepository(t *testing.T) {
	mock := NewMock(errors.New("error to read product"), nil)
	repo := NewRepository(&mock)
	pu := Product{}

	resultExpected := Product{}
	p, err := repo.Update(pu.Id, pu.Name, pu.Color, pu.Price, pu.Stock, pu.Code, pu.Published)

	assert.False(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, p, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestNegative2UpdateRepository(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	pu := Product{}

	resultExpected := Product{}
	p, err := repo.Update(87893279, pu.Name, pu.Color, pu.Price, pu.Stock, pu.Code, pu.Published)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, p, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestNegative3UpdateRepository(t *testing.T) {
	mock := NewMock(nil, errors.New("error to write product"))
	repo := NewRepository(&mock)
	pu := mock.Data[1]

	resultExpected := Product{}
	p, err := repo.Update(pu.Id, pu.Name, pu.Color, pu.Price, pu.Stock, pu.Code, pu.Published)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, p, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestUpdateNameRepository(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	resultExpected := "Test Update 1"
	p, _ := repo.UpdateName(mock.Data[0].Id, resultExpected)

	assert.True(t, mock.CallReadMethod)
	assert.True(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, p.Name, "los nombres deben ser iguales")
}

func TestNegative1UpdateNameRepository(t *testing.T) {
	mock := NewMock(errors.New("error to read product"), nil)
	repo := NewRepository(&mock)
	newName := "Test Update 1"
	resultExpected := Product{}
	pu, err := repo.UpdateName(mock.Data[0].Id, newName)

	assert.False(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, pu, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestNegative2UpdateNameRepository(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	newName := "Test Update 1"
	resultExpected := Product{}
	pu, err := repo.UpdateName(12398213, newName)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, pu, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestNegative3UpdateNameRepository(t *testing.T) {
	mock := NewMock(nil, errors.New("error to write product"))
	repo := NewRepository(&mock)
	newName := "Test Update 1"
	resultExpected := Product{}
	pu, err := repo.UpdateName(mock.Data[0].Id, newName)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, pu, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestUpdatePriceRepository(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	newPrice := 123456789.987654321
	p, _ := repo.UpdatePrice(mock.Data[0].Id, newPrice)

	assert.True(t, mock.CallReadMethod)
	assert.True(t, mock.CallWriteMethod)
	assert.Equal(t, newPrice, p.Price, "los precios deben ser iguales")
}

func TestNegative1UpdatePriceRepository(t *testing.T) {
	mock := NewMock(errors.New("error to read product"), nil)
	repo := NewRepository(&mock)
	newPrice := 123456789.987654321
	resultExpected := Product{}
	pu, err := repo.UpdatePrice(mock.Data[0].Id, newPrice)

	assert.False(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, pu, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestNegative2UpdatePriceRepository(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	newPrice := 123456789.987654321
	resultExpected := Product{}
	pu, err := repo.UpdatePrice(987429387, newPrice)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, pu, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestNegative3UpdatePriceRepository(t *testing.T) {
	mock := NewMock(nil, errors.New("error to write product"))
	repo := NewRepository(&mock)
	newPrice := 123456789.987654321
	resultExpected := Product{}
	pu, err := repo.UpdatePrice(mock.Data[0].Id, newPrice)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Equal(t, resultExpected, pu, "debe ser un producto vacio")
	assert.Error(t, err)
}

func TestDeleteRepository(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	idToDelete := mock.Data[0].Id
	err := repo.Delete(idToDelete)

	assert.True(t, mock.CallReadMethod)
	assert.True(t, mock.CallWriteMethod)
	assert.Nil(t, err)
}

func TestNegative1Deleteepository(t *testing.T) {
	mock := NewMock(errors.New("error to read product"), nil)
	repo := NewRepository(&mock)
	idToDelete := mock.Data[0].Id
	err := repo.Delete(idToDelete)

	assert.False(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Error(t, err)
}

func TestNegative2Deleteepository(t *testing.T) {
	mock := NewMock(nil, errors.New("error to write product"))
	repo := NewRepository(&mock)
	idToDelete := mock.Data[0].Id
	err := repo.Delete(idToDelete)

	assert.True(t, mock.CallReadMethod)
	assert.False(t, mock.CallWriteMethod)
	assert.Error(t, err)
}

func TestLastID(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	idExpected := len(mock.Data) + 1

	result, err := repo.LastID()

	assert.True(t, mock.CallReadMethod)
	assert.Nil(t, err)
	assert.Equal(t, idExpected, result, "el id debe ser el len del slice de productos mas 1")
}

func TestNegative1LastID(t *testing.T) {
	mock := NewMock(errors.New("error to read product"), nil)
	repo := NewRepository(&mock)
	idExpected := 0

	result, err := repo.LastID()

	assert.False(t, mock.CallReadMethod)
	assert.Error(t, err)
	assert.Equal(t, idExpected, result, "el id debe ser 0")
}

func TestNegative2LastID(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	idExpected := 1
	dataExpected := []Product{}
	mock.Data = dataExpected

	result, err := repo.LastID()

	assert.True(t, mock.CallReadMethod)
	assert.Nil(t, err)
	assert.Equal(t, idExpected, result, "el id debe ser 0")
}

func TestCheckExistence(t *testing.T) {
	mock := NewMock(nil, nil)
	repo := NewRepository(&mock)
	indexExpected := 0
	id := mock.Data[indexExpected].Id

	result, err := repo.CheckExistence(id)

	assert.True(t, mock.CallReadMethod)
	assert.Nil(t, err)
	assert.Equal(t, indexExpected, result, "el indice retornada debe ser igual al indice usado en el test")
}

func TestNegativeCheckExistence(t *testing.T) {
	mock := NewMock(errors.New("error to read product"), nil)
	repo := NewRepository(&mock)
	indexExpected := 0
	id := mock.Data[indexExpected].Id

	result, err := repo.CheckExistence(id)

	assert.False(t, mock.CallReadMethod)
	assert.Error(t, err)
	assert.Equal(t, indexExpected, result, "el indice retornado es 0 porque hubo un error")
}
