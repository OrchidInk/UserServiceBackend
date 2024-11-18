-- name: CreateCustomer :one
INSERT INTO
    "Customer" (
        "CustomerName",
        "ContractStartDate",
        "ContractEndDate",
        "IsActive"
    )
VALUES
    (
        sqlc.arg('CustomerName'),
        sqlc.arg('ContractStartDate'),
        sqlc.arg('ContractEndDate'),
        sqlc.arg('IsActive')
    ) RETURNING *;

-- name: FindByCustomerId :one
SELECT
    *
FROM
    "Customer"
WHERE
    "CustomerId" = sqlc.arg('CustomerId')
LIMIT
    1;

-- name: FindByCustomerIdAndIsActiveTrue :one
SELECT
    *
FROM
    "Customer"
WHERE
    "CustomerId" = sqlc.arg('CustomerId')
    AND "IsActive" = TRUE
LIMIT
    1;

-- name: FindByCustomerIdAndIsActiveFalse :one
SELECT
    *
FROM
    "Customer"
WHERE
    "CustomerId" = sqlc.arg('CustomerId')
    AND "IsActive" = FALSE
LIMIT
    1;

-- name: GetAllCustomers :many
SELECT
    *
FROM
    "Customer";

-- name: CountActiveCustomers :one
SELECT
    COUNT(*) AS active_count
FROM
    "Customer"
WHERE
    "IsActive" = TRUE;

-- name: CountInactiveCustomers :one
SELECT
    COUNT(*) AS inactive_count
FROM
    "Customer"
WHERE
    "IsActive" = FALSE;

-- name: UpdateCustomerIsActive :one
UPDATE
    "Customer"
SET
    "IsActive" = sqlc.arg('IsActive')
WHERE
    "CustomerId" = sqlc.arg('CustomerId') RETURNING *;

-- name: UpdateCustomerContractDates :one
UPDATE
    "Customer"
SET
    "ContractStartDate" = sqlc.arg('ContractStartDate'),
    "ContractEndDate" = sqlc.arg('ContractEndDate')
WHERE
    "CustomerId" = sqlc.arg('CustomerId') RETURNING *;

-- name: DeleteCustomerById :exec
DELETE FROM
    "Customer"
WHERE
    "CustomerId" = sqlc.arg('CustomerId');