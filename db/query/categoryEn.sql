-- name: CreateCategoryEn :one
INSERT INTO
    "categoryEn" (
        "CategoryNameEn"
    )
VALUES (
    sqlc.arg('CategoryNameEn') :: VARCHAR(100) 
) RETURNING *;

-- name: GetListByAllCategoryEn :many
SELECT 
    *
FROM
    "categoryEn";

-- name: UpdateCategoryEn :exec
UPDATE
  "categoryEn"
SET
  "CategoryNameEn" = sqlc.arg('CategoryNameEn') :: VARCHAR(100) 
WHERE
  "CategoryEnID" = sqlc.arg('CategoryEnID');

-- name: DeleteByIdCategoryEn :exec
DELETE FROM 
    "categoryEn"
WHERE
    "CategoryEnID" = sqlc.arg('CategoryEnID');

-- name: FindByCategoryEnId :one
SELECT
    *
FROM
    "categoryEn"
WHERE
    "CategoryEnID" = sqlc.arg('CategoryEnID')
LIMIT 1;

-- name: FindByNameCategoryEn :one
SELECT 
    *
FROM
    "categoryEn"
WHERE
    "CategoryNameEn" = sqlc.arg('CategoryNameEn') :: VARCHAR(100);