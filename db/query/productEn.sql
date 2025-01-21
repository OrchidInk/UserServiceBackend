-- name: CreateProductEn :one
INSERT INTO
    "productEn" (
        "ProductNameEn",
        "SubCategoryIDEn",
        "PriceEn",
        "StockQuantity",
        "ImagesPathEn",
        "DescriptionEn",
        "BrandEn",
        "ManufacturedCountryEn",
        "ColorEn",
        "SizeEn",
        "PenOutputEn",
        "FeaturesEn",
        "MaterialEn",
        "StapleSizeEn",
        "CapacityEn",
        "WeightEn",
        "ThicknessEn",
        "PackagingEn",
        "ProductCodeEn",
        "CostPriceEn",
        "RetailPriceEn",
        "WarehouseStockEn"
    )
VALUES
    (
        sqlc.arg('ProductNameEn'),
        sqlc.arg('SubCategoryIDEn'),
        sqlc.arg('PriceEn'),
        sqlc.arg('StockQuantity'),
        sqlc.arg('ImagesPathEn'),
        sqlc.arg('DescriptionEn'),
        sqlc.arg('BrandEn'),
        sqlc.arg('ManufacturedCountryEn'),
        sqlc.arg('ColorEn'),
        sqlc.arg('SizeEn'),
        sqlc.arg('PenOutputEn'),
        sqlc.arg('FeaturesEn'),
        sqlc.arg('MaterialEn'),
        sqlc.arg('StapleSizeEn'),
        sqlc.arg('CapacityEn'),
        sqlc.arg('WeightEn'),
        sqlc.arg('ThicknessEn'),
        sqlc.arg('PackagingEn'),
        sqlc.arg('ProductCodeEn'),
        sqlc.arg('CostPriceEn'),
        sqlc.arg('RetailPriceEn'),
        sqlc.arg('WarehouseStockEn')
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

-- name: UpdateByEnImagePath :one
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
    "ProductNameEn" ILIKE '%' || sqlc.arg ('ProductNameEn') || '%' -- Case-insensitive search for partial match
ORDER BY
    "Created_At" DESC;

-- name: FindByProductIdEn :one
SELECT
    *
FROM
    "productEn"
WHERE
    "ProductEnID" = sqlc.arg('ProductEnID')
LIMIT
    1;

-- name: DeductSockQuantityByProductEnID :one
UPDATE
    "productEn"
SET
    "StockQuantity" = "StockQuantity" - sqlc.arg('quantityPurchased')
WHERE
    "ProductEnID" = sqlc.arg('ProductEnID')
    AND "StockQuantity" >= sqlc.arg('quantityPurchased') RETURNING *;

-- name: UpdateProductEnSubCategory :one
UPDATE
    "productEn"
SET
    "SubCategoryIDEn" = sqlc.arg('SubCateogoryIDEn')
WHERE
    "ProductEnID" = sqlc.arg('ProductEnID') RETURNING *;