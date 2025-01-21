-- name: CreateSubCategoryEn :one
INSERT INTO
    "subCategoryEn" (
        "subCategoryNameEn",
        -- Correct casing here
        "CategoryEnID"
    )
VALUES
    (
        sqlc.arg('subCategoryNameEn') :: VARCHAR(100),
        sqlc.arg('CategoryEnID') :: INT
    ) RETURNING *;

-- name: GetListAllSubCategoriesEn :many
SELECT
    *
FROM
    "subCategoryEn";

-- name: UpdateSubCategoryNameEn :one
UPDATE
    "subCategoryEn"
SET
    "subCategoryNameEn" = sqlc.arg('subCategoryNameEn') :: VARCHAR(100)
WHERE
    "subCategoryIDEn" = sqlc.arg('subCategoryIDEn') RETURNING *;

-- name: DeleteSubCategoryEn :exec
DELETE FROM
    "subCategoryEn"
WHERE
    "subCategoryIDEn" = sqlc.arg('subCategoryIDEn');

-- name: FindBySubCategoryIDEn :one
SELECT
    *
FROM
    "subCategoryEn"
WHERE
    "subCategoryIDEn" = sqlc.arg('subCategoryIDEn')
LIMIT
    1;

-- name: FindByNameSubCategoryEn :one
SELECT
    *
FROM
    "subCategoryEn"
WHERE
    "subCategoryNameEn" = sqlc.arg('subCategoryNameEn')
LIMIT
    1;

-- name: GetProductsBySubCategoryEn :many
SELECT
    p."ProductEnID",
    p."ProductNameEn",
    p."PriceEn",
    p."StockQuantity",
    p."ImagesPathEn"
FROM
    "subCategoryEn" sc
    JOIN "productEn" p ON sc."subCategoryIDEn" = p."subCategoryIDEn"
WHERE
    sc."subCategoryIDEn" = sqlc.arg('subCategoryIDEn');

-- name: UpdateSubCategoryByIDEn :one
UPDATE
    "subCategoryEn"
SET
    "subCategoryNameEn" = sqlc.arg('subCategoryNameEn'),
    "CategoryEnID" = sqlc.arg('CategoryEnID')
WHERE
    "subCategoryIDEn" = sqlc.arg('subCategoryIDEn') RETURNING *;