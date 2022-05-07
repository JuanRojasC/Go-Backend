package products

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )"))
	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))
	productId := 1
	repo := NewRepository(db)
	user := Product{
		ID:    productId,
		Name:  "remera",
		Type:  "indumentaria",
		Count: 3,
		Price: 1500,
	}
	p, err := repo.Store(user.Name, user.Type, user.Count, user.Price)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, user.ID, p.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryGetWithTimeout(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	productId := 1
	columns := []string{"id", "name", "type", "count", "price"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "remera", "indumentaria", 3, 1500)
	mock.ExpectQuery("select id, name, type, count, price").WillDelayFor(30 * time.Second).WillReturnRows(rows)
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = repo.GetOneWithContext(ctx, productId)

	assert.Error(t, err)
}

func TestRepositoryStoreTxDB(t *testing.T) {

	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	repo := NewRepository(db)
	user := Product{
		Name:  "remera",
		Type:  "indumentaria",
		Count: 3,
		Price: 1500,
	}
	p, err := repo.Store(user.Name, user.Type, user.Count, user.Price)
	assert.NoError(t, err)
	assert.NotZero(t, p)
}
