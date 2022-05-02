# Consultas SQL (DML)

### SELECT

Clausula utilizada con el fin de obtener información de la bs en forma de registros

```sql
SELECT Campos FROM Tabla;
```

### WHERE

Filtra nuestra consulta dependiendo la necesidad.

```sql
SELECT campo1, campo2, campo3 FROM Tabla
WHERE campo1=valor;

SELECT campo1, campo2, campo3 FROM Tabla
WHERE campo1=valor AND campor2=valor;

SELECT campo1, campo2, campo3 FROM Tabla
WHERE campo1=valor OR campo=valor;
```

### ORDER BY

Nos permite ordenar los resultados a partir de 1 o más campos. Es más eficiente solicitarlos ordenados a la base de datos que ordenarlos en el backend. Es más eficiente usar el GROUP BY que el ORDER BY para ordenar los elementos.

```sql
SELECT campo1, campo2, campo3 FROM Tabla
WHERE campo1=valor
ORDER BY campo2;
```

### LIKE

Con like podemos filtrar cadenas de texto segun nuesta necesidad. No se recomienda usar en campos con una longitud de texto grande.

```sql
-- Termina en (t, T)
LIKE '%T'

-- Inicia en (t, T)
LIKE 'T%'

-- Contiene en el medio (t, T)
LIKE '%T%'
```

### LIMIT y OFFSET

Limit es usado para limitar el número de registro devueltos. Offset nos permite especificar apartir de qué fila comenzar la recuperación de datos.

```sql
-- Recuperar solo 10 registros
LIMIT 10;

-- Recuperar 4 registros a partir del 5
LIMIT 4 OFFSET 5
```

### DISTINCT

Nos permite omitir registros que contiene datos duplicados en los campos seleccioados. Solo nos retorna un registro y no todos lo campos.

```sql
SELECT DISTINCT campo2 FROM Tabla;
```

## Funciones de Agregación

* **COUNT:** Devuelve el número total de filas seleccionadas por la consulta
* **MIN:** Devulve el valor mínimo del campo especificado
* **MAX:** Devulve el valor máximo del campo especificado
* **SUM:** Suma los valores del campo que especifquemos
* **AVG:** Deveulve el valor promedio del campo que especifiquemos