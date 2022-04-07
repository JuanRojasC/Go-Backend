# Project Structure in Golang

## MVC Pattern

### Controller

* Recibe la peticion del cliente, validaa valores requeridos, ejecutar los diferentes servicios y retornar una respuesta.

### Service

* Realiza las funciones principales d ela aplicacion, procesamiento de datos, implementacion de funciones de negocio y administracion de recursos externos, por ejemplo, base de datos o api externas.

### Repository

* Se encarga de abstraer el acceso a los datos, siendo el encargado de interactuar con la base de datos o sistema de persistencia de datos.

## Project Structure

### Layer

Generar un paquete por cada capa, si tenemos varias entidades como Producto y Empleado, todas sus capas estaran en cada paquete. La desventaje es que si quisieeramos quitar la entidad porducto para implementarlo en otro microservicio, deberiamos modificar todos los paquetes.

```txt
controller :
    empleadoController.go
    productoController.go

repository:
    empleadoRepository.go
    productoRepository.go
service:
    empleadoServuce.go
    productService.go
```

### Domain

Generamos un pacquete por cada entidad, cada paquete tendra todas las capas de la entidad.

```txt
empleado
    controlador.go
    repositorio.go
    servicio.go

productos
    controlador.go
    repositorio.go
    servicio.go
```

## Packages Structure

Organización de las carpetas en un proyecto golang.

```txt
cmd (main & controllers)
    - server
        - handler
            - products.go
            - employees.go
        - main.go
internal (services & repositories)
    - products
        - repository.go
        - service.go
    - employees
        - repository.go
        - service.go
```
