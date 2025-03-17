-- name: CreateSize :one
insert into "Size" ("Size") values (sqlc.arg('Size')) RETURNING *;

-- name: GetAllSize :many
SELECT
    *
FROM
    "Size";

-- name: UpdateSize :one
UPDATE
    "Size"
SET
    "Size" = sqlc.arg('Size')
WHERE   
    "SizeId" = sqlc.arg('SizeId') RETURNING *;

-- name: DeleteSize :exec
DELETE FROM
    "Size"
WHERE
    "SizeId" = sqlc.arg('SizeId');


-- name: FindByIdSize :one
SELECT
    *
FROM
    "Size"
WHERE
    "SizeId" = sqlc.arg('SizeId');