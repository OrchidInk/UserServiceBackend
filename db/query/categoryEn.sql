-- name: CreateCategoryEn :one
INSERT INTO
    "categoryEn" ("CategoryNameEn")
VALUES
    (sqlc.arg ('CategoryNameEn')) RETURNING *;

-- name: GetListByAllCategoryEn :many
SELECT
    *
FROM
    "categoryEn";

-- name: UpdateCategoryEn :one
UPDATE
    "categoryEn"
SET
    "CategoryNameEn" = sqlc.arg('CategoryNameEn') :: VARCHAR(100)
WHERE
    "CategoryEnID" = sqlc.arg('CategoryEnID') :: INT RETURNING *;

-- name: DeleteCategoryById :exec
DELETE FROM
    "categoryEn"
WHERE
    "CategoryEnID" = sqlc.arg('CategoryEnID') :: INT;

-- name: FindByCategoryEnId :one
SELECT
    *
FROM
    "categoryEn"
WHERE
    "CategoryEnID" = sqlc.arg ('CategoryEnID')
LIMIT
    1;

-- name: FindByNameCategoryEn :one
SELECT
    *
FROM
    "categoryEn"
WHERE
    "CategoryNameEn" = sqlc.arg ('CategoryNameEn')
LIMIT
    1;

-- name: FindByIdCategoryEn :one
SELECT
    *
FROM
    "categoryEn"
WHERE
    "CategoryEnID" = sqlc.arg ('CategoryEnID')
LIMIT
    1;

-- name: GetCategoriesWithSubCategories :many
SELECT
    c."CategoryEnID",
    c."CategoryNameEn",
    sc."subCategoryIDEn",
    sc."subCategoryNameEn"
FROM
    "categoryEn" c
LEFT JOIN
    "subCategoryEn" sc ON c."CategoryEnID" = sc."CategoryEnID"
ORDER BY
    c."CategoryEnID";
