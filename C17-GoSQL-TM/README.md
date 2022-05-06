# Database/Sql Package

## Implement Methods

### GetAll

```go
func (r *repositpry) GetAll() ([]Struct, error) {
    var structs []Struct
    rows, err := r.db.Query("SELECT * FROM table")
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    for rows.Next() {
        var s Struct
        if err := rows.Scan(&s.ID, &s.Name, &s.Type, &s.Count); err != nil {
            log.Fatal(err)
            return nil, err
        }

        structs = append(structs, s)
    }

    return structs, nil
}
```

### Delete

```go
func (r *repository) Delete(id int) error {
    stmt, err := r.db.Prepare("DELETE FROM table WHERE id = ?", id)
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(id)

    if err != nil {
        return err
    }

    return nil
}
```

## Joins With Data Storage

```go

type Struct struct {
    ID int
    Name string
    Type string
    Count int
    IdJoin int `json:"id_join,omitempty"`
}

func (r *repository) GetFullData(id int) Struct {
    var stru Struct
    innerJoin := "SELECT * FROM table1 INNER JOIN table2 ON table1.id = table2.id WHERE table1.id = ?"
    rows, err := r.db.Query(innerJoin, id)
    if err != nil {
        return Struct{}
    }
    
    for rows.Next() {
        if err = rows.Scan(&s.ID, &s.Name, &s.Type, &s.Count, &s.IdJoin); err != nil {
            return err
        }
    }

    return stru
}
```

## Context

Nos brinda la posibilidad de efectuar cancelaciones de las queries mientras están en plena ejecución. Esto permite administrar los tiempos de duración y el uso de timeouts a las consultas a la DB. Normalmente el contexto es creado desde el handler y pasarlo hasta la capa repository para que lo implemente.

* **Demora en la Consulta:** Cuando una consulta demora más de lo habitual., el contct puede cancelar la ejecución de la misma tas un tiempo determinado.
* **Cancelación de la Request:** El cliente cancela la solicitud (cierra la apliacón)

```go
type Repository interface {
    GetOneWithContext(ctx context.Context, id int) (Struct, error)
}

func (r *respository) GetOneWithContext(Ctx contex.Context, id int) (Struct, error) {
    var stru Struct
    getQuery := "SELECT * FROM table WHERE id = ?"
    rows, err := r.db.QueryContext(ctx, getQuery, id)
    // Permite dormir la consulta para probar el timeout del context de ser necesario
    rows, err := r.db.QueryContext(ctx, "SELECT SLEEP(30) FROM table")

    if err != nil {
        return Struct{}
    }
    
    for rows.Next() {
        if err = rows.Scan(&s.ID, &s.Name, &s.Type, &s.Count, &s.IdJoin); err != nil {
            return err
        }
    }

    return stru
}
```
