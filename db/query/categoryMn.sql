-- name: CreateCategoryMn :one
INSERT INTO
    "categoryMn" ("CategoryNameMn")
VALUES
    (
        "CategoryNameMn" = sqlc.arg('CategoryNameMn') :: VARCHAR(100)
    ) RETURNING *;

-- name: GetListAllCategoryMn :many
SELECT
    *
FROM
    "categoryMn";

-- name: UpdateByCategoryMn :exec
UPDATE
    "categoryMn"
SET
    "CategoryNameMn" = sqlc.arg('CategoryNameMn') :: VARCHAR(100)
WHERE
    "CategoryMnID" = sqlc.arg('CategoryMnID');

-- name: DeleteFromCategoryMn :exec
DELETE FROM
    "categoryMn"
WHERE
    "CategoryMnID" = sqlc.arg('CategoryMnID');

-- name: GetByIdCategoryMn :one
SELECT
    *
FROM
    "categoryMn"
WHERE
    "CategoryMnID" = sqlc.arg('CategoryMnID')
LIMIT
    1;

-- name: FindByNameCategoryMn :one
SELECT
    *
FROM
    "categoryMn"
WHERE
    "CategoryNameMn" = sqlc.arg('CategoryNameMn')
LIMIT
    1;

-- name: FindByIdCategoryMn :one
SELECT
    *
FROM
    "categoryMn"
WHERE
    "CategoryMnID" - sqlc.arg('CategoryMnID')
LIMIT
    1;