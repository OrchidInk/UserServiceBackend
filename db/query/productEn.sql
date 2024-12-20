-- name: CreateProductEn :one
INSERT INTO
    "productEn" (
        "ProductNameEn",
        "subCategoryIDEn",
        "PriceEn",
        "StockQuantity",
        "ImagesPathEn"
    )
VALUES
    (
        sqlc.arg('ProductNameEn'),
        sqlc.arg('subCategoryIDEn'),
        sqlc.arg('PriceEn'),
        sqlc.arg('StockQuantity'),
        sqlc.arg('ImagesPathEn')
    ) RETURNING *;

-- name: GetListProductEn :many
SELECT
    *
FROM
    "productEn"
ORDER BY
    "Created_At" DESC;

-- name: UpdateByProductEnPrice :one
UPDATE
    "productEn"
SET
    "PriceEn" = sqlc.arg ('PriceEn')
WHERE
    "ProductEnID" = sqlc.arg ('ProductEnID') RETURNING *;

-- name: UpdateByProductEnStockQuantity :one
UPDATE
    "productEn"
SET
    "StockQuantity" = sqlc.arg ('StockQuantity')
WHERE
    "ProductEnID" = sqlc.arg ('ProductEnID') RETURNING *;

-- name: UpdateByProductEnImagePath :one
UPDATE
    "productEn"
SET
    "ImagesPathEn" = sqlc.arg ('ImagesPathEn')
WHERE
    "ProductEnID" = sqlc.arg ('ProductEnID') RETURNING *;

-- name: DeleteByProductEnId :exec
DELETE FROM
    "productEn"
WHERE
    "ProductEnID" = sqlc.arg ('ProductEnID');

-- name: FilterByProductEnName :many
SELECT
    *
FROM
    "productEn"
WHERE
    "ProductEnName" ILIKE '%' || sqlc.arg ('ProductEnName') || '%' -- Case-insensitive search for partial match
ORDER BY
    "Created_At" DESC;

-- name: FindByProductEnID :one
SELECT
    *
FROM
    "productEn"
WHERE
    "ProductEnID" = sqlc.arg('ProductEnID')
LIMIT
    1;

-- name: DeductStockQuantityByProductEnID :one
UPDATE
    "productEn"
SET
    "StockQuantity" = "StockQuantity" - sqlc.arg('quantityPurchased')
WHERE
    "ProductEnID" = sqlc.arg('ProductEnID')
    AND "StockQuantity" >= sqlc.arg('quantityPurchased') RETURNING *;

-- name: GetProductWithDetailsEn :many
SELECT
    p."ProductEnID",
    p."ProductNameEn",
    p."subCategoryIDEn",
    p."PriceEn",
    p."StockQuantity",
    p."ImagesPathEn",
    p."Created_At",
    p."Updated_At",
    d."detailEnId",
    d."ChoiceName",
    d."ChoiceValue"
FROM
    "productEn" p
    LEFT JOIN "detailEn" d ON p."ProductEnID" = d."ProductEnID"
ORDER BY
    p."ProductEnID",
    d."detailEnId";

-- name: FindProductWithDetailsByIDEn :many
SELECT
    p."ProductEnID",
    p."ProductNameEn",
    p."subCategoryIDEn",
    p."PriceEn",
    p."StockQuantity",
    p."ImagesPathEn",
    p."Created_At",
    p."Updated_At",
    d."detailEnId",
    d."ChoiceName",
    d."ChoiceValue"
FROM
    "productEn" p
    LEFT JOIN "detailEn" d ON p."ProductEnID" = d."ProductEnID"
WHERE
    p."ProductEnID" = sqlc.arg('ProductEnID');