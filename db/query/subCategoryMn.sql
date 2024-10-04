-- name: CreateSubCategoryMn :one
INSERT INTO
    "subCategoryMn" (
        "subCategoryNameMn",
        "CategoryMnID"
    )
VALUES
    (
        "subCategoryNameMn" = sqlc.arg('subCategoryNameMn') :: VARCHAR(100),
        "CategoryMnID" = sqlc.arg('CategoryMnID')
    ) RETURNING *;

-- name: GetListAllSubCategoryMn :many
SELECT
    *
FROM
    "subCategoryMn";

-- name: UpdateBySubCategoryNameMn :one
UPDATE
    "subCategoryMn"
SET
    "subCategoryNameMn" = sqlc.arg('subCategoryNameMn') :: VARCHAR(100)
WHERE
    "subCategoryIDMn" = sqlc.arg('subCategoryIDMn') RETURNING *;

-- name: FindBySubCategoryID :one
SELECT
    *
FROM
    "subCategoryMn"
WHERE
    "subCategoryIDMn" = sqlc.arg('subCategoryIDMn')
LIMIT
    1;

-- name: DeleteBySubCategoryMn :exec
DELETE FROM
    "subCategoryMn"
WHERE
    "subCategoryIDMn" = sqlc.arg('subCategoryIDMn');

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
    "subCategoryMn"  sc
JOIN 
    "productMn"  p ON sc."subCategoryIDMn" = p."subCategoryIDMn"
WHERE 
    sc."subCategoryIDMn" = $1;
