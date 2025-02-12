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
    p."ImagesPathMn"
FROM
    "categoryMn" c
    LEFT JOIN "subCategoryMn" sc ON c."CategoryMnID" = sc."CategoryMnID"
    LEFT JOIN "sCategoryMn" scc ON scc."SubCategoryIDMn" = sc."SubCategoryIDMn"
    LEFT JOIN "productMn" p ON scc."sCategoryIdMn" = p."sCategoryIdMn"
ORDER BY
    c."CategoryMnID",
    scc."sCategoryIdMn",
    p."ProductMnID";

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
    p."ImagesPathMn"
FROM
    "categoryMn" c
    LEFT JOIN "subCategoryMn" sc ON c."CategoryMnID" = sc."CategoryMnID"
    LEFT JOIN "sCategoryMn" scc ON scc."SubCategoryIDMn" = sc."SubCategoryIDMn"
    LEFT JOIN "productMn" p ON p."sCategoryIdMn" = scc."sCategoryIdMn"
WHERE
    c."CategoryMnID" = sqlc.arg('CategoryMnID')
ORDER BY
    scc."sCategoryIdMn",
    p."ProductMnID";
