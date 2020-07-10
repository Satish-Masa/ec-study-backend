-- name: create-users-table
CREATE TABLE users (
    id       INTEGER PRIMARY KEY UNIQUE NOT NULL,
    email    VARCHAR(255),
    password VARCHAR(255)
);