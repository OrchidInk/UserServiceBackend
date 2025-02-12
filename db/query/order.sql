-- name: CreateOrderItem :one
INSERT INTO "OrderItems" (
    "CustomerOrderId",
    "ProductMnID",
    "ProductEnID",
    "UserId",
    "PhoneNumber",
    "Quantity",
    "PriceAtOrder"
) VALUES (
    sqlc.arg('CustomerOrderID'),
    sqlc.arg('ProductMnID'),
    sqlc.arg('ProductEnID'),
    sqlc.arg('UserId'),
    sqlc.arg('PhoneNumber'),
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

-- name: CountTotalCustomers :one
SELECT 
    COUNT(*) AS total_customers
FROM 
    "Customer";

-- name: CountActiveCustomersOrder :one
SELECT 
    COUNT(*) AS active_customers
FROM 
    "Customer"
WHERE 
    "IsActive" = TRUE;

-- name: CountInactiveCustomersOrder :one
SELECT 
    COUNT(*) AS inactive_customers
FROM 
    "Customer"
WHERE 
    "IsActive" = FALSE;

-- name: GetCustomerAnalysisOrder :one
SELECT 
    (SELECT COUNT(*) FROM "Customer") AS total_customers,
    (SELECT COUNT(*) FROM "Customer" WHERE "IsActive" = TRUE) AS active_customers,
    (SELECT COUNT(*) FROM "Customer" WHERE "IsActive" = FALSE) AS inactive_customers;

-- name: GetDeliveriesByUserId :many
SELECT
    d.*
FROM
    "delivery" d
    JOIN "CustomerOrderDetail" cod ON d."OrderId" = cod."CustomerOrderId"
    JOIN "Customer" c ON cod."CustomerId" = c."CustomerId"
WHERE
    c."CustomerId" = sqlc.arg('CustomerId');

-- name: FindByOrderItemsId :one
SELECT
    *
FROM
    "OrderItems"
WHERE
    "OrderItemId" = sqlc.arg('OrderItemId');