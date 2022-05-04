# Bases de Datos NoSQL

## MongoDB

Un registro en MOngoDB es un documento, una estructura de datos compuesta por pares de campor y valores, similares a objetos JSON (BSON). Los valores de los campos pueden incluir otros documentos, listas, listas de documentos.

### Ventajas

* Reducen la necesidad de joins costosos
* Su esquema flexible permite polimorfismo

### De SQL a MongoDB

* table -> collection
* row -> document
* column -> field
* index -> index
* primary key (one or some columns) -> primary key (_id field)
* joins -> lookup, embedded documents
* aggregation (group by) -> aggregation pipeline

### CRUD

#### find

```mongodb
db.collection.find(
    {},
    {field_id: 1, status: 1, _id: 0}
)
```

El primer objeto es el filtro en este caso vacio porque queremos recuperar todos los documentos de la colección. El segundo objeto es la *proyección* es decir los campos por documento que queremos traer por eso el valor 1, en el caso de valor 0 se especifica que no queremos mostrar ese field.

#### equal

```mongodb
db.collection.find(
    {status: "A"}
)
```

En el objeto de filtrado, podemos especificar los documentos con determinado field y valor que deseamos obtener.

#### not equal

```mongodb
db.collection.find(
    {status: {$ne: "A"}}
)
```

$ne significa *not equal*, es una función de agregación que nos permite recuperar todos los documentos donde el field status no sea igual a "A".

#### and

```mongodb
db.collection.find(
    {status: "A", age: 50}
)
```

Varias condiciones se estrcuturan bajo el operador logico AND es decir todas sera necesarias de cumplir.

#### or

```mongodb
db.collection.find(
    {$or: [{status: "A"}, {age: 50}]}
)
```

```mongodb
db.collection.find(
    {status: {$in: ["A", "B"]}}
)
```

Bajo la función \$or que recibe un array con las condiciones si quiera una debe ser cumplida. Tambien podemos usar la funcion $in que recibe un array con las opciones posibles para determinado field.

#### lower & greater

```mongodb
$lt, $gt, $lte, $gte
db.collection.find(
    {status: {$gt: 80, $lt: 100}}
)
```

* $lt: less than
* $gt: greater than
* $lte: less than equals
* $gte: greater than equals

#### regex

```mongodb
// contiene la string
db.collection.find(
    {status: /bc/}
)

db.collection.find(
    {status: {$regex: /bc/}}
)

// comienza por
db.collection.find(
    {status: /^bc/}
)

db.collection.find(
    {status: {$regex: /^bc/}}
)
```

Mongo hace uso de expresiones regulares, con la función \$regex podemos buscar valores de ese campo que contengan dicha string, no es necesario hacer uso de la funcion \$regex, ya que se puede directamente con las barras diagonales. El simbolo ^ nos indica que el string debe comenzar con el valor especificado.

#### sort

```mongodb
db.collection.find(
    {status: "A"},
).sort(
    {field_id: -1}
)
```

La función sort nos permite ordenar los documentos con base en uno varios campos, con -1 para forma DESC o 1 para forma ASC pero por defecto la busquedad siempre trae ordenado los documentos de forma ASC.

#### count

```mongodb
db.collection.find().count()

// verifica si el campo es nulo, si no existe dicho campo es nulo
db.collection.find(
    {field_id: {$ne: null}}
).count()

db.collection.find().count(
    {field_id: {$ne: null}}
)

// verifica si el campo existe
db.collection.find().count(
    {field_id: {$exists: true}}
)

// filtra aquellos con valor nulo en el campo especificado
db.collection.find().count(
    {field_id: null}
)
```

Count nos permite contar la cantidad de documentos retornados por el find. Podemos filtrar el tipo de documentos a ser tenidos en cuenta según uno o varios campos para el conteo, especificandolo como filtro del find, o como parametro del count. Con la funcion \$exists podemos validar que el documento tenga el campo en cuestion, si lo tiene y su valor es nulo igualmente lo contará porque el campo si existe sin importar su valor.

#### distinct

```mongodb
db.collection.distinct("field_id")
```

Nos devuelve una colección con un solo documento por cada valor diferente de el campo en mención.

#### limit

```mongodb
db.collection.findOne()

db.collection.find().limit(1)
```

Con limit podemos especficar la cantidad de elementos que queremos recuperar, el findOne nos retorna el primer elemento que conincida con los filtros.

#### skip

```mongodb
db.collection.limit(5).skip(10)
```

Con skip podemos indicar a partir de que registro iniciar el retorno y junto como limit indicar la cantidad de registros a retornar.

### Consultas en MongoDB

#### dot notation

```mongodb
// importa el orden y valor, deben ser exactos
db.collection.find(
    {size: h:14, w:12, uom: "cm"}
)

db.collection.find(
    {"size.h": {$gte: 14}, "size.uom": "cm"}
)
```

#### by index and length

```mongodb
db.collection.find(
    {"tags.0": "red"}
)

db.collection.find(
    {tags: {$size: 3}}
)
```

```mongodb
db.collection.find(
    {"tags": {qty: 5, warehouse: "A"}}
)

db.collection.find(
    {"tags.qty": {$lte: 20}, "tags.warehouse": "A"}
)

// accedemos al documento en posicion 0 y su campo qty
db.collection.find(
    {"tags.0.qty": {$lte: 20}}
)
```

## De-normalizado de Documentos

usar doccumentos embebidos cuendo

* hay relaciones uno-uno
* En relaciones uno-a-muchos si es que "los muchos" siempre hacen falta en contexto "del uno"

## Normalizado

En genral, normalizar documentos cuando

* Anidar documentos resulta en duplicación de los datos pero no en suficiente rendimiento de lectura como para justificar la duplicación
* Hay relaciones muchos-a-muchos
* Jerárquicas (árboles)
* Complejas (redes)

## Agregaciones

### group

```mongodb
db.collection.aggregate([
    {$project: {field_id: 1, subtotal: {$multiply: ["$price", "$quantity"]}}},
    {$group: {field_id: "$field_id", total {$sum: "$subtotal"}}}
])
```

Su crea una lista de documentos donde el id es el field_id y el subtotal es la multiplcaciion del precio por la cantidad y luego agrupa por field_id y suma los subtotales de cada field_id repetido para generar un total.

### lookup

```mongodb
db.collection1.insertMany([
    $lookup: {
        from: "collection2",
        localField: "item",
        foreignField: "sku",
        as: "docs"
    }
])
```

Desde la coleccion 1 el campo item de la coleccion 1 asociar al campo sku de la colección 2 (asociar significa mismo valor) y el valor del campo sku sera el valor del campo docs de cada documento de la coleccion 1.

## Indices

Por defecto se crea un indice sobre el campo _id.

* **Simple:** Generado sobre el campo _id
* **Compuesto:** 