CREATE SCHEMA IF NOT EXISTS backend;
CREATE TABLE IF NOT EXISTS backend.Warehouse (
	id serial PRIMARY KEY,
	name varchar(255) UNIQUE NOT NULL,
	accessibility bool NOT NULL 
);
CREATE TABLE IF NOT EXISTS backend.Product (
	id serial PRIMARY KEY,	
	name varchar(255) NOT NULL UNIQUE,
	mass integer NOT NULL CHECK (mass >= 0)
);
CREATE TABLE IF NOT EXISTS backend.Placement (
	warehouse_id int REFERENCES backend.Warehouse (id) ON DELETE RESTRICT,
	product_id int REFERENCES backend.Product (id) ON DELETE RESTRICT,
	PRIMARY KEY (warehouse_id, product_id),
	total int NOT NULL CHECK (total >= 0),
	reserved int NOT NULL CHECK (reserved >= 0),
	CHECK (reserved <= total)
);

INSERT INTO backend.product ("name",mass) VALUES
	('adidas shoe',1000),
	('adidas pants',200),
	('red hat',100),
	('nike shoe',999),
	('uniqlo shirt',100),
	('uniqlo jacket',500),
	('uniqlo pants',300),
	('nike shirt',100);

INSERT INTO backend.warehouse ("name",accessibility) VALUES
	('mow',true),
	('cek',true),
	('svx',true);

INSERT INTO backend.placement (warehouse_id, product_id, total, reserved) VALUES
	(1,1,100,0),
	(1,2,100,0),
	(1,3,100,0),
	(1,4,100,0),
	(1,5,100,0),
	(2,3,100,0),
	(2,4,100,0),
	(2,6,100,0),
	(3,7,100,0),
	(3,1,100,0),
	(3,2,100,0),
	(3,5,100,0);
