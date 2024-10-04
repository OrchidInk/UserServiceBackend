-- name: CreateSubCategoryEn :one
INSERT INTO
    "subCategoryEn" (
        "subCategoryNameEn",
        "CategoryEnID"
    )
VALUES
    (
        "subCategoryNameEN" = sqlc.arg('subCategoryNameEN') :: VARCHAR(100),
        "CategoryEnID" = sqlc.arg('CategoryEnID') :: INT
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

-- name: GetProductsBySubCategoryEn :many
SELECT 
    p."ProductEnID",
    p."ProductNameEn",
    p."PriceEn",
    p."StockQuantity",
    p."ImagesPathEn"
FROM
    "subCategoryEn"  sc
JOIN 
    "productEn"  p ON sc."subCategoryIDEn" = p."subCategoryIDEn"
WHERE 
    sc."subCategoryIDEn" = $1;
