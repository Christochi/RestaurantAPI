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
   age INT
);