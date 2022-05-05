# Bases de Datos NoSQL

## Replicación

Es un grupo de procesos que mantienen el mismo dataset. Caada proceso corre en un máquina distinta (nodo) dentro de un grupo de máquinas (cluster). Las réplicas proveen redundancia y tolerancia a fallos, son la base de todos los deploys en producción.

* **Nodo Primario:** Recibe todas las operaciones de escritura y lectura; solo puede haber uno, replica los datos en los demás nodos.
* **Nodos Secundarios:** Replican las operaciones del primario en sus datasets, reflejando todo cambio.
* **Tolerancia a Fallos:** Cuando el primario falla (10 seg por defecto) (heartbeat), un secudnario llama a elección para nominarse como nuevo primario. El replica set no puede procesar operaciones de escritura durante la elección.
* **Write Concern:** Las operaciones de escritura requieren un reconocimiento de persistencia. Este reconocimiento es configurable, cuando su valor es mayoria, la escritura se confirma cuando las operaciones se hayan propagado en la mayoría de nodos.
* **Distribución Geografica:** Las colecciones fragmentadas son particionadas y distribuidas en los *shards* -nodos- del cluster. Las colecciones no particionadas son guardadas en un shard primario.

## Ventajas de la Fragmentación

Sharding en MongoDb es un método para distribuir los datos a lo largo de múltiples máquinas. Es útil cuando el dataset es muy grande o la cantidad de operaciones es muy alta. Dividir la carga entre múltiples servidores y añadir servidores según necesidad (escalabilidad horizontal).

## Conexión

Los clientes jamás deberían conectarse a un fragmento directamente. Deben hacerlo a traves de de un nodo enrutados (proceso mongos).

## Shard

### Shard Keys

La shard key es un índice simple o compuesto que determina la distribución de los documentos de una colección entre fragmentos del cluster. Mongo divide  los valores de la llave en rangos yuxtapuestos. Cada rango es asociado a un pedazo (chunk). Mongo intenta distribuirlos de manera equitativa entre los shards.

* Alta cardinalidad
* Baja frecuencia de valores
* No ser monótona

### Balanceador

Si la diferencia en el número de pedazos entre el shard más ocupado y el shard mas vacio supera un umbral el balanceador hace una migración para divir equitativamente la información.

### Operaciones dirigidas

Las consultas más rápidas son aquellas que el enrutador puede dirigir a un único shard utilizando la shard key.

## Caracteristicas de Bases NoSQL

* Escalabilidad Horizontal
* Tolerancia a la partición
* Consultas más rapidas (datos que se acceden juntos se guardan juntos)
* Facilidad de desarollo
* Falta de transacciones entre múltiples registros
* Posible duplicación de la información

## Tipos de Bases de Datos NoSQL

### Documentos

Almacenan coleciones de documentos similares a objetos JSON. Los documentos son similares pero no necesariamento iguales. *e.g. MongoDB*. Se encuentran pensadas pra datos con estructura compleja y variable.

### Llave-Valor

Almacenan colecciones como si fuera un diccionario, la llave funciona como identificador único, los valores pueden ser de cualquier tipo, estas bases tratan a los datos que contienen como BLOBs - 1s y 0-s sin estructura. Debido a que los BLOBs son opacos, esta bd no realizan búsquedades en base al contenido de los valores, ni establecen otro índice más que para las llaves. Son la opción más eficiente cuando las consultas para recuperar datos no complejas, usualmente son usadas en colas de mensajes, manejo de sesiones y almacenamiento cache.

### Columnares

Similares a las bases relaciones, su diferencia se encuentra en que almacenan la información paginas por columna, de tal forma que al solicitar una columna de una fila solo lee dicha columna y filtra por el valor indicado. Son óptimas para lectura rápidas de grandes vólumenes de datos y agregaciones de columnas, son usados especialmente en Big Data, series temporales (IoT, valores de mercados, telemetría) y Logs de aplicaciones.

### Grafos

Los documentos suelen ser entidades (personas, lugares, cosas); las relaciones son semánticamente relevantes (pertenece a, compró un) y pueden tener atributos (distancia, tiempo, costo). Se usan a hora de "atravesar" una red de documentos en busca de patrones.
