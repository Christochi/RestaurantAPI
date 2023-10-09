-- DROP TABLES
DROP TABLE IF EXISTS Chef;
DROP TABLE IF EXISTS Menu;

-- CREATE Chef Table
CREATE TABLE Chef (
   id SERIAL PRIMARY KEY,
   full_name  VARCHAR(50),
   about VARCHAR(255),
   image_name VARCHAR(50),
   gender VARCHAR(1), -- "M"/"F" (male/female)
   age INT,
   UNIQUE (about, image_name) -- unique constraint to avoide duplicate data
);

-- CREATE Menu Table
CREATE TABLE Menu (
   id SERIAL PRIMARY KEY,
   meal_type  VARCHAR(50),
   meal_name VARCHAR(50),
   price VARCHAR(50),
   about VARCHAR(255),
   image_name VARCHAR(50),
   UNIQUE (meal_name, image_name) -- unique constraint to avoide duplicate data
);