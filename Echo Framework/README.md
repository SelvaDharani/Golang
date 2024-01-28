# db_setup
To Run the project in local set , create a local db databases and tables

CREATE DATABASE training;

CREATE TABLE training.users (
    ID INT PRIMARY KEY,
    Name VARCHAR(255),
    Email VARCHAR(255)
);

# environment variables
export DB_PASSWORD = "local db password"

export JWT_KEY = "jwt signing key"
