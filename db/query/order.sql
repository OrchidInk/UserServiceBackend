-- name: CreateOrder :one
INSERT INTO "Orders" (
  "CustomerOrderId",
  "CompName",
  "UserName",
  "UserId",
  "PhoneNumber",
  "OrderItems",
  "CreatedAt"
) VALUES (
  sqlc.arg('CustomerOrderID'),
  sqlc.arg('CompName'),
  sqlc.arg('UserName'),
  sqlc.arg('UserId'),
  sqlc.arg('PhoneNumber'),
  sqlc.arg('OrderItems'),
  sqlc.arg('CreatedAt')
)
RETURNING *;


-- name: CreateOrderItem :one
INSERT INTO "OrderItems" (
  "OrderID",
  "ProductMnID",
  "ProductEnID",
  "Quantity",
  "PriceAtOrder",
  "SelectedColor",
  "SelectedSize"
) VALUES (
  sqlc.arg('OrderID'),
  sqlc.arg('ProductMnID'),
  sqlc.arg('ProductEnID'),
  sqlc.arg('Quantity'),
  sqlc.arg('PriceAtOrder'),
  sqlc.arg('SelectedColor'),
  sqlc.arg('SelectedSize')
)
RETURNING *;

-- name: GetOrdersWithDetails :many
SELECT
    *
from
    "Orders";

-- name: OrderDelete :exec
DELETE FROM 
  "Orders"
WHERE
  "OrderID" = sqlc.arg('OrderID');