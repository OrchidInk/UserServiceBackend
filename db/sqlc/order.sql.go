// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: order.sql

package db

import (
	"context"
	"database/sql"
)

const countActiveCustomersOrder = `-- name: CountActiveCustomersOrder :one
SELECT 
    COUNT(*) AS active_customers
FROM 
    "Customer"
WHERE 
    "IsActive" = TRUE
`

func (q *Queries) CountActiveCustomersOrder(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countActiveCustomersOrder)
	var active_customers int64
	err := row.Scan(&active_customers)
	return active_customers, err
}

const countInactiveCustomersOrder = `-- name: CountInactiveCustomersOrder :one
SELECT 
    COUNT(*) AS inactive_customers
FROM 
    "Customer"
WHERE 
    "IsActive" = FALSE
`

func (q *Queries) CountInactiveCustomersOrder(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countInactiveCustomersOrder)
	var inactive_customers int64
	err := row.Scan(&inactive_customers)
	return inactive_customers, err
}

const countTotalCustomers = `-- name: CountTotalCustomers :one
SELECT 
    COUNT(*) AS total_customers
FROM 
    "Customer"
`

func (q *Queries) CountTotalCustomers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTotalCustomers)
	var total_customers int64
	err := row.Scan(&total_customers)
	return total_customers, err
}

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO "OrderItems" (
    "CustomerOrderId",
    "ProductMnID",
    "ProductEnID",
    "UserId",
    "PhoneNumber",
    "Quantity",
    "PriceAtOrder",
    "CreatedAt"
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) RETURNING "OrderItemId", "CustomerOrderId", "ProductMnID", "ProductEnID", "UserId", "PhoneNumber", "Quantity", "PriceAtOrder", "CreatedAt"
`

type CreateOrderItemParams struct {
	CustomerOrderID sql.NullInt32
	ProductMnID     sql.NullInt32
	ProductEnID     sql.NullInt32
	UserId          int32
	PhoneNumber     string
	Quantity        int32
	PriceAtOrder    string
	CreatedAt       sql.NullTime
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrderItem,
		arg.CustomerOrderID,
		arg.ProductMnID,
		arg.ProductEnID,
		arg.UserId,
		arg.PhoneNumber,
		arg.Quantity,
		arg.PriceAtOrder,
		arg.CreatedAt,
	)
	var i OrderItem
	err := row.Scan(
		&i.OrderItemId,
		&i.CustomerOrderId,
		&i.ProductMnID,
		&i.ProductEnID,
		&i.UserId,
		&i.PhoneNumber,
		&i.Quantity,
		&i.PriceAtOrder,
		&i.CreatedAt,
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

const findByOrderItemsId = `-- name: FindByOrderItemsId :one
SELECT
    "OrderItemId", "CustomerOrderId", "ProductMnID", "ProductEnID", "UserId", "PhoneNumber", "Quantity", "PriceAtOrder", "CreatedAt"
FROM
    "OrderItems"
WHERE
    "OrderItemId" = $1
`

func (q *Queries) FindByOrderItemsId(ctx context.Context, orderitemid int32) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, findByOrderItemsId, orderitemid)
	var i OrderItem
	err := row.Scan(
		&i.OrderItemId,
		&i.CustomerOrderId,
		&i.ProductMnID,
		&i.ProductEnID,
		&i.UserId,
		&i.PhoneNumber,
		&i.Quantity,
		&i.PriceAtOrder,
		&i.CreatedAt,
	)
	return i, err
}

const getCustomerAnalysisOrder = `-- name: GetCustomerAnalysisOrder :one
SELECT 
    (SELECT COUNT(*) FROM "Customer") AS total_customers,
    (SELECT COUNT(*) FROM "Customer" WHERE "IsActive" = TRUE) AS active_customers,
    (SELECT COUNT(*) FROM "Customer" WHERE "IsActive" = FALSE) AS inactive_customers
`

type GetCustomerAnalysisOrderRow struct {
	TotalCustomers    int64
	ActiveCustomers   int64
	InactiveCustomers int64
}

func (q *Queries) GetCustomerAnalysisOrder(ctx context.Context) (GetCustomerAnalysisOrderRow, error) {
	row := q.db.QueryRowContext(ctx, getCustomerAnalysisOrder)
	var i GetCustomerAnalysisOrderRow
	err := row.Scan(&i.TotalCustomers, &i.ActiveCustomers, &i.InactiveCustomers)
	return i, err
}

const getDeliveriesByUserId = `-- name: GetDeliveriesByUserId :many
SELECT
    d."DeliverId", d."DeliverName", d."OrderId", d."DeliveryAmount", d."CreatedAt"
FROM
    "delivery" d
    JOIN "CustomerOrderDetail" cod ON d."OrderId" = cod."CustomerOrderId"
    JOIN "Customer" c ON cod."CustomerId" = c."CustomerId"
WHERE
    c."CustomerId" = $1
`

func (q *Queries) GetDeliveriesByUserId(ctx context.Context, customerid int32) ([]Delivery, error) {
	rows, err := q.db.QueryContext(ctx, getDeliveriesByUserId, customerid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Delivery
	for rows.Next() {
		var i Delivery
		if err := rows.Scan(
			&i.DeliverId,
			&i.DeliverName,
			&i.OrderId,
			&i.DeliveryAmount,
			&i.CreatedAt,
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

const getListAll = `-- name: GetListAll :many
SELECT
    "OrderItemId", "CustomerOrderId", "ProductMnID", "ProductEnID", "UserId", "PhoneNumber", "Quantity", "PriceAtOrder", "CreatedAt"
FROM
    "OrderItems"
`

func (q *Queries) GetListAll(ctx context.Context) ([]OrderItem, error) {
	rows, err := q.db.QueryContext(ctx, getListAll)
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
			&i.UserId,
			&i.PhoneNumber,
			&i.Quantity,
			&i.PriceAtOrder,
			&i.CreatedAt,
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

const getOrderItemsByCustomerOrderID = `-- name: GetOrderItemsByCustomerOrderID :many
SELECT "OrderItemId", "CustomerOrderId", "ProductMnID", "ProductEnID", "UserId", "PhoneNumber", "Quantity", "PriceAtOrder", "CreatedAt"
FROM "OrderItems"
WHERE "CustomerOrderId" = $1
`

func (q *Queries) GetOrderItemsByCustomerOrderID(ctx context.Context, customerorderid sql.NullInt32) ([]OrderItem, error) {
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
			&i.UserId,
			&i.PhoneNumber,
			&i.Quantity,
			&i.PriceAtOrder,
			&i.CreatedAt,
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
WHERE "OrderItemId" = $3 RETURNING "OrderItemId", "CustomerOrderId", "ProductMnID", "ProductEnID", "UserId", "PhoneNumber", "Quantity", "PriceAtOrder", "CreatedAt"
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
		&i.UserId,
		&i.PhoneNumber,
		&i.Quantity,
		&i.PriceAtOrder,
		&i.CreatedAt,
	)
	return i, err
}
