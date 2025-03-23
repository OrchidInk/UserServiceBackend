-- name: CreateProductMn :one
INSERT INTO
    "productMn" (
        "ProductNameMn",
        "sCategoryIdMn",
        "PriceMn",
        "StockQuantity",
        "ImagesPathMn",
        "DescriptionMn",
        "BrandMn",
        "ManufacturedCountryMn",
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
        sqlc.arg('sCategoryIdMn'),
        sqlc.arg('PriceMn'),
        sqlc.arg('StockQuantity'),
        sqlc.arg('ImagesPathMn'),
        sqlc.arg('DescriptionMn'),
        sqlc.arg('BrandMn'),
        sqlc.arg('ManufacturedCountryMn'),
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
    "sCategoryIdMn" = sqlc.arg('sCategoryIdMn')
WHERE
    "ProductMnID" = sqlc.arg('ProductMnID') RETURNING *;
-- name: UpdateSProductMn :one
UPDATE
    "productMn"
SET
    "StockQuantity" = sqlc.arg('StockQuantity')
WHERE
    "ProductMnID" = sqlc.arg('ProductMnID') RETURNING *;

-- name: UpdateProductMn :one
UPDATE
    "productMn"
SET
    "ProductNameMn" = sqlc.arg('ProductNameMn'),
    "sCategoryIdMn" = sqlc.arg('sCategoryIdMn'),
    "PriceMn" = sqlc.arg('PriceMn'),
    "StockQuantity" = sqlc.arg('StockQuantity'),
    "ImagesPathMn" = sqlc.arg('ImagesPathMn'),
    "DescriptionMn" = sqlc.arg('DescriptionMn'),
    "BrandMn" = sqlc.arg('BrandMn'),
    "ManufacturedCountryMn" = sqlc.arg('ManufacturedCountryMn'),
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

-- name: GetProductMnWithAllColorsAndSizes :many
SELECT 
    p."ProductMnID",
    p."ProductNameMn",
    p."sCategoryIdMn",
    p."PriceMn",
    p."StockQuantity",
    p."ImagesPathMn",
    p."DescriptionMn",
    p."BrandMn",
    p."ManufacturedCountryMn",
    p."PenOutputMn",
    p."FeaturesMn",
    p."MaterialMn",
    p."StapleSizeMn",
    p."CapacityMn",
    p."WeightMn",
    p."ThicknessMn",
    p."PackagingMn",
    p."UsageMn",
    p."InstructionsMn",
    p."ProductCodeMn",
    p."CostPriceMn",
    p."RetailPriceMn",
    p."WarehouseStockMn",
    p."Created_At",
    p."Updated_At",

    COALESCE(
        ARRAY_AGG(DISTINCT c."ColorId") FILTER (WHERE c."ColorId" IS NOT NULL),
        '{}'
    ) AS "ColorIds",
    COALESCE(
        ARRAY_AGG(DISTINCT c."Color") FILTER (WHERE c."Color" IS NOT NULL),
        '{}'
    ) AS "ColorNames",

    COALESCE(
        ARRAY_AGG(DISTINCT s."SizeId") FILTER (WHERE s."SizeId" IS NOT NULL),
        '{}'
    ) AS "SizeIds",
    COALESCE(
        ARRAY_AGG(DISTINCT s."Size") FILTER (WHERE s."Size" IS NOT NULL),
        '{}'
    ) AS "SizeNames",
    COALESCE(
        ARRAY_AGG(DISTINCT pi."imagesPath") FILTER (WHERE pi."ImagesPath" IS NOT NULL),
        '{}'
    ) AS "ImagesPath"

FROM "productMn" p
LEFT JOIN "productMn_colors" pc ON p."ProductMnID" = pc."ProductMnID"
LEFT JOIN "productMn_sizes" ps ON p."ProductMnID" = ps."ProductMnID"
LEFT JOIN "Color" c ON pc."ColorId" = c."ColorId"
LEFT JOIN "Size" s ON ps."SizeId" = s."SizeId"
LEFT JOIN "productImagesMn" pi ON p."ProductMnID" = pi."ProductMnID"
GROUP BY p."ProductMnID";

-- name: GetProductMnWithAllColorsAndSizesByID :one
SELECT 
    p."ProductMnID",
    p."ProductNameMn",
    p."sCategoryIdMn",
    p."PriceMn",
    p."StockQuantity",
    p."ImagesPathMn",
    p."DescriptionMn",
    p."BrandMn",
    p."ManufacturedCountryMn",
    p."PenOutputMn",
    p."FeaturesMn",
    p."MaterialMn",
    p."StapleSizeMn",
    p."CapacityMn",
    p."WeightMn",
    p."ThicknessMn",
    p."PackagingMn",
    p."UsageMn",
    p."InstructionsMn",
    p."ProductCodeMn",
    p."CostPriceMn",
    p."RetailPriceMn",
    p."WarehouseStockMn",
    p."Created_At",
    p."Updated_At",

    COALESCE(
        ARRAY_AGG(DISTINCT c."ColorId") FILTER (WHERE c."ColorId" IS NOT NULL),
        '{}'
    ) AS "ColorIds",
    COALESCE(
        ARRAY_AGG(DISTINCT c."Color") FILTER (WHERE c."Color" IS NOT NULL),
        '{}'
    ) AS "ColorNames",

    COALESCE(
        ARRAY_AGG(DISTINCT s."SizeId") FILTER (WHERE s."SizeId" IS NOT NULL),
        '{}'
    ) AS "SizeIds",
    COALESCE(
        ARRAY_AGG(DISTINCT s."Size") FILTER (WHERE s."Size" IS NOT NULL),
        '{}'
    ) AS "SizeNames",
    COALESCE(
        ARRAY_AGG(DISTINCT pi."ImagePath") FILTER (WHERE pi."ImagePath" IS NOT NULL),
        '{}'
    ) AS "ImagePaths"

FROM "productMn" p
LEFT JOIN "productMn_colors" pc ON p."ProductMnID" = pc."ProductMnID"
LEFT JOIN "productMn_sizes" ps ON p."ProductMnID" = ps."ProductMnID"
LEFT JOIN "Color" c ON pc."ColorId" = c."ColorId"
LEFT JOIN "Size" s ON ps."SizeId" = s."SizeId"
LEFT JOIN "productImagesEn" pi ON p."ProductMnID" = pi."ProductMnID"
WHERE p."ProductMnID" = sqlc.arg('ProductMnID')
GROUP BY p."ProductMnID";