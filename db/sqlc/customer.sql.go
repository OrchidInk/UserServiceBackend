// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: customer.sql

package db

import (
	"context"
)

const countActiveCustomers = `-- name: CountActiveCustomers :one
SELECT
    COUNT(*) AS active_count
FROM
    "Customer"
WHERE
    "IsActive" = TRUE
`

func (q *Queries) CountActiveCustomers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countActiveCustomers)
	var active_count int64
	err := row.Scan(&active_count)
	return active_count, err
}

const countInactiveCustomers = `-- name: CountInactiveCustomers :one
SELECT
    COUNT(*) AS inactive_count
FROM
    "Customer"
WHERE
    "IsActive" = FALSE
`

func (q *Queries) CountInactiveCustomers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countInactiveCustomers)
	var inactive_count int64
	err := row.Scan(&inactive_count)
	return inactive_count, err
}

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO
    "Customer" (
        "CustomerName",
        "ContractStartDate",
        "ContractEndDate",
        "IsActive"
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4
    ) RETURNING "CustomerId", "CustomerName", "ContractStartDate", "ContractEndDate", "IsActive", "Created_At", "Updated_At"
`

type CreateCustomerParams struct {
	CustomerName      string
	ContractStartDate int32
	ContractEndDate   int32
	IsActive          bool
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.CustomerName,
		arg.ContractStartDate,
		arg.ContractEndDate,
		arg.IsActive,
	)
	var i Customer
	err := row.Scan(
		&i.CustomerId,
		&i.CustomerName,
		&i.ContractStartDate,
		&i.ContractEndDate,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCustomerById = `-- name: DeleteCustomerById :exec
DELETE FROM
    "Customer"
WHERE
    "CustomerId" = $1
`

func (q *Queries) DeleteCustomerById(ctx context.Context, customerid int32) error {
	_, err := q.db.ExecContext(ctx, deleteCustomerById, customerid)
	return err
}

const findByCustomerId = `-- name: FindByCustomerId :one
SELECT
    "CustomerId", "CustomerName", "ContractStartDate", "ContractEndDate", "IsActive", "Created_At", "Updated_At"
FROM
    "Customer"
WHERE
    "CustomerId" = $1
LIMIT
    1
`

func (q *Queries) FindByCustomerId(ctx context.Context, customerid int32) (Customer, error) {
	row := q.db.QueryRowContext(ctx, findByCustomerId, customerid)
	var i Customer
	err := row.Scan(
		&i.CustomerId,
		&i.CustomerName,
		&i.ContractStartDate,
		&i.ContractEndDate,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findByCustomerIdAndIsActiveFalse = `-- name: FindByCustomerIdAndIsActiveFalse :one
SELECT
    "CustomerId", "CustomerName", "ContractStartDate", "ContractEndDate", "IsActive", "Created_At", "Updated_At"
FROM
    "Customer"
WHERE
    "CustomerId" = $1
    AND "IsActive" = FALSE
LIMIT
    1
`

func (q *Queries) FindByCustomerIdAndIsActiveFalse(ctx context.Context, customerid int32) (Customer, error) {
	row := q.db.QueryRowContext(ctx, findByCustomerIdAndIsActiveFalse, customerid)
	var i Customer
	err := row.Scan(
		&i.CustomerId,
		&i.CustomerName,
		&i.ContractStartDate,
		&i.ContractEndDate,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findByCustomerIdAndIsActiveTrue = `-- name: FindByCustomerIdAndIsActiveTrue :one
SELECT
    "CustomerId", "CustomerName", "ContractStartDate", "ContractEndDate", "IsActive", "Created_At", "Updated_At"
FROM
    "Customer"
WHERE
    "CustomerId" = $1
    AND "IsActive" = TRUE
LIMIT
    1
`

func (q *Queries) FindByCustomerIdAndIsActiveTrue(ctx context.Context, customerid int32) (Customer, error) {
	row := q.db.QueryRowContext(ctx, findByCustomerIdAndIsActiveTrue, customerid)
	var i Customer
	err := row.Scan(
		&i.CustomerId,
		&i.CustomerName,
		&i.ContractStartDate,
		&i.ContractEndDate,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllCustomers = `-- name: GetAllCustomers :many
SELECT
    "CustomerId", "CustomerName", "ContractStartDate", "ContractEndDate", "IsActive", "Created_At", "Updated_At"
FROM
    "Customer"
`

func (q *Queries) GetAllCustomers(ctx context.Context) ([]Customer, error) {
	rows, err := q.db.QueryContext(ctx, getAllCustomers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.CustomerId,
			&i.CustomerName,
			&i.ContractStartDate,
			&i.ContractEndDate,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateCustomerContractDates = `-- name: UpdateCustomerContractDates :one
UPDATE
    "Customer"
SET
    "ContractStartDate" = $1,
    "ContractEndDate" = $2
WHERE
    "CustomerId" = $3 RETURNING "CustomerId", "CustomerName", "ContractStartDate", "ContractEndDate", "IsActive", "Created_At", "Updated_At"
`

type UpdateCustomerContractDatesParams struct {
	ContractStartDate int32
	ContractEndDate   int32
	CustomerId        int32
}

func (q *Queries) UpdateCustomerContractDates(ctx context.Context, arg UpdateCustomerContractDatesParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, updateCustomerContractDates, arg.ContractStartDate, arg.ContractEndDate, arg.CustomerId)
	var i Customer
	err := row.Scan(
		&i.CustomerId,
		&i.CustomerName,
		&i.ContractStartDate,
		&i.ContractEndDate,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCustomerIsActive = `-- name: UpdateCustomerIsActive :one
UPDATE
    "Customer"
SET
    "IsActive" = $1
WHERE
    "CustomerId" = $2 RETURNING "CustomerId", "CustomerName", "ContractStartDate", "ContractEndDate", "IsActive", "Created_At", "Updated_At"
`

type UpdateCustomerIsActiveParams struct {
	IsActive   bool
	CustomerId int32
}

func (q *Queries) UpdateCustomerIsActive(ctx context.Context, arg UpdateCustomerIsActiveParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, updateCustomerIsActive, arg.IsActive, arg.CustomerId)
	var i Customer
	err := row.Scan(
		&i.CustomerId,
		&i.CustomerName,
		&i.ContractStartDate,
		&i.ContractEndDate,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}