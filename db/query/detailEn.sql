-- name: CreateDetailEn :one
INSERT INTO
    "detailEn" (
        "ProductEnID",
        "ChoiceName",
        "ChoiceValue"
    )
VALUES
    (
        sqlc.arg('ProductEnID'),
        sqlc.arg('ChoiceName'),
        sqlc.arg('ChoiceValue')
    ) RETURNING *;

-- name: GetDetailsByProductEnID :many
SELECT
    "detailEnId",
    "ProductEnID",
    "ChoiceName",
    "ChoiceValue"
FROM
    "detailEn"
WHERE
    "ProductEnID" = sqlc.arg('ProductEnID');

-- name: UpdateDetailEn :one
UPDATE
    "detailEn"
SET
    "ChoiceName" = sqlc.arg('ChoiceName'),
    "ChoiceValue" = sqlc.arg('ChoiceValue')
WHERE
    "detailEnId" = sqlc.arg('detailEnId') RETURNING *;

-- name: DeleteDetailEn :exec
DELETE FROM
    "detailEn"
WHERE
    "detailEnId" = sqlc.arg('detailEnId');

-- name: GetAllDetailsEn :many
SELECT
    "detailEnId",
    "ProductEnID",
    "ChoiceName",
    "ChoiceValue"
FROM
    "detailEn"
ORDER BY
    "ProductEnID";

-- name: FindByDetailEn :one
SELECT
    *
FROM
    "detailEn"
WHERE
    "detailEnId" = sqlc.arg('detailEnId')
LIMIT
    1;