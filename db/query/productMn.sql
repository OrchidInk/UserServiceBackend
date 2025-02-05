-- name: CreateProductMn :one
INSERT INTO
    "productMn" (
        "ProductNameMn",
        "SubCategoryIDMn",
        "PriceMn",
        "StockQuantity",
        "ImagesPathMn",
        "DescriptionMn",
        "BrandMn",
        "ManufacturedCountryMn",
        "ColorMn",
        "SizeMn",
        "PenOutputMn",
        "FeaturesMn",
        "MaterialMn",
        "StapleSizeMn",
        "CapacityMn",
        "WeightMn",
        "ThicknessMn",
        "PackagingMn",
        "ProductCodeMn",
        "CostPriceMn",
        "RetailPriceMn",
        "WarehouseStockMn"
    )
VALUES
    (
        sqlc.arg('ProductNameMn'),
        sqlc.arg('SubCategoryIDMn'),
        sqlc.arg('PriceMn'),
        sqlc.arg('StockQuantity'),
        sqlc.arg('ImagesPathMn'),
        sqlc.arg('DescriptionMn'),
        sqlc.arg('BrandMn'),
        sqlc.arg('ManufacturedCountryMn'),
        sqlc.arg('ColorMn'),
        sqlc.arg('SizeMn'),
        sqlc.arg('PenOutputMn'),
        sqlc.arg('FeaturesMn'),
        sqlc.arg('MaterialMn'),
        sqlc.arg('StapleSizeMn'),
        sqlc.arg('CapacityMn'),
        sqlc.arg('WeightMn'),
        sqlc.arg('ThicknessMn'),
        sqlc.arg('PackagingMn'),
        sqlc.arg('ProductCodeMn'),
        sqlc.arg('CostPriceMn'),
        sqlc.arg('RetailPriceMn'),
        sqlc.arg('WarehouseStockMn')
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
    "ProductNameMn" ILIKE '%' || sqlc.arg ('ProductNameMn') || '%' -- Case-insensitive search for partial match
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

-- name: DeductSockQuantityByProductMnID :one
UPDATE
    "productMn"
SET
    "StockQuantity" = "StockQuantity" - sqlc.arg('quantityPurchased')
WHERE
    "ProductMnID" = sqlc.arg('ProductMnID')
    AND "StockQuantity" >= sqlc.arg('quantityPurchased') RETURNING *;

-- name: UpdateProductMnSubCategory :one
UPDATE
    "productMn"
SET
    "SubCategoryIDMn" = sqlc.arg('SubCategoryIDMn')
WHERE
    "ProductMnID" = sqlc.arg('ProductMnID') RETURNING *;


-- name: UpdateProductMn :one
UPDATE
    "productMn"
SET
    "ProductNameMn" = sqlc.arg('ProductNameMn'),
    "SubCategoryIDMn" = sqlc.arg('SubCategoryIDMn'),
    "PriceMn" = sqlc.arg('PriceMn'),
    "StockQuantity" = sqlc.arg('StockQuantity'),
    "ImagesPathMn" = sqlc.arg('ImagesPathMn'),
    "DescriptionMn" = sqlc.arg('DescriptionMn'),
    "BrandMn" = sqlc.arg('BrandMn'),
    "ManufacturedCountryMn" = sqlc.arg('ManufacturedCountryMn'),
    "ColorMn" = sqlc.arg('ColorMn'),
    "SizeMn" = sqlc.arg('SizeMn'),
    "PenOutputMn" = sqlc.arg('PenOutputMn'),
    "FeaturesMn" = sqlc.arg('FeaturesMn'),
    "StapleSizeMn" = sqlc.arg('StapleSizeMn'),
    "CapacityMn" = sqlc.arg('CapacityMn'),
    "WeightMn" = sqlc.arg('WeightMn'),
    "ThicknessMn" = sqlc.arg('ThicknessMn'),
    "PackagingMn" = sqlc.arg('PackagingMn'),
    "UsageMn" = sqlc.arg('UsageMn'),
    "InstructionsMn" = sqlc.arg('InstructionsMn'),
    "ProductCodeMn" = sqlc.arg('ProductCodeMn'),
    "CostPriceMn" = sqlc.arg('CostPriceMn'),
    "RetailPriceMn" = sqlc.arg('RetailPriceMn'),
    "WarehouseStockMn" = sqlc.arg('WarehouseStockMn')
WHERE
    "ProductMnID" = sqlc.arg('ProductMnID') RETURNING *;
