-- name: CreateImages :one
INSERT INTO
    "images" ("ImagePath")
VALUES
    (sqlc.arg('ImagePath') :: TEXT) RETURNING *;

-- name: FindByListImages :many
SELECT
    *
FROM
    "images"
ORDER BY
    "Created_At" DESC;

-- name: FindByImageUrl :one
SELECT
    *
FROM
    "images"
WHERE
    "ImagePath" = sqlc.arg('ImagePath') :: TEXT
LIMIT
    1;

-- name: DeleteByImages :exec
DELETE FROM
    "images"
WHERE
    "Id" = sqlc.arg('Id');

-- name: FindByImageId :one
SELECT
    *
FROM
    "images"
WHERE
    "Id" = sqlc.arg('Id')
LIMIT
    1;