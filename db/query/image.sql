-- name: CreateProductImageEn :one
INSERT INTO "productImagesEn" (
    "ProductEnID",
    "ImagePath"
)
VALUES (
    sqlc.arg('ProductEnID'),
    sqlc.arg('ImagePath')
)
RETURNING *;

-- name: GetProductImagesEnByProductID :many
SELECT *
FROM "productImagesEn"
WHERE "ProductEnID" = sqlc.arg('ProductEnID')
ORDER BY "Created_At" ASC;

-- name: UpdateProductImageEn :one
UPDATE "productImagesEn"
SET "ImagePath" = sqlc.arg('ImagePath')
WHERE "ImageID" = sqlc.arg('ImageID')
RETURNING *;

-- name: DeleteProductImageEn :exec
DELETE FROM "productImagesEn"
WHERE "ImageID" = sqlc.arg('ImageID');

-- name: CreateProductImageMn :one
INSERT INTO "productImagesMn" (
    "ProductMnID",
    "ImagePath"
)
VALUES (
    sqlc.arg('ProductMnID'),
    sqlc.arg('ImagePath')
)
RETURNING *;

-- name: GetProductImagesMnByProductID :many
SELECT *
FROM "productImagesMn"
WHERE "ProductMnID" = sqlc.arg('ProductMnID')
ORDER BY "Created_At" ASC;

-- name: UpdateProductImageMn :one
UPDATE "productImagesMn"
SET "ImagePath" = sqlc.arg('ImagePath')
WHERE "ImageID" = sqlc.arg('ImageID')
RETURNING *;

-- name: DeleteProductImageMn :exec
DELETE FROM "productImagesMn"
WHERE "ImageID" = sqlc.arg('ImageID');
