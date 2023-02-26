-- Set Environment Variables
\set userdb `echo "$APP_DB_USER"`
\set passdb `echo "$APP_DB_PASS"`
\set dbname `echo "$APP_DB_NAME"`

-- Create Database
CREATE DATABASE :"dbname";

-- Create User and Grant Privileges
CREATE USER :"userdb" WITH ENCRYPTED PASSWORD :'passdb';
GRANT ALL PRIVILEGES ON DATABASE :"dbname" TO :"userdb";
\connect :"dbname" :"userdb"

-- Create Extension and Install into Database
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create User Table
CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(25) NOT NULL
);

-- Create Sample Table
CREATE TABLE "sample" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    description VARCHAR(255),
    owner_id INT NOT NULL REFERENCES "user" (id)
);

-- Insert User Data 
INSERT INTO "user" (username) VALUES 
    ('User 1'),
    ('User 2'),
    ('User 3'),
    ('User 4'),
    ('User 5');

-- Insert Sample Data
INSERT INTO "sample" (name, description, owner_id) VALUES 
    ('Sample 1', 'Sample 1 Description', 1),
    ('Sample 2', 'Sample 2 Description', 2),
    ('Sample 3', 'Sample 3 Description', 2),
    ('Sample 4', 'Sample 4 Description', 3),
    ('Sample 5', 'Sample 5 Description', 4);

