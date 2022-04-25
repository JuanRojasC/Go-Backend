# Functional Testing & Test Driven Development (TDD)

## Funtional Testing

Es un tiipo de test black box que tiene como objetvio probar un requerimiento funcional especifico del sfotware

* Probar un requerimiento funcional concreto
* Es necesario la integracion múltiples componentes
* Son más costosos de realizar
* Son lentos a comparación de los unit tests

### httptest Package

Nos permite consstruir Functional Test, End to End Tests e Integration Tests. Una de las ventajas es que puede ser usado junto con librerías nativas de GO como librerias de terceros (Gin).

```go
req := httptest.NewRequest("GET", "http://localhost:8080/endpoint", nil)
var response *httptest.REsponseRecorder = httptest.NewRecorder()
func (engine *Engine) serverHTTP(w http.ResponseWriter, req *http.Request)
```

## Test Driven Development

Técnica para el desarrollo de software que tiene como objetivo entender los casos de uso, escribir los test y finalmente de forma iterativa implementar la solución.

* Primero escribir el test luego el código
* Con TDD se puede alcanzar una cobertura de 100% mas facilmente
* Promueve el refactor continuo
