-- name: CreateDetailMn :one
INSERT INTO
    "detailMn" (
        "ProductMnID",
        "ChoiceName",
        "ChoiceValue"
    )
VALUES
    (
        sqlc.arg('ProductMnId'),
        sqlc.arg('ChoiceName'),
        sqlc.arg('ChoiceValue')
    ) RETURNING *;

-- name: GetDetailsByProductMnID :many
SELECT
    "detailMnId",
    "ProductMnID",
    "ChoiceName",
    "ChoiceValue"
FROM
    "detailMn"
WHERE
    "ProductMnID" = sqlc.arg('ProductMnID');

-- name: UpdateDetailMn :one
UPDATE
    "detailMn"
SET
    "ChoiceName" = sqlc.arg('ChoiceName'),
    "ChoiceValue" = sqlc.arg('ChoiceValue')
WHERE
    "detailMnId" = sqlc.arg('detailMnId') RETURNING *;

-- name: DeleteDetailMn :exec
DELETE FROM
    "detailMn"
WHERE
    "detailMnId" = sqlc.arg('detailMnId');

-- name: GetAllDetailsMn :many
SELECT
    "detailMnId",
    "ProductMnID",
    "ChoiceName",
    "ChoiceValue"
FROM
    "detailMn"
ORDER BY
    "ProductMnID";

-- name: FindByDetailMnID :one
SELECT
    *
FROM
    "detailMn"
WHERE
    "detailMnId" = sqlc.arg('detailMnId')
LIMIT
    1;