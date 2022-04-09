# API Documentation & Middlewares

## API Docs

Es una pieza básica para el mantenimiento y escalabilidad de un proyecto.

* Descripcion de endpoint detallada, desribir que hace cada endpoint desde un lenguaje comprensible, no técnico.
* Detallar parametros y respuestas esperadas, con el fin de facilitar el conocimiento de envio y recepcion de información

### Swagger

Swagger es una herramienta útil para describir, producti, cosumir y visualizar APIs RESTful.

#### Package swaggo

#### install

```bash
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/swag/files
go get -u github.com/swaggo/gin-swagger
```

#### use

```bash
export PATH=$PATH:$HOME/go/bin

cd
nano .bashrc
export PATH=$PATH:$HOME/go/bin

swag init -g cmd/server/main.go
```

```go
import (
    "github.com/JuanDRojasC/C9-GoWeb-TT/go-web/docs"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

docs.SwaggerInfo.Host = os.Getenv("HOST")
r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiler.Handler))
```

#### Annotations

##### Main

```go
// @title Bootcamp API
// @version 1.0
// @description This API Handle Products
// @termOfService https://developers.com/es_co/tyc
// @contac.name API Support
// @contac.url https://developers.com/support
// @license.name Apache 2.0
// @license.url https://apache.org/licenses/LICENSE-2.0.html
func main() {}
```

##### Endpoints

```go
// SaveProduct godoc
// @Summary Save a new product
// @Tags Products
// @Description Save a new product
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body fullRequest true "Product to Save"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response "Invalid Token"
// @Failure 422 {object} web.Response "Body malformed"
// @Router /products [post]
func (ph *ProductHandler) SaveProduct() gin.HandlerFunc {}
```

*[Annotations Description](https://github.com/swaggo/swag#api-operation)*

## Middlewares

### Methods

* **gin.Default().Use():** Permite setear un middleware que afectara a todas las rutas
* **gin.Context.Next():** Permite continuar con los middlewares siguientes
* **gin.Context.AbortWithStatusJSON():** Permite abortar la ejecución de los siguientes middlewares
* **gin.Deafult().METHOD("/route", Middleware, Action):** Seteamos un middleware a un ruta en especifico

### Code Example

```go
package main

import (
    "fmt"
    "github.con/gin-gonic/gin"
)

func GetDummyEndpoint(c *gin.COntext) {
    fmt.Println("I am dummy")
}

func Welcome(c *gin.Context) {
    fmt.Println("Hello to BootCamp")
}

func Middleware(c *gin.Context) {
    fmt.Printl("I am Middleware")
}

func MiddleWareDummy(c *gin.Context) {
    fmt.Println("Middleware for route")
    c.Next()
}

func main() {
    api = gin.Default()

    api.Use(Middleware)

    api.GET("/", Welcome)
    api.GET("/dummy", MiddelWareDummy, GetDummyEndpoint)
}

```
