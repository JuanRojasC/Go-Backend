# POST Method, Group Routes and Headers

## Group Routes

Podemos agrupar diferentes metodos con sus respectivos endpoints a una ruta general.

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    group := r.Group("/products")
    group.GET("/", func(ctx *gin.Context){})
    group.POST("/", func(ctx *gin.Context){})
}
```

## Headers

Obtener acceso a los headers nos permitira verificar y setear metadatos del request.

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.POST("/headers", func(c *gin.Context) {
        token := c.GetHeader("token")
        if token != "123456" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "token invalido",
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "message": "hola!",
        })

    })
}
```

## POST Method

Utilizamos el metodo POST para enviar request que generalmente concluyen con la creación de un nuevo recurso. Los datos contenidos en el body pueden ser obtenidos con el metodo ShouldBindJSON(*pointer) que retorna un error si no puede matchear la información del JSON con los campos de la estructura.

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type request struct {
    ID       int
    Nombre   string
    Tipo     string
    Cantidad int
    Precio   float64
}

func main() {
    r := gin.Default()
    r.POST("/hello", func(ctx *gin.Context){
        var req request
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.JSON(http.StatusOK, req)
    })
}
```
