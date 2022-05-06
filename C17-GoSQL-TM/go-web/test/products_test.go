package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	engine "github.com/JuanDRojasC/C17-GoSQL-TM/go-web/cmd/server/engine"
	"github.com/JuanDRojasC/C17-GoSQL-TM/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"nombre"`
	Color       string  `json:"color"`
	Price       float64 `json:"precio"`
	Stock       float64 `json:"stock"`
	Code        string  `json:"codigo"`
	Published   bool    `json:"publicado"`
	CreatedDate string  `json:"fecha_creacion"`
}

type Response struct {
	Code  int     `json:"status"`
	Data  Product `json:"data,omitempty"`
	Error string  `json:"error,omitempty"`
}

var FILENAME = "products_test.json"
var BACKUPNAME = "products_test_temp.json"

/* Create server for functional test */
func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	createBackup()
	db := store.NewStore(store.FileType, "./"+FILENAME)
	return engine.GetEngine(db)
}

func createBackup() {
	original, err := ioutil.ReadFile("./" + FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(original)
	err = ioutil.WriteFile(BACKUPNAME, original, 0755)
	if err != nil {
		log.Fatal(err)
	}

}

func resetStore() {
	original, err := ioutil.ReadFile("./" + BACKUPNAME)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(FILENAME, original, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("./" + BACKUPNAME)
}

/* Facilita la creacion de la request para los tests */
func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestFunctionalGetAllProducts(t *testing.T) {

	type response struct {
		Code  int       `json:"status"`
		Data  []Product `json:"data,omitempty"`
		Error string    `json:"error,omitempty"`
	}

	// Create server for test
	r := createServer()

	// Create request for test
	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	// Var that store the response
	var res response

	// Indicar el servidor que atendera la request
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.Nil(t, err)
	assert.True(t, len(res.Data) > 0)

	resetStore()
}

func TestFunctionalGetOne(t *testing.T) {

	// Create server for test
	r := createServer()

	// Create request for test
	req, rr := createRequestTest(http.MethodGet, "/products/1", "")

	// Var that store the response
	var res Response

	// Indicar el servidor que atendera la request
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.Nil(t, err)
	assert.Equal(t, 1, res.Data.Id, "id debe ser igual a 1")
	assert.Equal(t, "MacBook Pro", res.Data.Name, "el nombre debe ser igual a MacBook Pro")
	assert.Equal(t, "gray", res.Data.Color, "color debe ser igual a gray")
	assert.Equal(t, 989.0, res.Data.Price, "el precio debe ser igual a 989")
	assert.Equal(t, 134.0, res.Data.Stock, "el stock debe ser igual a 134")
	assert.Equal(t, "8977JNFJKD", res.Data.Code, "id debe ser igual a 8977JNFJKD")
	assert.Equal(t, true, res.Data.Published, "el producto debe estar publicado")
	assert.Equal(t, "24/04/2022", res.Data.CreatedDate, "la fecha de cracion debe ser igual a 24/04/2022")

	resetStore()
}

func TestFunctionalSaveProduct(t *testing.T) {

	// Create server for test
	r := createServer()

	// Create request for test
	p := Product{0, "Test Product", "color test", 21.45, 34.5, "SDFHK35", false, "24/04/2022"}
	d, _ := json.Marshal(p)
	req, rr := createRequestTest(http.MethodPost, "/products/", string(d))

	// Var that store the response
	var res Response

	// Indicar el servidor que atendera la request
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.Nil(t, err)
	assert.Equal(t, p.Name, res.Data.Name, "los nombres deben ser iguales")
	assert.Equal(t, p.Color, res.Data.Color, "los colores deben ser iguales")
	assert.Equal(t, p.Price, res.Data.Price, "los precios deben ser iguales")
	assert.Equal(t, p.Stock, res.Data.Stock, "la cantidad disponibles deben ser iguales")
	assert.Equal(t, p.Code, res.Data.Code, "los codigos deben ser iguales")
	assert.Equal(t, p.Published, res.Data.Published, "ningun producto debe estar publicado")
	assert.NotEqual(t, p.CreatedDate, res.Data.CreatedDate, "sus fechas de creacion deben ser distintas")

	resetStore()
}

func TestFunctionalUpdateProduct(t *testing.T) {

	// Create server for test
	r := createServer()

	// Create request for test
	id := 1
	p := Product{0, "Test Product Update", "color test update", 21.45, 34.5, "SDFHK35", false, "24/04/2022"}
	d, _ := json.Marshal(p)
	req, rr := createRequestTest(http.MethodPut, fmt.Sprintf("/products/%d", id), string(d))

	// Var that store the response
	var res Response

	// Indicar el servidor que atendera la request
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.Nil(t, err)
	assert.Equal(t, id, res.Data.Id, "el id debe ser igual al del producto respondido")
	assert.Equal(t, p.Name, res.Data.Name, "los nombres deben ser iguales")
	assert.Equal(t, p.Color, res.Data.Color, "los colores deben ser iguales")
	assert.Equal(t, p.Price, res.Data.Price, "los precios deben ser iguales")
	assert.Equal(t, p.Stock, res.Data.Stock, "la cantidad disponibles deben ser iguales")
	assert.Equal(t, p.Code, res.Data.Code, "los codigos deben ser iguales")
	assert.Equal(t, p.Published, res.Data.Published, "ningun producto debe estar publicado")
	assert.Equal(t, p.CreatedDate, res.Data.CreatedDate, "sus fechas de creacion deben ser distintas")

	resetStore()
}

func TestFunctionalUpdateNameProduct(t *testing.T) {

	// Create server for test
	r := createServer()

	// Create request for test
	id := 1
	p := Product{Name: "Test Update Name"}
	d, _ := json.Marshal(p)
	req, rr := createRequestTest(http.MethodPatch, fmt.Sprintf("/products/%d", id), string(d))

	// Var that store the response
	var res Response

	// Indicar el servidor que atendera la request
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.Nil(t, err)
	assert.Equal(t, id, res.Data.Id, "el id debe ser igual al del producto respondido")
	assert.Equal(t, p.Name, res.Data.Name, "los nombres deben ser iguales")
	assert.NotNil(t, res.Data.Color, "este valor no puede ser nulo")
	assert.NotNil(t, res.Data.Price, "este valor no puede ser nulo")
	assert.NotNil(t, res.Data.Stock, "este valor no puede ser nulo")
	assert.NotNil(t, res.Data.Code, "este valor no puede ser nulo")
	assert.NotNil(t, res.Data.Published, "este valor no puede ser nulo")
	assert.NotNil(t, res.Data.CreatedDate, "este valor no puede ser nulo")

	resetStore()
}

func TestFunctionalUpdatePriceProduct(t *testing.T) {

	// Create server for test
	r := createServer()

	// Create request for test
	id := 1
	p := Product{Price: 123456789.987654321}
	d, _ := json.Marshal(p)
	req, rr := createRequestTest(http.MethodPatch, fmt.Sprintf("/products/%d", id), string(d))

	// Var that store the response
	var res Response

	// Indicar el servidor que atendera la request
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.Nil(t, err)
	assert.Equal(t, id, res.Data.Id, "el id debe ser igual al del producto respondido")
	assert.Equal(t, p.Price, res.Data.Price, "los nombres deben ser iguales")
	assert.NotNil(t, res.Data.Name, "este valor no puede ser nulo")
	assert.NotNil(t, res.Data.Color, "este valor no puede ser nulo")
	assert.NotNil(t, res.Data.Stock, "este valor no puede ser nulo")
	assert.NotNil(t, res.Data.Code, "este valor no puede ser nulo")
	assert.NotNil(t, res.Data.Published, "este valor no puede ser nulo")
	assert.NotNil(t, res.Data.CreatedDate, "este valor no puede ser nulo")

	resetStore()
}

func TestFunctionalDeleteProduct(t *testing.T) {

	// Create server for test
	r := createServer()

	// Create request for test
	id := 1
	req, rr := createRequestTest(http.MethodDelete, fmt.Sprintf("/products/%d", id), "")

	// Indicar el servidor que atendera la request
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	resetStore()
}
