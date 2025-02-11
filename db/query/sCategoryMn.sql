-- name: CreateSCategoryMn :one
INSERT INTO
    "sCategoryMn" (
        "sCategoryNameMn",
        "SubCategoryIDMn"
    )
VALUES
    (
        sqlc.arg('sCategoryNameMn'),
        sqlc.arg('SubCategoryIDMn')
    ) RETURNING *;

-- name: GetAllSCategoriesMn :many
SELECT
    *
FROM
    "sCategoryMn";

-- name: UpdateSCategoryNameMn :one
UPDATE "sCategoryMn"
SET "sCategoryNameMn" = sqlc.arg('sCategoryNameMn')
WHERE "sCategoryIdMn" = sqlc.arg('sCategoryIdMn') RETURNING *;


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
    "sCategoryNameMn" = sqlc.arg('sCategoryNameMn')
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
    "sCategoryNameMn" = sqlc.arg('sCategoryNameMn'),
    "SubCategoryIDMn" = sqlc.arg('SubCategoryIDMn')
WHERE
    "sCategoryIdMn" = sqlc.arg('sCategoryIdMn') RETURNING *;