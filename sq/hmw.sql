




CREATE TABLE userr (
    id serial,
    name VARCHAR(34),
    country VARCHAR(43),
    salary  int
    age int
);
INSERT INTO userr (name,country,salary,age) 
VALUES ('Adam', 'USA', 4000,34),
('Adrian', 'USA', 4500,27),
('Jack', 'USA', 3500,18),
('Jonathan', 'USA', 4100,29),
('Kevin', 'USA', 3200,20),
('Max', 'USA', 5000,19),
('Nicholas', 'USA', 5500,21),
('Robert', 'USA', 4700,24),
('Paul', 'USA', 4300,23),
('Peter', 'USA', 7000,22),
('Stephen', 'USA', 6400,25),
('Tim', 'USA', 6300,28),
('Sebastian', 'USA', 3900,17),
('Warren', 'USA', 5400,26),
('Thomas', 'USA', 2200,18);
SELECT COUNT(id) FROM userr;
/*
 count 
-------
    15
(1 row)
*/
SELECT MIN(age) FROM userr;
/*
 min 
-----
  17
(1 row)

*/
SELECT MAX(salary) FROM userr;
/*
 max  
------
 7000
(1 row)

*/
SELECT AVG(age) FROM userr;
/*
         avg         
---------------------
 23.4000000000000000
(1 row)
*/

SELECT SUM(salary) FROM userr;
/*
 total_salary 
--------------
        70000
(1 row)

this one  has issues 

*/