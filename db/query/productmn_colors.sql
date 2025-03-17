-- name: InsertProductMnColor :one
INSERT into "productMn_colors" ("ProductMnID", "ColorId") VALUES (
    sqlc.arg('ProductMnID'),
    sqlc.arg('ColorId')
) RETURNING *;