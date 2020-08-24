-- +migrate Up
CREATE TABLE ordereds (
    created_at datetime default current_timestamp,
    item_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    number  INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY (item_id)
        REFERENCES items(id)
        ON DELETE CASCADE,
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);