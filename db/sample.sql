DROP TABLE IF EXISTS venuePizzas;
DROP TABLE IF EXISTS venues;
DROP TABLE IF EXISTS pizzas;
DROP TABLE IF EXISTS tastes;
DROP TABLE IF EXISTS images;


CREATE TABLE images (
	image_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	filename VARCHAR ( 255 ) NOT NULL,
	content_type VARCHAR ( 50 ) NOT NULL,
	location VARCHAR ( 1000 ) NOT NULL
);

CREATE TABLE tastes (
	taste_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	cheesiness INTEGER NOT NULL,
	flavor INTEGER NOT NULL,
	sauciness INTEGER NOT NULL,
	saltiness INTEGER NOT NULL,
	charness INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL
);

CREATE TABLE pizzas (
	pizza_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name VARCHAR ( 250 ) NOT NULL,
	category VARCHAR ( 250 ) NOT NULL,
	description TEXT NOT NULL,
	image_id INT REFERENCES images(image_id) ON DELETE CASCADE,
	taste_id INT REFERENCES tastes(taste_id) ON DELETE CASCADE,
	updated_at TIMESTAMP,
	created_at TIMESTAMP NOT NULL
);

CREATE TABLE venues (
	venue_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name VARCHAR ( 150 ) NOT NULL,
	lat DOUBLE PRECISION,
	lon DOUBLE PRECISION,
	address VARCHAR ( 255 )
);

CREATE TABLE venuePizzas (
	venue_pizza_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	pizza_id INT REFERENCES pizzas(pizza_id) ON DELETE CASCADE,
	venue_id INT REFERENCES venues(venue_id) ON DELETE CASCADE
);