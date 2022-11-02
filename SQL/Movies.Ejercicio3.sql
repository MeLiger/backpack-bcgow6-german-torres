select * from movies;
select first_name, last_name, rating
from actors;
select title
from series;
select * from actors
where rating >= 7.5;
select title, rating, awards 
from movies
where  rating >= 7.5;
select title, rating
from movies
order by rating ASC;
select title
from movies
limit 3;

select * from movies
order by rating desc
limit 5;

select * from actors
limit 10;

select title, rating
from movies
where title like 'Toy Story';

