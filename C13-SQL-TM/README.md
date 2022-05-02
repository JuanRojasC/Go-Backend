# Data Bases

## Teorema de CAP

Consistency, Availability and Partition Tolerance, establece que en un sistema distribuido de almacenamiento de datos cuando el sistema sufre una partición o un fallo no se puede garantizar consistencia y disiponibilidad al mismo tiempo.

* **Consistencia:** Propiedad que establece que la lectura de datos recibe como respuesta la escritura mas reciente
* **Disponibilidad:** Propiedad que establece que cualquier petición debería recibir una respuesta no errónea, pero sin la garantía de que esta respuesta sea la escritura más reciente.
* **Tolerancia al Particionamiento:** El sistema debe seguir funcionando aunque algunos nodos no se encuentren disponibles ya que la información es consistente en todos los nodos.

## Bases de Datos Relacionales

Las bases de datos relaciones esta compuesta de una o más tablas que se encargan de almacenar el conjunto de datos, cada tabla tiene una o más columnas y filas y las intersección entre ambas es condierada como celda o campo.

### Caracteristicas

* Independencia lógica y física de los datos
* Redudancia mínima
* Acceso concurrente por parte de múltiples usuarios
* Integridad de los datos
* Consultas complejas optimizadas
* Seguridad de acceso y auditoría
* Respaldo y recuperación
* Acceso a través de lenaguajes de programación estandar

## Tipos de Bases De Datos

* Relaciones (SQL)
* Transaccionales (MySQL)
* Documentales (MongoDB)
* Orientada a objetos (PostgreSQL)
* Jerárquicas (Gráfos, Redis)

## DB Relaciones

* Nos permite establecer relaciones entre cada uno de los dato secistentes en sus tablas
* Su funcinamiento radica en introducir todos los datos en registros, organizados en tablas
* usan el Structured Query Language como lenguaje predominante

## Diagramas Entidad Relación (DER)

Herramienta que nos permite modelar los datos para representar las entidades relevantes de un sistema de información así como sus interrelaciones y propiedades.

* **Entidad:** Objeto con atributos
* **Atributos:** Propiedades que describen a una entidad
* **Relación:** Asociacón o relación entre varias entidades
* **Cardinalidad:** Cantidad de valores relacionadas entre tablas

## Keys

### Primary Key

Columna o conjunto de columnas en una tabla cuyos valores identifican de forma única una fila de la tabla.

### Foreign Key

Columna o conjunto de columnas en una tabla cuyos valores corresponden a los valors de la clave primaria de otra tabla.

## Relaciones

* **Relacion 1:1:** Una instancia de la entidad A se relacion unicamente con una de la entidad B
* **Relacion 1:n:** Una instancia de la entidad A se relaciona con varias instancias de la entidad B
* **Relacion n:n:** Cualquier instancia de la entidad A se relaciona con cualquier instancia de la entidad B

## Tipos de datos

Cada campo de la tabla de una base de datos relacional tiene un tiipo de valor

* **Alfanúmericos:** TEXT, CHAR, CHARACTER, STRING, VARCHAR
* **Numéricos:** INTEGER, FLOAT, DOUBLE, NUMBER, NUMERIC
* **Booleanos:** BIT
* **Fechas:** DATETIME, TIME, TIMESTAMP
* **Autoincrementales:** COUNTER
