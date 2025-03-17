-- name: CreateOrder :one
INSERT INTO "Orders" (
  "CustomerOrderId",
  "UserId",
  "PhoneNumber",
  "OrderItems",
  "CreatedAt"
) VALUES (
  sqlc.arg('CustomerOrderID'),
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