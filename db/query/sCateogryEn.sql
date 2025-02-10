-- name: CreateSCategoryEn :one
INSERT INTO
    "sCategoryEn" (
        "sCategoryNameEn",
        "SubCategoryIDEn"   
    ) 
VALUES
    (
        sqlc.arg('sCategoryNameEn') :: VARCHAR(100),
        sqlc.arg('SubCategoryIDEn') :: INT
    ) RETURNING *;

-- name: GetAllSCategoryEn :many
SELECT
    *
FROM
    "sCategoryEn";

-- name: UpdateSCategoryNameEn :one
UPDATE
    "sCategoryEn"
SET
    "sCategoryNameEn" = sqlc.arg('sCategoryNameEn') :: VARCHAR(100)
WHERE
    "sCategoryIdEn" = sqlc.arg('sCategoryIdEn') RETURNING *;

-- name: DeleteSCategoryEn :exec
DELETE FROM
    "sCategoryEn"
WHERE
    "sCategoryIdEn" = sqlc.arg('sCategoryIdEn');

-- name: FindByNameSCategoryNameEn :one
SELECT
    *
FROM
    "sCategoryEn"
WHERE
    "sCategoryNameEn" = sqlc.arg('sCategoryNameEn')
LIMIT 1;

-- name: FindBySCategoryIdEn :one
SELECT
    *
FROM
    "sCategoryEn"
WHERE
    "sCategoryIdEn" = sqlc.arg('sCategoryIdEn')
LIMIT 1;

-- name: GetProductBySCategoryEn :many
SELECT
    p."ProductEnID",
    p."ProductNameEn",
    p."PriceEn",
    p."StockQuantity",
    p."ImagesPathEn"
FROM
    "sCategoryEn" sc
    JOIN "productEn" p ON sc."sCategoryIdEn" = p."sCategoryIdEn"
WHERE
    sc."sCategoryIdEn" = sqlc.arg('sCategoryIdEn');

-- name: UpdateSCategoryByIdEn :exec
UPDATE
    "sCategoryEn"
SET
    "sCategoryNameEn" = sqlc.arg('sCategoryNameEn'),
    "SubCategoryIDEn" = sqlc.arg('SubCategoryIDEn')
WHERE
    "sCategoryIdEn" = sqlc.arg('sCategoryIdEn') RETURNING *;