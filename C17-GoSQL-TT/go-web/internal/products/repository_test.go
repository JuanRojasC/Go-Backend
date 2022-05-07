package products

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func dataSource() string {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_TABLE"))
}

func TestGetAllRepository(t *testing.T) {
	txdb.Register("txdb", "mysql", dataSource())
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(db)
	ps, err := repo.GetAll()
	currentRegisters := 6

	assert.Nil(t, err, "el error debe ser nulo")
	assert.Equal(t, currentRegisters, len(ps), "la cantidad de productos debe ser igual")
}

func TestGetByIDRepository(t *testing.T) {
	txdb.Register("txdb", "mysql", dataSource())
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(db)
	idExpected := 1
	p, err := repo.GetByID(idExpected)

	assert.NoError(t, err)
	assert.Equal(t, idExpected, p.Id, "debe ser el mismo id")
}

func TestGetByNameRepository(t *testing.T) {
	txdb.Register("txdb", "mysql", dataSource())
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(db)
	nameExpected := "test nombre"
	minProductsExpected := 1
	ps, err := repo.GetByName(nameExpected)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(ps), minProductsExpected, "por lo menos 1 producto")
}

func TestSaveRepository(t *testing.T) {
	txdb.Register("txdb", "mysql", dataSource())
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(db)
	np := buildProduct()
	rp, err := repo.Save(np.Name, np.Color, np.Price, np.Stock, np.Code, np.Published)

	assert.NoError(t, err)
	assert.Equal(t, np.Name, rp.Name)
	assert.Equal(t, np.Color, rp.Color)
	assert.Equal(t, np.Price, rp.Price)
	assert.Equal(t, np.Stock, rp.Stock)
	assert.Equal(t, np.Code, rp.Code)
	assert.Equal(t, np.Published, rp.Published)
}

func TestUpdateRepository(t *testing.T) {
	txdb.Register("txdb", "mysql", dataSource())
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(db)
	np := buildProduct()
	_, err = repo.Save(np.Name, np.Color, np.Price, np.Stock, np.Code, np.Published)
	up := Product{0, "unit test update", "gray update", 349.6, 439, "JDJGF84NJUPDATE", false, time.Now().Format("2006-01-02 15:04:05")}

	assert.NoError(t, err)
	assert.Equal(t, np.Name, up.Name)
	assert.Equal(t, np.Color, up.Color)
	assert.Equal(t, np.Price, up.Price)
	assert.Equal(t, np.Stock, up.Stock)
	assert.Equal(t, np.Code, up.Code)
	assert.Equal(t, np.Published, up.Published)
}

func TestDeleteRepository(t *testing.T) {
	txdb.Register("txdb", "mysql", dataSource())
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(db)
	np := buildProduct()
	sp, err := repo.Save(np.Name, np.Color, np.Price, np.Stock, np.Code, np.Published)
	errDelete := repo.Delete(sp.Id)
	rp, err := repo.GetByID(sp.Id)

	assert.NoError(t, errDelete)
	assert.Equal(t, Product{}, rp)
}

func TestUnitaryGetAllRepositoryOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := buildRowsProducts()
	mock.ExpectQuery(getAll).WillReturnRows(rows)

	repo := NewRepository(db)
	ps, err := repo.GetAll()
	expectedItems := 3

	assert.NoError(t, err)
	assert.Equal(t, expectedItems, len(ps))
}

func TestUnitaryGetByIDOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	p := Product{1, "test product", "color test", 2343.5, 245.4, "JSDOJ34GF", false, "2022-05-06 20:05:00"}
	rows := buildRowProduct(p.Id, p.Name, p.Color, p.Price, p.Stock, p.Code, p.Published, p.CreatedDate)
	mock.ExpectQuery(getByID).WillReturnRows(rows)

	repo := NewRepository(db)
	rp, err := repo.GetByID(p.Id)

	assert.NoError(t, err)
	assert.Equal(t, p.Id, rp.Id)
	assert.Equal(t, p.Name, rp.Name)
	assert.Equal(t, p.Color, rp.Color)
	assert.Equal(t, p.Price, rp.Price)
	assert.Equal(t, p.Stock, rp.Stock)
	assert.Equal(t, p.Code, rp.Code)
	assert.Equal(t, p.Published, rp.Published)
	assert.Equal(t, p.CreatedDate, rp.CreatedDate)
}

func TestUnitaryGetByNameOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	itemsExp := 1
	p := Product{1, "test product", "color test", 2343.5, 245.4, "JSDOJ34GF", false, "2022-05-06 20:05:00"}
	rows := buildRowProduct(p.Id, p.Name, p.Color, p.Price, p.Stock, p.Code, p.Published, p.CreatedDate)
	mock.ExpectQuery(getByName).WillReturnRows(rows)

	repo := NewRepository(db)
	ps, err := repo.GetByName(p.Name)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(ps), itemsExp)
	assert.Equal(t, p.Id, ps[0].Id)
	assert.Equal(t, p.Name, ps[0].Name)
	assert.Equal(t, p.Color, ps[0].Color)
	assert.Equal(t, p.Price, ps[0].Price)
	assert.Equal(t, p.Stock, ps[0].Stock)
	assert.Equal(t, p.Code, ps[0].Code)
	assert.Equal(t, p.Published, ps[0].Published)
	assert.Equal(t, p.CreatedDate, ps[0].CreatedDate)
}

func TestUnitarySaveRepositoryOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(saveProduct))
	mock.ExpectExec(regexp.QuoteMeta(saveProduct)).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)
	np := buildProduct()
	p, err := repo.Save(np.Name, np.Color, np.Price, np.Stock, np.Code, np.Published)

	assert.NoError(t, err)
	assert.Equal(t, np.Name, p.Name)
	assert.Equal(t, np.Color, p.Color)
	assert.Equal(t, np.Price, p.Price)
	assert.Equal(t, np.Stock, p.Stock)
	assert.Equal(t, np.Code, p.Code)
	assert.Equal(t, np.Published, p.Published)
}

func TestUnitarySaveRepositoryInternalError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(saveProduct))
	mock.ExpectExec(regexp.QuoteMeta(saveProduct)).WillReturnError(ErrInternalServer)

	repo := NewRepository(db)
	np := buildProduct()
	_, err = repo.Save(np.Name, np.Color, np.Price, np.Stock, np.Code, np.Published)

	assert.Error(t, err)
}

func TestUnitaryUpdateRepositoryOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	mock.ExpectPrepare(regexp.QuoteMeta(updateProduct))
	mock.ExpectExec(regexp.QuoteMeta(updateProduct)).WillReturnResult(sqlmock.NewResult(1, 1))
	p := buildProduct()
	rows := buildRowProduct(p.Id, p.Name, p.Color, p.Price, p.Stock, p.Code, p.Published, p.CreatedDate)
	mock.ExpectQuery(getByID).WillReturnRows(rows)

	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	up, err := repo.Update(ctx, p.Id, p.Name, p.Color, p.Price, p.Stock, p.Code, p.Published)

	assert.NoError(t, err)
	assert.Equal(t, p.Id, up.Id)
	assert.Equal(t, p.Name, up.Name)
	assert.Equal(t, p.Color, up.Color)
	assert.Equal(t, p.Price, up.Price)
	assert.Equal(t, p.Stock, up.Stock)
	assert.Equal(t, p.Code, up.Code)
	assert.Equal(t, p.Published, up.Published)
}

func TestUnitaryDeleteRepositoryOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	p := buildProduct()
	rows := buildRowProduct(p.Id, p.Name, p.Color, p.Price, p.Stock, p.Code, p.Published, p.CreatedDate)
	mock.ExpectQuery(getByID).WillReturnRows(rows)
	mock.ExpectPrepare(regexp.QuoteMeta(deleteByID))
	mock.ExpectExec(regexp.QuoteMeta(deleteByID)).WillReturnResult(sqlmock.NewResult(1, 1))

	id := 1
	repo := NewRepository(db)
	err = repo.Delete(id)

	assert.NoError(t, err)
}

func TestUnitaryDeleteRepositoryNotFoundError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	columns := []string{"id", "name", "color", "price", "stock", "code", "published", "created_date"}
	rows := sqlmock.NewRows(columns)
	mock.ExpectQuery(getByID).WillReturnRows(rows)

	id := 1
	repo := NewRepository(db)
	err = repo.Delete(id)

	assert.Error(t, err)
}

func buildProduct() Product {
	return Product{1, "unit test", "gray", 3445.6, 123, "JDJGF84NJ", false, time.Now().Format("2006-01-02 15:04:05")}
}

func buildRowProduct(id int, name string, color string, price float64, stock float64, code string, published bool, createdDate string) *sqlmock.Rows {
	columns := []string{"id", "name", "color", "price", "stock", "code", "published", "created_date"}
	rows := *sqlmock.NewRows(columns)
	rows.AddRow(id, name, color, price, stock, code, published, createdDate)
	return &rows
}

func buildRowsProducts() *sqlmock.Rows {
	columns := []string{"id", "name", "color", "price", "stock", "code", "published", "created_date"}
	rows := *sqlmock.NewRows(columns)
	var products []Product
	file, err := os.ReadFile("../../products.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(file, &products); err != nil {
		log.Fatal(err)
	}
	for _, f := range products {
		rows.AddRow(f.Id, f.Name, f.Color, f.Price, f.Stock, f.Code, f.Published, f.CreatedDate)
	}
	return &rows
}
