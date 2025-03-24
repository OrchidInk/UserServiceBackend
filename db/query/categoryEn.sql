-- name: CreateCategoryEn :one
INSERT INTO
    "categoryEn" ("CategoryNameEn")
VALUES
    (sqlc.arg ('CategoryNameEn')) RETURNING *;

-- name: GetListByAllCategoryEn :many
SELECT
    *
FROM
    "categoryEn";

-- name: UpdateCategoryEn :one
UPDATE
    "categoryEn"
SET
    "CategoryNameEn" = sqlc.arg('CategoryNameEn') :: VARCHAR(100)
WHERE
    "CategoryEnID" = sqlc.arg('CategoryEnID') :: INT RETURNING *;

-- name: DeleteCategoryById :exec
DELETE FROM
    "categoryEn"
WHERE
    "CategoryEnID" = sqlc.arg('CategoryEnID') :: INT;

-- name: FindByCategoryEnId :one
SELECT
    *
FROM
    "categoryEn"
WHERE
    "CategoryEnID" = sqlc.arg ('CategoryEnID')
LIMIT
    1;

-- name: FindByNameCategoryEn :one
SELECT
    *
FROM
    "categoryEn"
WHERE
    "CategoryNameEn" = sqlc.arg ('CategoryNameEn')
LIMIT
    1;

-- name: FindByIdCategoryEn :one
SELECT
    *
FROM
    "categoryEn"
WHERE
    "CategoryEnID" = sqlc.arg ('CategoryEnID')
LIMIT
    1;

-- name: GetCategoriesWithSubCategories :many
SELECT
    c."CategoryEnID",
    c."CategoryNameEn",
    sc."SubCategoryIDEn",
    sc."subCategoryNameEn",
    scc."sCategoryIdEn",
    scc."sCategoryNameEn"
FROM
    "categoryEn" c
    LEFT JOIN "subCategoryEn" sc ON c."CategoryEnID" = sc."CategoryEnID"
    LEFT JOIN "sCategoryEn" scc ON scc."SubCategoryIDEn" = sc."SubCategoryIDEn"
ORDER BY
    c."CategoryEnID";

-- name: GetCategoriesWithSubCategoriesAndProductsEn :many
SELECT
    c."CategoryEnID",
    c."CategoryNameEn",
    sc."SubCategoryIDEn",
    sc."subCategoryNameEn",
    scc."sCategoryIdEn",
    scc."sCategoryNameEn",
    p."ProductEnID",
    p."ProductNameEn",
    p."PriceEn",
    p."StockQuantity",
    COALESCE(pi."ImagePath", p."ImagesPathEn") as "ImagesPathEn"
FROM
    "categoryEn" c
    LEFT JOIN "subCategoryEn" sc ON c."CategoryEnID" = sc."CategoryEnID"
    LEFT JOIN "sCategoryEn" scc ON scc."SubCategoryIDEn" = sc."SubCategoryIDEn"
    LEFT JOIN "productEn" p ON scc."sCategoryIdEn" = p."sCategoryIdEn"
    LEFT JOIN "productImagesEn" pi ON p."ProductEnID" = pi."ProductEnID"
ORDER BY
    c."CategoryEnID",
    scc."sCategoryIdEn",
    p."ProductEnID";

-- name: FindSubCategoriesAndProductsByCategoryIDEn :many
SELECT
    c."CategoryEnID",
    c."CategoryNameEn",
    sc."SubCategoryIDEn",
    sc."subCategoryNameEn",
    scc."sCategoryIdEn",
    scc."sCategoryNameEn",
    p."ProductEnID",
    p."ProductNameEn",
    p."PriceEn",
    p."StockQuantity",
    COALESCE(pi."ImagePath", p."ImagesPathEn") as "ImagesPathEn"
FROM
    "categoryEn" c
    LEFT JOIN "subCategoryEn" sc ON c."CategoryEnID" = sc."CategoryEnID"
    LEFT JOIN "sCategoryEn" scc ON scc."SubCategoryIDEn" = sc."SubCategoryIDEn"
    LEFT JOIN "productEn" p ON scc."sCategoryIdEn" = p."sCategoryIdEn"
    LEFT JOIN "productImagesEn" pi ON p."ProductEnID" = pi."ProductEnID"
WHERE
    c."CategoryEnID" = sqlc.arg('CategoryEnID')
ORDER BY
    scc."sCategoryIdEn",
    p."ProductEnID";
