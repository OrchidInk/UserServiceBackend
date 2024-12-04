-- name: CreateOrderItem :one
INSERT INTO "OrderItems" (
    "CustomerOrderId",
    "ProductMnID",
    "ProductEnID",
    "Quantity",
    "PriceAtOrder"
) VALUES (
    sqlc.arg('CustomerOrderID'),
    sqlc.arg('ProductMnID'),
    sqlc.arg('ProductEnID'),
    sqlc.arg('Quantity'),
    sqlc.arg('PriceAtOrder')
) RETURNING *;

-- name: GetOrderItemsByCustomerOrderID :many
SELECT *
FROM "OrderItems"
WHERE "CustomerOrderId" = sqlc.arg('CustomerOrderID');

-- name: UpdateOrderItem :one
UPDATE "OrderItems"
SET "Quantity" = sqlc.arg('Quantity'),
    "PriceAtOrder" = sqlc.arg('PriceAtOrder')
WHERE "OrderItemId" = sqlc.arg('OrderItemID') RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM "OrderItems"
WHERE "OrderItemId" = sqlc.arg('OrderItemID');
