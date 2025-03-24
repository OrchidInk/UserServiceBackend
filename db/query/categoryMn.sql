-- name: CreateCategoryMn :one
INSERT INTO
    "categoryMn" ("CategoryNameMn")
VALUES
    (
        sqlc.arg('CategoryNameMn') :: VARCHAR(100)
    ) RETURNING *;

-- name: UpdateCategoryMn :one
UPDATE
    "categoryMn"
SET
    "CategoryNameMn" = sqlc.arg('CategoryNameMn') :: VARCHAR(100)
WHERE
    "CategoryMnID" = sqlc.arg('CategoryMnID') :: INT RETURNING *;

-- name: GetListByAllCategoryMn :many
SELECT
    *
FROM
    "categoryMn";

-- name: DeleteCategoryByMnId :exec
DELETE FROM
    "categoryMn"
WHERE
    "CategoryMnID" = sqlc.arg('CategoryMnID') :: INT;

-- name: FindByCategoryMnId :one
SELECT
    *
FROM
    "categoryMn"
WHERE
    "CategoryMnID" = sqlc.arg('CategoryMnID')
LIMIT
    1;

-- name: FindByNameMnCategoryMn :one
SELECT
    *
FROM
    "categoryMn"
WHERE
    "CategoryNameMn" = sqlc.arg('CategoryNameMn') :: VARCHAR(100)
LIMIT
    1;

-- name: FindByCategoryMn :one 
SELECT
    *
FROM
    "categoryMn"
WHERE
    "CategoryMnID" = sqlc.arg('CategoryMnID')
LIMIT
    1;

-- name: GetCategoriesWithSubCategoriesMn :many
SELECT
    c."CategoryMnID",
    c."CategoryNameMn",
    sc."SubCategoryIDMn",
    sc."subCategoryNameMn",
    scc."sCategoryIdMn",
    scc."sCategoryNameMn"
FROM
    "categoryMn" c
    LEFT JOIN "subCategoryMn" sc 
        ON c."CategoryMnID" = sc."CategoryMnID"
    LEFT JOIN "sCategoryMn" scc 
        ON scc."SubCategoryIDMn" = sc."SubCategoryIDMn"
ORDER BY
    c."CategoryMnID";


-- name: FindSubCategoriesAndProductsByCategoryIDMn :many
SELECT
    c."CategoryMnID",
    c."CategoryNameMn",
    sc."SubCategoryIDMn",
    sc."subCategoryNameMn",
    scc."sCategoryIdMn",
    scc."sCategoryNameMn",
    p."ProductMnID",
    p."ProductNameMn",
    p."PriceMn",
    p."StockQuantity",
    COALESCE(
        MIN(pi."ImagePath"),
        p."ImagesPathMn"
    ) AS "ImagesPathMn"
FROM "categoryMn" c
LEFT JOIN "subCategoryMn" sc ON c."CategoryMnID" = sc."CategoryMnID"
LEFT JOIN "sCategoryMn" scc ON scc."SubCategoryIDMn" = sc."SubCategoryIDMn"
LEFT JOIN "productMn" p ON p."sCategoryIdMn" = scc."sCategoryIdMn"
LEFT JOIN "productImagesMn" pi ON p."ProductMnID" = pi."ProductMnID"
WHERE c."CategoryMnID" = sqlc.arg('CategoryMnID')
GROUP BY
    c."CategoryMnID",
    c."CategoryNameMn",
    sc."SubCategoryIDMn",
    sc."subCategoryNameMn",
    scc."sCategoryIdMn",
    scc."sCategoryNameMn",
    p."ProductMnID",
    p."ProductNameMn",
    p."PriceMn",
    p."StockQuantity",
    p."ImagesPathMn"
ORDER BY
    scc."sCategoryIdMn",
    p."ProductMnID";

-- name: GetProductMnWithAllColorsAndSizesByIDMn :one
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
LEFT JOIN "productImagesMn" pi ON p."ProductMnID" = pi."ProductMnID"
WHERE p."ProductMnID" = sqlc.arg('ProductMnID')
GROUP BY p."ProductMnID";
-- name: GetCategoriesWithSubCategoriesAndProductMn :many
SELECT 
    c."CategoryMnID",
    c."CategoryNameMn",
    sc."SubCategoryIDMn",
    sc."subCategoryNameMn",
    scc."sCategoryIdMn",
    scc."sCategoryNameMn",
    p."ProductMnID",
    p."ProductNameMn",
    p."PriceMn",
    p."StockQuantity",
    COALESCE(
        MIN(pi."ImagePath"),
        p."ImagesPathMn"
    ) AS "ImagesPathMn"
FROM "categoryMn" c
LEFT JOIN "subCategoryMn" sc ON c."CategoryMnID" = sc."CategoryMnID"
LEFT JOIN "sCategoryMn" scc ON scc."SubCategoryIDMn" = sc."SubCategoryIDMn"
LEFT JOIN "productMn" p ON scc."sCategoryIdMn" = p."sCategoryIdMn"
LEFT JOIN "productImagesMn" pi ON p."ProductMnID" = pi."ProductMnID"
GROUP BY 
    c."CategoryMnID",
    c."CategoryNameMn",
    sc."SubCategoryIDMn",
    sc."subCategoryNameMn",
    scc."sCategoryIdMn",
    scc."sCategoryNameMn",
    p."ProductMnID",
    p."ProductNameMn",
    p."PriceMn",
    p."StockQuantity",
    p."ImagesPathMn"
ORDER BY 
    c."CategoryMnID",
    scc."sCategoryIdMn",
    p."ProductMnID";

