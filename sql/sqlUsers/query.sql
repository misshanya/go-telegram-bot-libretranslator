-- name: CreateUser :one
INSERT INTO user (
    tg_id
) VALUES (
    ?
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM user
WHERE tg_id = ? LIMIT 1;

-- name: ChangeLangAutodetect :one
UPDATE user
SET lang_autodetect = NOT lang_autodetect
WHERE tg_id = ?
RETURNING *;