-- name: CreatePayment :one
INSERT INTO "payments" (
    "OrderID",
    "UserID",
    "PaymentMethod",
    "PaymentStatus",
    "Amount"
) VALUES (
    sqlc.arg('OrderID'),
    sqlc.arg('UserID'),
    sqlc.arg('PaymentMethod'),
    sqlc.arg('PaymentStatus'),
    sqlc.arg('Amount')
) RETURNING *;


-- name: GetListPayment :many
SELECT
    *
FROM
    "payments";

-- name: UpdatePaymentStatus :one
UPDATE
    "payments"
SET
    "PaymentStatus" = sqlc.arg('PaymentStatus')
WHERE
    "PaymentID" = sqlc.arg('PaymentID') RETURNING *;

-- name: DeletePayment :exec
DELETE FROM "payments"
WHERE
    "PaymentID" = sqlc.arg('PaymentID');

-- name: FindByPaymentsId :one
SELECT
    *
FROM
    "payments"
WHERE
    "PaymentID" = sqlc.arg('PaymentID')
LIMIT 
1;