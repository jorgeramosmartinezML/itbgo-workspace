-- 1. Con la base de datos “movies”, se propone crear una tabla temporal llamada “TWD” y 
-- guardar en la misma los episodios de todas las temporadas de “The Walking Dead”.
CREATE TEMPORARY TABLE twd
	SELECT e.title AS episode, s.number AS season
	FROM episodes e
	INNER JOIN seasons s ON e.season_id = s.id
	INNER JOIN series s2 ON s.serie_id = s2.id
	WHERE s2.title = 'The Walking Dead' 

--  2. Realizar una consulta a la tabla temporal para ver los episodios de la primera temporada.
SELECT *
FROM twd 
WHERE season = 1;
