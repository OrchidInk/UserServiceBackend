-- name: CreateDeliveryAddress :one
INSERT INTO
    "deliveryAddress" (
        "DeliverId",
        "Street",
        "City",
        "State",
        "PostalCode"
    )
VALUES
    () RETURNING *;