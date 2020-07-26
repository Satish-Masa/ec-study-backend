-- +migrate Up
CREATE TABLE users (
    id       INTEGER NOT NULL AUTO_INCREMENT,
    email    VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

-- +migrate Up
CREATE TABLE items (
    id          INTEGER NOT NULL AUTO_INCREMENT,
    name        VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    price       INTEGER NOT NULL,
    PRIMARY KEY (id)
);

-- +migrate UP
CREATE TABLE basket (
    id      INTEGER NOT NULL AUTO_INCREMENT,
    item_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    name    VARCHAR(255) NOT NULL,
    price   INTEGER NOT NULL,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE users;
DROP TABLE items;
DROP TABLE basket;