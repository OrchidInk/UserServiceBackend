-- name: InsertProductEnSize :one
INSERT INTO "productEn_sizes" ("ProductEnID", "SizeId")
VALUES (
    sqlc.arg('ProductEnID'), 
    sqlc.arg('SizeId')) RETURNING *;