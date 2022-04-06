# Responses & Validation

## Fields Validation

Usar el *binding:"required"* tag en la declaracion de la estructura, para controlar el error usar la siguiente funcion.

```go
if err := c.ShouldbindJSON(&req); err != nil {
    var ve validator.ValidationErrors
    if errors.As(err, &ve) {
        out := make([]ErrorMsh, len(ve))
        for i, fe := range ve {
            out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe, fe.Field())}
        }
    }
}
```

## Generic Response

Implementar un package que maneje las respuestas que retornaremos al cliente, para que las respuestas tengan la misma estructura

### Format

```json
{
"code": 401
"error": "Token invalido"
}
```

### Code Example

```go
type Response struct {
    Code  string
    Data  interface{}  `json:"data,omitempty`
    Error string       `json:"error,omitempty`
}

func NewResponse(code int, data interface{}, err error) Response {
    if code < 400 {
        return Response[strconv.FormatInt(int64(code), 10), data, ""]
    }
    return Response[strconv.FormatInt(int64(code, 10), nil, err.Error())]
}
```
