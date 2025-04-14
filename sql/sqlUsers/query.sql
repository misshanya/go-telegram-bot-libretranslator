-- name: CreateUser :one
INSERT INTO user (tg_id)
VALUES (?)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM user
WHERE tg_id = ?
LIMIT 1;

-- name: ChangeLangAutodetect :one
UPDATE user
SET lang_autodetect = NOT lang_autodetect
WHERE tg_id = ?
RETURNING *;

-- name: GetSourceLang :one
SELECT source_lang
FROM user
WHERE tg_id = ?
LIMIT 1;

-- name: GetTargetLang :one
SELECT target_lang
FROM user
WHERE tg_id = ?
LIMIT 1;

-- name: SetSourceLang :exec
UPDATE user
SET source_lang = ?
WHERE tg_id = ?;

-- name: SetTargetLang :exec
UPDATE user
SET target_lang = ?
WHERE tg_id = ?;