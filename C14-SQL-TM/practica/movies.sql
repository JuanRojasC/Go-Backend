USE movies_db;

-- Mostrar el título y el nombre del género de todas las series.
SELECT s.title, g.name
FROM genres g
INNER JOIN series s
WHERE g.id = s.id;

-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT 
e.title AS chapter, 
ac.first_name AS actor_name, 
ac.last_name AS last_name_actor
FROM episodes e
INNER JOIN actor_episode ae
ON e.id = ae.episode_id
INNER JOIN actors ac
ON ae.actor_id = ac.id;

-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT s.title AS serie, COUNT(*) AS cantidad_temporadas
FROM series s
INNER JOIN seasons ss
ON s.id = ss.serie_id
GROUP BY s.title;

-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT g.name, COUNT(m.id) AS cantidad_peliculas
FROM genres g
LEFT JOIN movies m
ON g.id = m.genre_id
GROUP BY g.name
HAVING cantidad_peliculas >= 3;

-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT ac.first_name, ac.last_name 
FROM actors ac
INNER JOIN actor_movie am
ON ac.id = am.actor_id
INNER JOIN movies m
ON am.movie_id = m.id AND m.title LIKE '%Guerra de las galaxias%'