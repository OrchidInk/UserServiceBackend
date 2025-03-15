-- name: CreateColor :one
insert into "Color" (
    "Color"
) values (
    sqlc.arg('Color')
) RETURNING *;

-- name: GetAllColor :many
SELECT
    *
FROM
    "Color";

-- name: UpdateColor :one
update 
    "Color"
SET
    "Color" = sqlc.arg('Color')
WHERE
    "ColorId" = sqlc.arg('ColorId') RETURNING *;

-- name: DeleteColor :exec
DELETE FROM
    "Color"
WHERE
    "ColorId" = sqlc.arg('ColorId');

-- name: FindByColorId :one
SELECT
    *
FROM
    "Color"
WHERE
    "ColorId" = sqlc.arg('ColorID')
limit 1;