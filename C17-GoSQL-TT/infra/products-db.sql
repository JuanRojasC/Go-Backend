DROP DATABASE IF EXISTS products;
CREATE DATABASE IF NOT EXISTS products;
USE products;

DROP TABLE IF EXISTS products;
CREATE TABLE products(
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  `color` VARCHAR(25) NOT NULL,
  `price` DECIMAL(18, 2) NOT NULL,
  `stock` DECIMAL(20,2) NOT NULL,
  `code` VARCHAR(50) NOT NULL,
  `published` BOOLEAN NOT NULL,
  `created_date` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO products VALUES (1,'MacBook Pro','gray',989,134,'8977JNFJKD',true,'2022-04-05 22:23');
INSERT INTO products VALUES (2,'iPhone','gray',1145,67,'949JJ54LF',true,'2022-04-05 22:23');
INSERT INTO products VALUES (3,'Magic Mouse','white',79.99,96,'JFOVN405',true,'2022-04-05 22:23');
INSERT INTO products VALUES (4,'test nombre','green update',5456.76,345,'OFJ30585KF',true,'2022-04-05 22:23');
INSERT INTO products VALUES (5,'test nombre','green update',5456.76,345,'OFJ30585KF',true,'2022-04-24 20:29');