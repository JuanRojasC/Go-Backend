# PUT, PATCH and DELETE Methods

## PUT

Usamos PUT cuando queremos actualizar un recurso de manera completa, es decir sobreescribir todas sus propiedades.

## PATCH

Usamos PATCH cuando solo queremos actualizar determinadas propiedades del recurso, no queremos obligar al cliente a enviarnos la estructura completa

## DELETE

Usamos DELETE cuando vamos a eliminar un recurso.

### Code Example

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.PUT("/put", func(ctx *gin.Context){})
    r.PATCH("/put", func(ctx *gin.Context){})
    r.DELETE("/put", func(ctx *gin.Context){})
}
```
