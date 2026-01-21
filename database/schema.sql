-- database/schema.sql
CREATE TABLE subjects (id INTEGER PRIMARY KEY, name TEXT);
CREATE TABLE actions (id INTEGER PRIMARY KEY, description TEXT);
CREATE TABLE locations (id INTEGER PRIMARY KEY, name TEXT);

-- Ejemplo de datos:
INSERT INTO subjects (name) VALUES ('Elon Musk'), ('Los Gatos'), ('La NASA');
INSERT INTO actions (name) VALUES ('están clonando'), ('ocultan un portal en'), ('venden datos a');
INSERT INTO locations (name) VALUES ('la Luna'), ('tu refrigerador'), ('el Área 51');