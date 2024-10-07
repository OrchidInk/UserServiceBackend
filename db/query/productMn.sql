-- name: CreateProductMn :one
INSERT INTO
    "productMn" (
        "ProductNameMn",
        "subCategoryIDMn",
        "PriceMn",
        "StockQuantity",
        "ImagesPathMn"
    )
VALUES
    (
        sqlc.arg('ProductNameMn'),
        sqlc.arg('subCategoryIDMn'),
        sqlc.arg('PriceMn'),
        sqlc.arg('StockQuantity'),
        sqlc.arg('ImagesPathMn')
    ) RETURNING *;

-- name: GetListProductMn :many
SELECT
    *
FROM
    "productMn"
ORDER BY
    "Created_At" DESC;

-- name: UpdateByProductMnPrice :one
UPDATE
    "productMn"
SET
    "PriceMn" = sqlc.arg ('PriceMn')
WHERE
    "ProductMnID" = sqlc.arg ('ProductMnID') RETURNING *;

-- name: UpdateByProductMnStockQuantity :one
UPDATE
    "productMn"
SET
    "StockQuantity" = sqlc.arg ('StockQuantity')
WHERE
    "ProductMnID" = sqlc.arg ('ProductMnID') RETURNING *;

-- name: UpdateByMnImagePath :one
UPDATE
    "productMn"
SET
    "ImagesPathMn" = sqlc.arg ('ImagesPathMn')
WHERE
    "ProductMnID" = sqlc.arg ('ProductMnID') RETURNING *;

-- name: DeleteByProductMnId :exec
DELETE FROM
    "productMn"
WHERE
    "ProductMnID" = sqlc.arg ('ProductMnID');

-- name: FilterByProductMnName :many
SELECT
    *
FROM
    "productMn"
WHERE
    "ProductMnName" ILIKE '%' || sqlc.arg ('ProductMnName') || '%' -- Case-insensitive search for partial match
ORDER BY
    "Created_At" DESC;

-- name: FindByProductIdMn :one
SELECT
    *
FROM
    "productMn"
WHERE
    "ProductMnID" = sqlc.arg('ProductMnID')
LIMIT
    1;