CREATE DATABASE hezzl;

\connect hezzl;

CREATE TABLE
    users(
        id SERIAL NOT NULL PRIMARY KEY,
        name VARCHAR(100),
        age INT
    );