-- +goose Up
ALTER TABLE user ADD COLUMN source_lang TEXT NOT NULL DEFAULT 'ru';
ALTER TABLE user ADD COLUMN target_lang TEXT NOT NULL DEFAULT 'en';

-- +goose Down
ALTER TABLE user DROP COLUMN source_lang;
ALTER TABLE user DROP COLUMN target_lang;