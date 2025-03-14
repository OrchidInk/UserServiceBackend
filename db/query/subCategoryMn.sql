-- name: CreateSubCategoryMn :one
INSERT INTO
    "subCategoryMn" (
        "subCategoryNameMn",
        "CategoryMnID"
    )
VALUES
    (
        sqlc.arg('subCategoryNameMn') :: VARCHAR(100),
        sqlc.arg('CategoryMnID') :: INT
    ) RETURNING *;

-- name: GetListAllSubCategoryMn :many
SELECT
    *
FROM
    "subCategoryMn";

-- name: FindByNameSubCategoryMn :one
SELECT
    *
FROM
    "subCategoryMn"
WHERE
    "subCategoryNameMn" = sqlc.arg('subCategoryNameMn')
LIMIT
    1;

-- name: UpdateBySubCategoryNameMn :one
UPDATE
    "subCategoryMn"
SET
    "subCategoryNameMn" = sqlc.arg('subCategoryNameMn') :: VARCHAR(100)
WHERE
    "SubCategoryIDMn" = sqlc.arg('SubCategoryIDMn') RETURNING *;

-- name: FindBySubCategoryID :one
SELECT
    *
FROM
    "subCategoryMn"
WHERE
    "SubCategoryIDMn" = sqlc.arg('SubCategoryIDMn')
LIMIT
    1;

-- name: DeleteBySubCategoryMn :exec
DELETE FROM
    "subCategoryMn"
WHERE
    "SubCategoryIDMn" = sqlc.arg('SubCategoryIDMn');

-- name: UpdateByCategoryIDMn :one
UPDATE
    "subCategoryMn"
SET
    "CategoryMnID" = sqlc.arg('CategoryMnID') RETURNING *;

-- name: GetProductsBySubCategoryMn :many
SELECT
    p."ProductMnID",
    p."ProductNameMn",
    p."PriceMn",
    p."StockQuantity",
    p."ImagesPathMn"
FROM
    "subCategoryMn" sc
    JOIN "productMn" p ON sc."SubCategoryIDMn" = p."SubCategoryIDMn"
WHERE
    sc."SubCategoryIDMn" = sqlc.arg('SubCategoryIDMn');

-- name: UpdateSubCategoryByIDMn :one
UPDATE
    "subCategoryMn"
SET
    "subCategoryNameMn" = sqlc.arg('subCategoryNameMn'),
    "CategoryMnID" = sqlc.arg('CategoryMnID')
WHERE
    "SubCategoryIDMn" = sqlc.arg('SubCategoryIDMn') RETURNING *;