-- name: InsertProductEnColor :one
INSERT INTO "productEn_colors" ("ProductEnID", "ColorId")
VALUES (
    $1, 
   $2 
) RETURNING *;

