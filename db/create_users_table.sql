-- +migrate Up
CREATE TABLE users (
    id          INTEGER NOT NULL AUTO_INCREMENT,
    email       VARCHAR(255) UNIQUE NOT NULL,
    password    VARCHAR(255) NOT NULL,
    token       VARCHAR(255) NOT NULL,
    validation  BOOLEAN DEFAULT 0,
    PRIMARY KEY (id)
);