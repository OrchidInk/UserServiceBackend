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
    "IsActive" = sqlc.arg('IsActive'),
    "ContractEndDate" = sqlc.arg('ContractEndDate')
WHERE
    "CustomerId" = sqlc.arg('CustomerId') RETURNING *;

-- name: UpdateCustomerContractDates :one
UPDATE
    "Customer"
SET
    "ContractStartDate" = sqlc.arg('ContractStartDate')
WHERE
    "CustomerId" = sqlc.arg('CustomerId') RETURNING *;

-- name: DeleteCustomerById :exec
DELETE FROM
    "Customer"
WHERE
    "CustomerId" = sqlc.arg('CustomerId');

-- name: GetCustomerCountByStatus :many
SELECT
    "IsActive" AS status,
    COUNT(*) AS count
FROM
    "Customer"
GROUP BY
    "IsActive";



-- name: GetExpiredContracts :many
SELECT
    *
FROM
    "Customer"
WHERE
    "ContractEndDate" < CURRENT_DATE;

-- name: GetContractsEndingSoon :many
SELECT
    *
FROM
    "Customer"
WHERE
    "ContractEndDate" BETWEEN CURRENT_DATE AND (CURRENT_DATE + INTERVAL '30 days');

-- name: GetCustomerStatusOverTime :many
SELECT
    DATE_TRUNC('month', "Created_At") AS month,
    SUM(CASE WHEN "IsActive" = TRUE THEN 1 ELSE 0 END) AS active_customers,
    SUM(CASE WHEN "IsActive" = FALSE THEN 1 ELSE 0 END) AS inactive_customers
FROM
    "Customer"
GROUP BY
    DATE_TRUNC('month', "Created_At")
ORDER BY
    month ASC;
