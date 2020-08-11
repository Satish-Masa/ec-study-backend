-- +migrate Up
CREATE TABLE carts (
    item_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (item_id)
        REFERENCES items(id)
        ON DELETE CASCADE,
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);