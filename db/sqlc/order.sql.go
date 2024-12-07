// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: order.sql

package db

import (
	"context"
)

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO "OrderItems" (
    "CustomerOrderId",
    "ProductMnID",
    "ProductEnID",
    "Quantity",
    "PriceAtOrder"
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING "OrderItemId", "CustomerOrderId", "ProductMnID", "ProductEnID", "Quantity", "PriceAtOrder"
`

type CreateOrderItemParams struct {
	CustomerOrderID int32
	ProductMnID     int32
	ProductEnID     int32
	Quantity        int32
	PriceAtOrder    string
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrderItem,
		arg.CustomerOrderID,
		arg.ProductMnID,
		arg.ProductEnID,
		arg.Quantity,
		arg.PriceAtOrder,
	)
	var i OrderItem
	err := row.Scan(
		&i.OrderItemId,
		&i.CustomerOrderId,
		&i.ProductMnID,
		&i.ProductEnID,
		&i.Quantity,
		&i.PriceAtOrder,
	)
	return i, err
}

const deleteOrderItem = `-- name: DeleteOrderItem :exec
DELETE FROM "OrderItems"
WHERE "OrderItemId" = $1
`

func (q *Queries) DeleteOrderItem(ctx context.Context, orderitemid int32) error {
	_, err := q.db.ExecContext(ctx, deleteOrderItem, orderitemid)
	return err
}

const getOrderItemsByCustomerOrderID = `-- name: GetOrderItemsByCustomerOrderID :many
SELECT "OrderItemId", "CustomerOrderId", "ProductMnID", "ProductEnID", "Quantity", "PriceAtOrder"
FROM "OrderItems"
WHERE "CustomerOrderId" = $1
`

func (q *Queries) GetOrderItemsByCustomerOrderID(ctx context.Context, customerorderid int32) ([]OrderItem, error) {
	rows, err := q.db.QueryContext(ctx, getOrderItemsByCustomerOrderID, customerorderid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrderItem
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.OrderItemId,
			&i.CustomerOrderId,
			&i.ProductMnID,
			&i.ProductEnID,
			&i.Quantity,
			&i.PriceAtOrder,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrderItem = `-- name: UpdateOrderItem :one
UPDATE "OrderItems"
SET "Quantity" = $1,
    "PriceAtOrder" = $2
WHERE "OrderItemId" = $3 RETURNING "OrderItemId", "CustomerOrderId", "ProductMnID", "ProductEnID", "Quantity", "PriceAtOrder"
`

type UpdateOrderItemParams struct {
	Quantity     int32
	PriceAtOrder string
	OrderItemID  int32
}

func (q *Queries) UpdateOrderItem(ctx context.Context, arg UpdateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, updateOrderItem, arg.Quantity, arg.PriceAtOrder, arg.OrderItemID)
	var i OrderItem
	err := row.Scan(
		&i.OrderItemId,
		&i.CustomerOrderId,
		&i.ProductMnID,
		&i.ProductEnID,
		&i.Quantity,
		&i.PriceAtOrder,
	)
	return i, err
}
