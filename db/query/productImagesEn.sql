-- name: CreateImagesEn :one
INSERT INTO
    "imagesEn" (
        "ProductEnID",
        "ImagePath"
    )
VALUES
    (
        sqlc.arg('ProductEnID'),
        sqlc.arg('ImagePath')
    ) RETURNING *;

-- name: GetListImagesEn :many
SELECT
    *
FROM
    "imagesEn";

-- name: UpdateImagesEn :one
UPDATE
    "imagesEn"
SET
    "ImagePath" = sqlc.arg('ImagePath')
WHERE
    "Id" = sqlc.arg('Id') RETURNING *;

-- name: DeleteImagesEn :exec
DELETE FROM
    "imagesEn"
WHERE
    "Id" = sqlc.arg('Id');