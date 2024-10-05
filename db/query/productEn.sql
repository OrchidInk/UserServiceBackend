-- name: CreateProductEn :one
INSERT INTO
    "productEn" (
        "ProductNameEn",
        "ImageID",
        "subCategoryIDEn",
        "PriceEn",
        "StockQuantity",
        "ImagesPathEn"
    )
VALUES
    (
        "ProductNameEn" = sqlc.arg('ProductNameEn'),
        "ImageID" = sqlc.arg('ImageID'),
        "subCategoryIDEn" = sqlc.arg('subCategoryIDEn'),
        "PriceEn" = sqlc.arg('PriceEn'),
        "StockQuantity" = sqlc.arg('StockQuantity'),
        "ImagesPathEn" = sqcl.arg('ImagesPathEn')

    ) RETURNING *;

-- name: GetListProductEn :many
SELECT
    *
FROM
    "productEn"
ORDER BY
    "Created_At" DESC;

-- name: UpdateByProductEnPrice :one
UPDATE "productEn"
SET
    "PriceEn" = sqlc.arg ('PriceEn')
WHERE
    "ProductEnID" = sqlc.arg ('ProductEnID') RETURNING *;

-- name: UpdateByProductEnStockQuantity :one
UPDATE "productEn"
SET
    "StockQuantity" = sqlc.arg ('StockQuantity')
WHERE
    "ProductEnID" = sqlc.arg ('ProductEnID') RETURNING *;

-- name: UpdateByProductEnImagePath :one
UPDATE "productEn"
SET
    "ImagesPathEn" = sqlc.arg ('ImagesPathEn')
WHERE
    "ProductEnID" = sqlc.arg ('ProductEnID') RETURNING *;

-- name: DeleteByProductEnId :exec
DELETE FROM "productEn"
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
