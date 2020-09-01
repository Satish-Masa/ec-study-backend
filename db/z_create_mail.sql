-- +migrate Up
CREATE TABLE mails (
    user_id INTEGER NOT NULL,
    created_at datetime default current_timestamp,   
    token       VARCHAR(255) NOT NULL,
    validation  BOOLEAN DEFAULT 0,
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);