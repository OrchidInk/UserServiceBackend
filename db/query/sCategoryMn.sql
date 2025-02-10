-- name: CreateSCategoryMn :one
INSERT INTO
    "sCategoryMn" (
        "sCategoryName",
        "SubCategoryIDMn"
    )
VALUES
    (
        sqlc.arg('sCategoryName') :: VARCHAR(100),
        sqlc.arg('SubCategoryIDMn') :: INT
    ) RETURNING *;

-- name: GetAllSCategoriesMn :many
SELECT
    *
FROM
    "sCategoryMn";

-- name: UpdateSCategoryNameMn :one
UPDATE
    "sCategoryMn"
SET
    "sCategoryName" = sqlc.arg('sCategoryName') :: VARCHAR(100)
WHERE
    "sCategoryIdMn" = sqlc.arg('sCategoryIdMn') RETURNING *;

-- name: DeleteSCategoryMn :exec
DELETE FROM
    "sCategoryMn"
WHERE
    "sCategoryIdMn" = sqlc.arg('sCategoryIdMn');

-- name: FindBySCategoryIdMn :one
SELECT
    *
FROM
    "sCategoryMn"
WHERE
    "sCategoryIdMn" = sqlc.arg('sCategoryIdMn')
LIMIT 1;

-- name: FindBySCategoryNameMn :one
SELECT  
    *
FROM
    "sCategoryMn"
WHERE
    "sCategoryName" = sqlc.arg('sCategoryName')
LIMIT
 1;

-- name: GetProductBySCategoriesMn :many
SELECT
    p."ProductMnID",
    p."ProductNameMn",
    p."PriceMn",
    p."StockQuantity",
    p."ImagesPathMn"
FROM
    "sCategoryMn" sc
    JOIN "productMn" p ON sc."sCategoryIdMn" = p."sCategoryIdMn"
WHERE
    sc."sCategoryIdMn" = sqlc.arg('sCategoryIdMn');

-- name: UpdateSCategoryByIdMn :exec
UPDATE  
    "sCategoryMn"
SET
    "sCategoryName" = sqlc.arg('sCategoryName'),
    "SubCategoryIDMn" = sqlc.arg('SubCategoryIDMn')
WHERE
    "sCategoryIdMn" = sqlc.arg('sCategoryIdMn') RETURNING *;