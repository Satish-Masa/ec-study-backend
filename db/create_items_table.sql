-- +migrate Up
CREATE TABLE items (
    id          INTEGER NOT NULL AUTO_INCREMENT,
    name        VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    price       INTEGER NOT NULL,
    stock       INTEGER NOT NULL,
    PRIMARY KEY (id)
);