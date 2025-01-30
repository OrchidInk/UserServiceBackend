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


-- name: UpdateProductEn :one
UPDATE 
    "productEn"
SET
    "ProductNameEn" = sqlc.arg('ProductNameEn'),
    "SubCategoryIDEn" = sqlc.arg('SubCategoryIDEn'),
    "PriceEn" = sqlc.arg('PriceEn'),
    "StockQuantity" = sqlc.arg('StockQuantity'),
    "ImagesPathEn" = sqlc.arg('ImagesPathEn'),
    "DescriptionEn" = sqlc.arg('DescriptionEn'),
    "BrandEn" = sqlc.arg('BrandEn'),
    "ManufacturedCountryEn" = sqlc.arg('ManufacturedCountryEn'),
    "ColorEn" = sqlc.arg('ColorEn'),
    "SizeEn" = sqlc.arg('SizeEn'),
    "PenOutputEn" = sqlc.arg('PenOutputEn'),
    "FeaturesEn" = sqlc.arg('FeaturesEn'),
    "MaterialEn" = sqlc.arg('MaterialEn'),
    "StapleSizeEn" = sqlc.arg('StapleSizeEn'),
    "CapacityEn" = sqlc.arg('CapacityEn'),
    "WeightEn" = sqlc.arg('WeightEn'),
    "ThicknessEn" = sqlc.arg('ThicknessEn'),
    "PackagingEn" = sqlc.arg('PackagingEn'),
    "UsageEn" = sqlc.arg('UsageEn'),
    "InstructionsEn" = sqlc.arg('InstructionsEn'),
    "ProductCodeEn" = sqlc.arg('ProductCodeEn'),
    "CostPriceEn" = sqlc.arg('CostPriceEn'),
    "RetailPriceEn" = sqlc.arg('RetailPriceEn'),
    "WarehouseStockEn" = sqlc.arg('WarehouseStockEn')
    WHERE
    "ProductEnID" = sqlc.arg('ProductEnID') RETURNING *;
