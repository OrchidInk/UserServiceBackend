-- name: InsertProductMnSize :one
INSERT into "productMn_sizes" ("ProductMnID", "SizeId") VALUES (
    sqlc.arg('ProductMnID'),
    sqlc.arg('SizeId')
) RETURNING *;