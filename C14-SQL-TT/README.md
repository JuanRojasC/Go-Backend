# Bases de Datos Relacionales

## Normalización

Proceso de estandarizacion y validacion de daots ue conssite en eliminar las redudancias o incosistencias, completando datos mediante una serie de reglas que actualizan la información, protegiendo su integridad y favoreciendo la interpretación, para que así sea mas simple de consultas y más eficiente para quien la gestiona.

* **1NF:** Elimina datos duplicados en atributos. Crea registros independientes.
*e.g. Atributos con múltiples valores*
* **2NF:** Eliminación de columnas que no dependen de la clave principal, maximiar el uso de FK.
* **3NF:** Elimina subgrupos de datos en múltiples columnas de una tabla y crea tablas nuevas, con relaciones entre ellas.
*e.g. una persona puede tener todos lo numeros telefonicos que quiera, se crea una tabla para los numeros y relaciona con el id de la persona*
* **4NF:** Desaparecen todas las dependencias.

## Sentencias DML

Son aquelllas utilizadas para insertar, leer, actualizar o eliminar registros de la base de datos.

* **Create**

```sql
INSERT INTO actors (first_name,last_name, rating, favorite_movie_id) VALUES ('Charles', 'Creek', 9.0, 11);
```

* Read

```sql
SELECT * FROM movies WHERE id = 1
```

* Update

```sql
UPDATE movies
SET length = 250, genre_id = 5
WHERE id = 1;
```

* Delete

```sql
DELETE FROM movies
WHERE id = 1;
```

## Tablas Temporales

Utilizadas general,ente para hacer pruebas, consultas, análisis, cargas en tablas de staging. La tabla y sus datos se eliminan al finalizar la sesión, se usa para evitar el uso de multiples joins en una consulta. Estas tablan son unicas del usuario no pueden ser compartidas entre usuarios. Las tablas temporales no pueden hacer referencia a Foreign Keys.

```sql
CREATE TEMPORARY TABLE tabla
SELECT * 
FROM table
INNER JOIN table2;

CREATE TEMPORARY TABLE tabla(
    'id' INTEGER NOT NULL PRIMARY KEY
);
```

## Planes de Ejecución

Nos permite visualizar la forma y procedimientos que se llevaron a cabo para materializar la consulta.

```sql
EXPLAIN SELECT * FROM movies;
```

## Indices

* Mejoran los tiempos de respuesta de Queries Complejas
* MEjoran el acceso a los datos al proporcionar una ruta más directa a los registros
* Evitan realizar escaneos (barridas) completas o lineales de los datos en una tabla

### Tipos de indice

* **Indice de Primary Key:** No admite PK duplicadas
* **Indice Ordinario:** Admite duplicados
* **Indice Unico:** Son como los ordinarios pero no admiten duplicados

### Mostrar Indices

```sql
SHOW INDEX FROM tabla
HELP INDEX schema.tabla
```

### Creación de Indices

* **Creacion de tabla**

```sql
CREATE TABLE tabla() UNIQUE PRIMARY INDEX(column1, ..., columnN)
CREATE TABLE tabla() PRIMARY INDEX(column1, ..., columnN)
```

* **Sintaxis**

```sql
CREATE INDEX movies_idx ON MOVIES (id)
ALTER TABLE movies ADD INDEX movies_idx (id)
```
