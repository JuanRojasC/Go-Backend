# Database/Sql Package

Implementación nativa de Go, que expone una interfaz que permite gestionar las conexión y datos de una base de datos.

## Catacterísticas

* Forma pparte de las libreias standar de Go
* Compatible con Bases de Datos SQL
* Debe ser usado conjuntamente con un drive de base de datos
* Hacer uso del type sql.DB para gestionar la conexión y la ejecución.
* Puede generar una conexión o administrar un pool de conexiones.

## Install

```bash
go get -u github.com/go-sql-driver/mysql
```

```go
import (
    "databse/sql"
    "log"
    _ "github.com/go-sql-driver-mysql")
```

## Implement

```go

var (
    StorageDB *sql.DB
)

dataSource := "user:pass@tcp(localhost:3306)/dbName"

StorageDB, err := sql.Open("mysql", dataSource)

if err != nil {
    panic(err)
}

if err = db.Ping(); err != nil {
    log.Fatal(err)
}
```

### Notas

* Una sola vez se usa el sql.Open es decir solo se genra una conexión a la BD
* El dataSource es la configuración de la conexión

## Repository

Componentes que ecapsulan la lógican necesaria para el acceso a una fuente de datos.

### Interface

* Deefinimos una interfaz de nuestro Repository.
* Definimos una función constructura que recibe por parametro un puntero a *sql.DB
* Definimos un struc que será nuestro modelo del recurso.

### Implement Methods

#### Store

```go
func ( r. *repository) Store (nombre string, tipo string, cantidad int, precio float64) (Struct, error) {
    stmt, err := r.db.Prepare("INSERT INTO talbe(name, type, count, price) VALUES (?, ?, ?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    // Cerramos la sentencia al terminar para evitar consumos innecesarios de memoria
    defer stmt.Close()
    stru := Struct{name, type, count, price}
    result, err := stmt.Exec(nombre, tipo, cantidad, precio)
    if err != nil {
        return Struct{}, err
    }
    // del sql.Result obtenemos el ultimo Id insertado
    insertId, _ := result.LastInsertId()
    stru.ID = int(insertId)

    return stru, nil
}
```

#### GetOne

```go
func ( r. *repository) GetOne (id int) (Struct, error) {
    var stru Struct
    stmt, err := r.db.Query("SELECT id, name, type, count, price FROM table WHERE id = ?", id)
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next(){
        if err := rows.Scan(&stru.ID, &stru.Name, %stru.Type, &stru.Count, &stru.Price); err != nil {
            return stru, err
        }
    }

    return stru, nil
}
```

#### UPDATE

```go
func ( r. *repository) Store (id int, nombre string, tipo string, cantidad int, precio float64) (Struct, error) {
    stmt, err := r.db.Prepare("UPDATE table SET name = ?, type = ?, count = ?, price = ? WHERE id = ?")
    if err != nil {
        log.Fatal(err)
    }
    // Cerramos la sentencia al terminar para evitar consumos innecesarios de memoria
    defer stmt.Close()
    stru := Struct{name, type, count, price}
    result, err := stmt.Exec(nombre, tipo, cantidad, precio, id)
    if err != nil {
        return Struct{}, err
    }
    // del sql.Result obtenemos el ultimo Id insertado
    insertId, _ := result.LastInsertId()
    stru.ID = int(insertId)

    return stru, nil
}
```
