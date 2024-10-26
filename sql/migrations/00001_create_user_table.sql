-- +goose Up
CREATE TABLE IF NOT EXISTS user (
    id INTEGER NOT NULL PRIMARY KEY,
    tg_id BIGINT UNIQUE NOT NULL,
    lang_autodetect BOOLEAN NOT NULL DEFAULT TRUE,
    registered_at timestamp DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE user;