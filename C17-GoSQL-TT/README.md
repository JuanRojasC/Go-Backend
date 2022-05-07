# Repository Testing

## Package DATA-DOG/go-sqlmock

Librería que implementa el paquete sql/driver con el proposito de simular un motor de vase de datos ene los test sin la necesidad de una real.

* Hace uso del type sql.DB
* Permite simular una base de datos
* Brinda la posibilidad de simular errores o probar casos border

### Install

```bash
go get github.com/DATA-DOG/go-sqlmock
```

### Implement

```go
import "github.com/DATA-DOG/go-sqlmock"

func TestRepositoryStore(t *testing.T) {
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()
    // With regexp
    mock.ExpectPrepare(regexp.QuoteMeta("INSER INTO table(name, price) VALUES (?,?)"))
    mock.ExpectPrepare("INSERT INTO table")
    // With regexp
    mock.ExpectExec(regexp.QuoteMeta("INSER INTO table(name, price) VALUES (?,?)")).WillReturnResult(sqlmock.NewResult(1,1))
    // NewResult(id, rows affected)
    mock.ExpectExec("INSERT INTO table").WillReturnResult(sqlmock.NewResult(1,1))
    repo := NewRepository(db)
    stru := Struct{
        Id: resultId,
        Name: "shirt",
        Price: 100.15
    }
    p, err := repo.Store(stru.Name, stru.Price)
    assert.NoError(t, err)
    assert.NotZero(t, p)
    assert.Equal(t, stru.ID, p.Id)
    assert.NoError(t, mock.ExpectationWereMet())
}

func TestRepositoryGetWithTimeout(t *testing.T) {
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    resultId := 1
    columns := []string{"id", "name", "price"}
    rows := sqlmock.NewRows(columns, resultId, "shirt", 1034.46)
    rows.AddRow(resultId, "shirt", 100.15)
    mock.ExpectQuery("SELECT id, name, price").WillDelayFor(30 * time.Second).WillReturnRows(rows)

    repo := NewRepository(db)
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    _, err := repo.GetOneWithContext(ctx, resultId)
    assert.Error(t, err)
}
```

## Package DATA-DOG/go-txdb

Es un paquete que proprociona un driver de sql que al generar una conexión inicia una transaciión y genera un rollback al realizar un cierre.

* Permite realizar test más robustos ya que utilizara las restricciones reales del motor de base de datos que se utilice (unique, tipos de datos, claves foraneas)
* Depende de una conexión real, solo aisla el codigo dentro de una transacción

### Install go-txdb

```bash
go get github.com/DATA-DOG/go-txdb
```

### Implement go-txdb

```go
import (
    "database/sql"
    "github.com/go-sql-driver/mysql"
    "github.com/DATA-DOG/go-txdb"
    "github.com/google/uuid"
)

func TestRepository(t *testing.T) {
    txdb.Register("txdb", "mysql", "root@tcp(localhost:3306/storage"))
    db, err := sql.Open("txdb", uuid.New().String())
    repo := NewRepository(db)
    stru := Struct{
        Id: resultId,
        Name: "shirt",
        Price: 100.15
    }
    p, err := repo.Store(stru.Name, stru.Price)
    assert.NoError(t, err)
    assert.NotZero(t, p)
}
```

go-txdb no funciona al ser ejecutado en entornos de CI debido a que requiere una conexión real a la base de datos.
