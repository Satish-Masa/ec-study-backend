-- name: create-users-table
CREATE TABLE users (
    id       INTEGER PRIMARY KEY,
    email    VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- name: create-items-table
CREATE TABLE items (
    id          INTEGER PRIMARY KEY UNIQUE,
    name        VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    price       INTEGER NOT NULL
);

-- name: create-basket-table
CREATE TABLE basket (
    item_id INTEGER UNIQUE NOT NULL,
    user_id INTEGER NOT NULL,
    name    VARCHAR(255) NOT NULL,
    price   INTEGER NOT NULL
);