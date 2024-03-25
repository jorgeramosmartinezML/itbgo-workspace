-- Ejercicios MySQL

-- 1. Listar todos los actores
SELECT first_name, last_name
FROM actors;

-- 2. Encontrar películas con una calificación superior a 8.5
SELECT title, rating
FROM movies
WHERE rating > 8.5;

-- 3. Buscar actores cuyo nombre de pila comience con 'J'
SELECT first_name
FROM actors
WHERE first_name LIKE 'J%';

-- 4. Listar las 5 películas más recientes
SELECT title, release_date 
FROM movies
ORDER BY release_date DESC 
LIMIT 5;

-- 5. Mostrar los géneros que están activos (active = 1)
SELECT name
FROM genres
WHERE active;

-- 6. Encontrar episodios de la temporada 3
SELECT e.title, s.title 
FROM episodes e
INNER JOIN seasons s
ON e.season_id = s.id
WHERE s.title LIKE 'Tercer%'

-- 7. Listar actores cuyos apellidos contengan 'Smith' y tengan una calificación superior a 7
SELECT first_name, last_name, rating 
FROM actors
WHERE last_name LIKE '%Smith%' AND rating > 7;

-- 8. Mostrar series que comenzaron después del 1 de enero de 2020
SELECT title 
FROM series
WHERE YEAR(release_date) >= 2020;

-- 9. Encontrar películas con una duración entre 90 y 120 minutos
SELECT title, `length` 
FROM movies
WHERE `length` BETWEEN 90 AND 120;

-- 10. Listar los 3 géneros con el ranking más alto
SELECT name, ranking 
FROM genres
ORDER BY ranking DESC
LIMIT 3;


