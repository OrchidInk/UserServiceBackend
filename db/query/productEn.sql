-- name: CreateProductEn :one
INSERT INTO
    "productEn" (
        "ProductNameEn",
        "sCategoryIdEn",
        "PriceEn",
        "StockQuantity",
        "ImagesPathEn",
        "DescriptionEn",
        "BrandEn",
        "ManufacturedCountryEn",
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
        sqlc.arg('sCategoryIdEn'),
        sqlc.arg('PriceEn'),
        sqlc.arg('StockQuantity'),
        sqlc.arg('ImagesPathEn'),
        sqlc.arg('DescriptionEn'),
        sqlc.arg('BrandEn'),
        sqlc.arg('ManufacturedCountryEn'),
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
    "sCategoryIdEn" = sqlc.arg('sCategoryIdEn')
WHERE
    "ProductEnID" = sqlc.arg('ProductEnID') RETURNING *;

-- name: UpdateSProductEn :one
UPDATE
    "productEn"
SET
    "StockQuantity" = sqlc.arg('StockQuantity')
WHERE
    "ProductEnID" = sqlc.arg('ProductEnID') RETURNING *;
-- name: UpdateProductEn :one
UPDATE 
    "productEn"
SET
    "ProductNameEn" = sqlc.arg('ProductNameEn'),
    "sCategoryIdEn" = sqlc.arg('sCategoryIdEn'),
    "PriceEn" = sqlc.arg('PriceEn'),
    "StockQuantity" = sqlc.arg('StockQuantity'),
    "ImagesPathEn" = sqlc.arg('ImagesPathEn'),
    "DescriptionEn" = sqlc.arg('DescriptionEn'),
    "BrandEn" = sqlc.arg('BrandEn'),
    "ManufacturedCountryEn" = sqlc.arg('ManufacturedCountryEn'),
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

-- -- name: GetProductEnWithDetails :many
-- SELECT 
--     p."ProductEnID",
--     p."ProductNameEn",
--     p."sCategoryIdEn",
--     p."PriceEn",
--     p."StockQuantity",
--     p."ImagesPathEn",
--     p."DescriptionEn",
--     p."BrandEn",
--     p."ManufacturedCountryEn",
--     p."ColorId",
--     c."Color" AS colorName,
--     p."SizeId",
--     s."Size" AS sizeName,
--     p."PenOutputEn",
--     p."FeaturesEn",
--     p."MaterialEn",
--     p."StapleSizeEn",
--     p."CapacityEn",
--     p."WeightEn",
--     p."ThicknessEn",
--     p."PackagingEn",
--     p."UsageEn",
--     p."InstructionsEn",
--     p."ProductCodeEn",
--     p."CostPriceEn",
--     p."RetailPriceEn",
--     p."WarehouseStockEn",
--     p."Created_At",
--     p."Updated_At"
-- FROM "productEn" p
-- LEFT JOIN "Color" c ON p."ColorId" = c."ColorId"
-- LEFT JOIN "Size" s ON p."SizeId" = s."SizeId";


-- name: GetProductEnWithAllColorsAndSizes :many
SELECT 
    p."ProductEnID",
    p."ProductNameEn",
    p."sCategoryIdEn",
    p."PriceEn",
    p."StockQuantity",
    p."ImagesPathEn",
    p."DescriptionEn",
    p."BrandEn",
    p."ManufacturedCountryEn",
    p."PenOutputEn",
    p."FeaturesEn",
    p."MaterialEn",
    p."StapleSizeEn",
    p."CapacityEn",
    p."WeightEn",
    p."ThicknessEn",
    p."PackagingEn",
    p."UsageEn",
    p."InstructionsEn",
    p."ProductCodeEn",
    p."CostPriceEn",
    p."RetailPriceEn",
    p."WarehouseStockEn",
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

FROM "productEn" p
LEFT JOIN "productEn_colors" pc 
       ON p."ProductEnID" = pc."ProductEnID"
LEFT JOIN "productEn_sizes" ps 
       ON p."ProductEnID" = ps."ProductEnID"
LEFT JOIN "Color" c 
       ON pc."ColorId" = c."ColorId"
LEFT JOIN "Size" s 
       ON ps."SizeId" = s."SizeId"
LEFT JOIN "productImagesEn" pi ON p."ProductEnID" = pi."ProductEnID"
GROUP BY p."ProductEnID";


-- name: GetProductEnWithAllColorsAndSizesByID :one
SELECT 
    p."ProductEnID",
    p."ProductNameEn",
    p."sCategoryIdEn",
    p."PriceEn",
    p."StockQuantity",
    p."ImagesPathEn",
    p."DescriptionEn",
    p."BrandEn",
    p."ManufacturedCountryEn",
    p."PenOutputEn",
    p."FeaturesEn",
    p."MaterialEn",
    p."StapleSizeEn",
    p."CapacityEn",
    p."WeightEn",
    p."ThicknessEn",
    p."PackagingEn",
    p."UsageEn",
    p."InstructionsEn",
    p."ProductCodeEn",
    p."CostPriceEn",
    p."RetailPriceEn",
    p."WarehouseStockEn",
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
FROM "productEn" p
LEFT JOIN "productEn_colors" pc ON p."ProductEnID" = pc."ProductEnID"
LEFT JOIN "productEn_sizes" ps ON p."ProductEnID" = ps."ProductEnID"
LEFT JOIN "Color" c ON pc."ColorId" = c."ColorId"
LEFT JOIN "Size" s ON ps."SizeId" = s."SizeId"
LEFT JOIN "productImagesEn" pi ON p."ProductEnID" = pi."ProductEnID"
WHERE p."ProductEnID" = sqlc.arg('ProductEnID')
GROUP BY p."ProductEnID";
