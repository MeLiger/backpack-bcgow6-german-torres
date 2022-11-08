SELECT * from movies_db.movies;


SELECT m.id, m.title, a.first_name, a.last_name, a.favorite_movie_id 
FROM movies m
INNER JOIN actors a
ON m.id = a.favorite_movie_id;

SELECT m.id, m.title, a.first_name, a.last_name, a.favorite_movie_id,
CONCAT(a.first_name," ",a.last_name) AS "nombre y apellido", a.favorite_movie_id
FROM movies m
INNER JOIN actors a
ON m.id = a.favorite_movie_id;

#LEFT JOIN
SELECT m.title AS "Pelicula Favorita", CONCAT(a.first_name," ",a.last_name) AS "Actor"
FROM movies m
LEFT JOIN actors a ON m.id = a.favorite_movie_id;

#RIGHT JOIN
SELECT m.title AS "Pelicula Favorita", CONCAT(a.first_name," ",a.last_name) AS "Actor"
FROM movies m
RIGHT JOIN actors a ON m.id = a.favorite_movie_id;

SELECT m.title AS "Pelicula Favorita", CONCAT(a.first_name," ",a.last_name) AS "Actor"
FROM movies m
LEFT JOIN actors a ON m.id = a.favorite_movie_id
WHERE a.favorite_movie_id IS NULL;

SELECT m.title AS "Pelicula Favorita", CONCAT(a.first_name," ",a.last_name) AS "Actor"
FROM movies m
RIGHT JOIN actors a ON m.id = a.favorite_movie_id
WHERE a.favorite_movie_id IS NULL;

#Otras consultas
SELECT * FROM movies;

SELECT m.title, SUM(m.awards) AS total_awards
FROM movies m
GROUP BY id HAVING total_awards = 2; 



