-- name: CreateSubCategoryEn :one
INSERT INTO
    "subCategoryEn" (
        "subCategoryNameEn",
        -- Correct casing here
        "CategoryEnID"
    )
VALUES
    (
        sqlc.arg('subCategoryNameEn') :: VARCHAR(100),
        sqlc.arg('CategoryEnID') :: INT
    ) RETURNING *;

-- name: GetListAllSubCategoriesEn :many
SELECT
    *
FROM
    "subCategoryEn";

-- name: UpdateSubCategoryNameEn :one
UPDATE
    "subCategoryEn"
SET
    "subCategoryNameEn" = sqlc.arg('subCategoryNameEn') :: VARCHAR(100)
WHERE
    "SubCategoryIDEn" = sqlc.arg('SubCategoryIDEn') RETURNING *;

-- name: DeleteSubCategoryEn :exec
DELETE FROM
    "subCategoryEn"
WHERE
    "SubCategoryIDEn" = sqlc.arg('SubCategoryIDEn');

-- name: FindBySubCategoryIDEn :one
SELECT
    *
FROM
    "subCategoryEn"
WHERE
    "SubCategoryIDEn" = sqlc.arg('SubCategoryIDEn')
LIMIT
    1;

-- name: FindByNameSubCategoryEn :one
SELECT
    *
FROM
    "subCategoryEn"
WHERE
    "subCategoryNameEn" = sqlc.arg('subCategoryNameEn')
LIMIT
    1;

-- name: GetProductsBySubCategoryEn :many
SELECT
    scc."sCategoryIdEn",
    scc."sCategoryNameEn",
    p."ProductEnID",
    p."ProductNameEn",
    p."PriceEn",
    p."StockQuantity",
    p."ImagesPathEn"
FROM
    "subCategoryEn" sc
    JOIN "sCategoryEn" scc ON scc."SubCategoryIDEn" = sc."SubCategoryIDEn"
    JOIN "productEn" p ON scc."sCategoryIdEn" = p."sCategoryIdEn"
WHERE
    scc."sCategoryIdEn" = sqlc.arg('sCategoryIdEn');

-- name: UpdateSubCategoryByIDEn :one
UPDATE
    "subCategoryEn"
SET
    "subCategoryNameEn" = sqlc.arg('subCategoryNameEn'),
    "CategoryEnID" = sqlc.arg('CategoryEnID')
WHERE
    "SubCategoryIDEn" = sqlc.arg('SubCategoryIDEn') RETURNING *;