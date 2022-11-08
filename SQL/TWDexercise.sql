SELECT e.* FROM episodes e 
INNER JOIN series sr
INNER JOIN seasons s ON e.season_id = s.id WHERE s.serie_id = (
SELECT id FROM series where title = "The Walking Dead"
 );
 
CREATE TEMPORARY TABLE `TWD`
SELECT e.* FROM episodes e
INNER JOIN seasons s ON e.season_id = s.id WHERE s.serie_id = (
SELECT id FROM series WHERE title = "The Walking Dead"
);

SHOW CREATE TABLE TWD;