-- name: CreateDelivery :one
INSERT INTO
    "delivery" (
        "DeliverName",
        "OrderId",
        "DeliveryAmount"
    )
VALUES
    (
        sqlc.arg('DeliverName'),
        sqlc.arg('OrderId'),
        sqlc.arg('DeliveryAmount')
    ) RETURNING *;

-- name: DeleteDelivery :exec
DELETE FROM
    "delivery"
WHERE
    "DeliverId" = sqlc.arg('DeliverId');

-- name: UpdateDeliver :one
UPDATE
    "delivery"
SET
    "DeliverName" = sqlc.arg('DeliverName')
WHERE
    "DeliverId" = sqlc.arg('DeliverId') RETURNING *;

-- name: GetListDeliver :many
SELECT
    *
FROM
    "delivery";

-- name: FindByDeliveryId :one
SELECT
    *
FROM
    "delivery"
WHERE
    "DeliverId" = sqlc.arg('DeliverId')
LIMIT
    1;