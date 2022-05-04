# Cuestionario

1. **¿Cuántas colecciones tiene la base de datos?**
Tiene una coleccion *restaurantes*

2. **¿Cuántos documentos en cada colección? ¿Cuánto pesa cada colección?**
Tiene 25359 documentos y pesa 4.1KB

3. **¿Cuántos índices en cada colección? ¿Cuánto espacio ocupan los índices de cada colección?**
Tiene 1 indice y el tamaño que ocupa es de 266240

4. **Traer un documento de ejemplo de cada colección. db.collection.find(...).pretty() nos da un formato más legible**

```mongodb
db.restaurantes.findOne()
db.restaurantes.find().limit(1)
```

5. **Para cada colección, listar los campos a nivel raíz (ignorar campos dentro de documentos anidados) y sus tipos de datos.**

```mongodb
db.restaurantes.find({}, {direccion: {edificio: 0, coord:0, calle: 0, codigo_postal: 0}, grados: {$slice: 0}})
```

6. **Devolver restaurante_id, nombre, barrio y tipo_cocina pero excluyendo _id para un documento (el primero).**

```mongodb
db.restaurantes.find({}, {_id: 0, direccion: 0, grados: 0})
```

7. **Devolver restaurante_id, nombre, barrio y tipo_cocina para los primeros 3 restaurantes que contengan 'Bake' en alguna parte de su nombre.**

```mongodb
db.restaurantes.find({nombre: /Bake/}, {_id: 0, direccion: 0, grados: 0}).limit(3)
```

8. **Contar los restaurantes de comida (tipo_cocina) china (Chinese) o tailandesa (Thai) del barrio (barrio) Bronx. Consultar or versus in.**

```mongdb
db.restaurantes.find({$or: [{tipo_cocina: "Chinese"}, {tipo_cocina: "Thai"}], barrio: "Bronx"}).count()
db.restaurantes.find({tipo_cocina: {$in: ["Chinese", "Thai"]}, barrio: "Bronx"}).count()
```

9. **Traer 3 restaurantes que hayan recibido al menos una calificación de grado 'A' con puntaje mayor a 20. Una misma calificación debe cumplir con ambas condiciones simultáneamente; investigar el operador elemMatch.**

```mongdb
db.restaurantes.find({grados: { $elemMatch: {grado: "A", puntaje: {$gt: 20}}}}).limit(3)
```

10. **¿A cuántos documentos les faltan las coordenadas geográficas? En otras palabras, revisar si el tamaño de direccion.coord es 0 y contar.**

```mongdb
db.restaurantes.find({"direccion.coord": {$size: 0}}).count()
db.restaurantes.count({"direccion.coord": {$size: 0}})
```

11. **Devolver nombre, barrio, tipo_cocina y grados para los primeros 3 restaurantes; de cada documento solo la última calificación. Ver el operador slice.**

```mongdb
db.restaurantes.find({}, {direccion: 0, _id: 0, restaurante_id: 0, grados: {$slice: -1}}).limit(3)
```

12. **¿Cuál es top 3 de tipos de cocina (cuisine) que podemos encontrar entre los datos? Googlear "mongodb group by field, count it and sort it". Ver etapa limit del pipeline de agregación.** 

```mongodb
```
