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

-- Create Sample Table
CREATE TABLE "sample" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    description VARCHAR(255)
);

-- Insert Sample Data
INSERT INTO "sample" (name, description) VALUES 
    ('Sample 1', 'This is a sample description'),
    ('Sample 2', 'This is a sample description'),
    ('Sample 3', 'This is a sample description'),
    ('Sample 4', 'This is a sample description'),
    ('Sample 5', 'This is a sample description');
