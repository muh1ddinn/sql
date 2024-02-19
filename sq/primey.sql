CREATE  TABLE  country (

country_id int PRIMARY KEY,
country_name VARCHAR(50) NOT NULL
);

CREATE TABLE city (

city_id SERIAL PRIMARY KEY,
city_ame VARCHAR(50) NOT NULL,
country_id int REFERENCES country(country_id),
population int NOT NULL

);


INSERT INTO country (country_id,country_name)
VALUES 
(1,'GERMANY'),
(2,'TURKEY'),
(3,'UZBEKISTON'),
(4,'USA'),
(5,'CANADA');

INSERT INTO city ( ncity_ame,country_id) VALUES

('Califorina',4),
('ANTALAY',2),
('TERMIZ',3),
('KANSAS',4);
SELECT * FROM counrty ;
SELECT * FROM city;

SELECT AVG(LENGTH(ncity_ame)) AS average_name from city;
/*
problem
SELECT country_id,(SELECT COUNT(city_id)  AS city_count FROM city where country_id=1  ) FROM counrty;

more than 3 
*/

SELECT ncity_ame from city where ncity_ame like 'A';

/*SELECT * FROM city order by ncity_ame desc limit 1;*/



SELECT * FROM city order by city_id asc limit 1;
SELECT * FROM city order by city_id desc limit 1;

/* alter table this add after last row so how can add above rows ?


*/

ALTER TABLE city ADD population int  NOT NULL;

INSERT INTO city ( city_ame,country_id,population)  VALUES

('Ottawa',5,20000),
('TOSHKENT',3,36000),
('NEW YORK',4,3423),
('TORONTO',5,23400),
('ISTANBUL',2,60000),
('SAMARQAND',3,6454),
('OHIO',4,34324),
('Califorina',4,350000),
('ANTALAY',2,34524),
('TERMIZ',3,1204),
('BERLIN',1,45000),
('KANSAS',4,1500);




SELECT SUM(population) AS allcity_population FROM city; 

UPDATE city SET population = 350000 where city_id = 2;

DELETE WHERE city WHERE population >= 100000;

INSERT INTO city (city_ame, country_id,population, creation_time)
VALUES ('New City', 1,4500, CURRENT_TIMESTAMP);


ALTER TABLE city 
ADD creation_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP;


ALTER TABLE city 
ADD update_time TIMESTAMP

UPDATE city SET city_ame = 'NUKUS',update_time= current_timestamp  WHERE city_id = 10;

/* last one have problem*/