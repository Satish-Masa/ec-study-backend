-- +migrate Up
CREATE TABLE items (
    id          INTEGER NOT NULL AUTO_INCREMENT,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    price       INTEGER NOT NULL,
    stock       INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);