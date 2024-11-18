-- name: CreateImageMn :one
INSERT INTO
    "imagesMn" ("ProductMnID", "ImagePath")
VALUES
    (
        sqlc.arg('ProductMnID'),
        sqlc.arg('ImagePath')
    ) RETURNING *;

-- name: GetListImagesMn :many
SELECT
    *
FROM
    "imagesMn";

-- name: DeleteImagesMn :exec
DELETE FROM
    "imagesMn"
WHERE
    "Id" = sqlc.arg('Id');

-- name: UpdateImagesMn :one
UPDATE
    "imagesMn"
SET
    "ImagePath" = sqlc.arg('ImagePath')
WHERE
    "Id" = sqlc.arg('Id') RETURNING *;