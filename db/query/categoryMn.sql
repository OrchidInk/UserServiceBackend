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
    sc."subCategoryNameMn"
FROM
    "categoryMn" c
    LEFT JOIN "subCategoryMn" sc ON c."CategoryMnID" = sc."CategoryMnID"
ORDER BY
    c."CategoryMnID";

-- name: GetCategoriesWithSubCategoriesAndProductMn :many
SELECT
    c."CategoryMnID",
    c."CategoryNameMn",
    sc."SubCategoryIDMn",
    sc."subCategoryNameMn",
    p."ProductMnID",
    p."ProductNameMn",
    p."PriceMn",
    p."StockQuantity",
    p."ImagesPathMn"
FROM
    "categoryMn" c
    LEFT JOIN "subCategoryMn" sc ON c."CategoryMnID" = sc."CategoryMnID"
    LEFT JOIN "productMn" p ON sc."SubCategoryIDMn" = p."SubCategoryIDMn"
ORDER BY
    c."CategoryMnID",
    sc."SubCategoryIDMn",
    p."ProductMnID";

-- name: FindSubCategoriesAndProductsByCategoryIDMn :many
SELECT
    c."CategoryMnID",
    c."CategoryNameMn",
    sc."SubCategoryIDMn",
    sc."subCategoryNameMn",
    p."ProductMnID",
    p."ProductNameMn",
    p."PriceMn",
    p."StockQuantity",
    p."ImagesPathMn"
FROM
    "categoryMn" c
    LEFT JOIN "subCategoryMn" sc ON c."CategoryMnID" = sc."CategoryMnID"
    LEFT JOIN "productMn" p ON p."SubCategoryIDMn" = sc."SubCategoryIDMn"
WHERE
    c."CategoryMnID" = sqlc.arg('CategoryMnID')
ORDER BY
    sc."SubCategoryIDMn",
    p."ProductMnID";