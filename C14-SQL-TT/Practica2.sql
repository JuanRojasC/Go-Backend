#Explicar el concepto de normalización y para que se utiliza.
	#Es un proceso de estandarización y validación de los datos para eliminar redundancias e 
    #inconsistencias, por medio de reglas para así proteger la información, protegiendo su integridad 
    #y favorecer la interpretación, así es más sencillo consultar y se gestiona de manera más eficiente.

#Agregar una película a la tabla movies.
INSERT INTO movies_db.movies 
		(created_at, 
		title, 
		rating, 
		awards, 
		release_date, 
		length, 
		genre_id)
	VALUES 
		(sysdate(), 
		"Edward Scissorhands", 
		10.0, 
		8, 
		"1998-05-23", 
		61, 
		5);
SET @id_movies = last_insert_id();
SELECT @id_movies; # Id = 23

SELECT * FROM movies_db.movies WHERE id = @id_movies;

#Agregar un género a la tabla genres.
INSERT INTO movies_db.genres 
	(ranking,
    created_at,
    active,
    name)
    VALUES
    (13,
	sysdate(),
    1,
    "Belico"
    );

SET @id_genre = last_insert_id();
SELECT @id_genre; # Id = 13
    
SELECT * FROM movies_db.genres WHERE id = @id_genre;

#Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies_db.movies 
	SET	genre_id = 13
    WHERE id = 23;

SELECT * FROM movies_db.movies WHERE id = 23;

#Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
SELECT * FROM movies_db.actors;

UPDATE movies_db.actors
	SET favorite_movie_id = 23
    WHERE id = 4; #Actualizamos a Leonardo Di Caprio
    
SELECT * FROM movies_db.actors WHERE id = 4;

#Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE temp_movies 
	SELECT *
		FROM movies_db.movies;
SELECT * FROM temp_movies;

#SET SQL_SAFE_UPDATES = 0; # Deshabilitar la actualizacion segura para la manipulación de tabla temporal

#Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM temp_movies WHERE awards < 5;
SELECT * FROM temp_movies;

#Obtener la lista de todos los géneros que tengan al menos una película.
SELECT genres.id, genres.name, genres.ranking
	FROM movies_db.genres genres
		INNER JOIN movies_db.movies movies ON movies.genre_id = genres.id
	GROUP BY genres.id, genres.name, genres.ranking;
    
SELECT *
	FROM movies_db.genres genres
    WHERE (SELECT COUNT(*) 
			FROM movies_db.movies 
            WHERE genre_id = genres.id) >= 1;

#Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
SELECT actors.id, actors.first_name, actors.last_name, movies.awards
	FROM movies_db.actors actors
		INNER JOIN movies_db.movies movies ON movies.id = actors.favorite_movie_id
	GROUP BY actors.id, actors.first_name, actors.last_name, movies.awards
    HAVING movies.awards > 3;
    
SELECT *
	FROM movies_db.actors actors
    WHERE (SELECT awards
			FROM movies_db.movies 
            WHERE id = actors.favorite_movie_id) > 3;

#Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
EXPLAIN SELECT genres.id, genres.name, genres.ranking
	FROM movies_db.genres genres
		INNER JOIN movies_db.movies movies ON movies.genre_id = genres.id
	GROUP BY genres.id, genres.name, genres.ranking;

EXPLAIN SELECT actors.id, actors.first_name, actors.last_name, movies.awards
	FROM movies_db.actors actors
		INNER JOIN movies_db.movies movies ON movies.id = actors.favorite_movie_id
	GROUP BY actors.id, actors.first_name, actors.last_name, movies.awards
    HAVING movies.awards > 3;

#¿Qué son los índices? ¿Para qué sirven?
	#Son mecanismos para el acceso más directo a los registros
	#logrando así optimizar las consultas reduciendo el tiempo
	#de respuesta en querys complejas y evitando una busqueda lineal.

#Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX idx_movies_title
	ON movies_db.movies(title);
    
#Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;

	#DROP INDEX idx_movies_title
	#	ON movies_db.movies;
