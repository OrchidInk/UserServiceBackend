// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: detailMn.sql

package db

import (
	"context"
)

const createDetailMn = `-- name: CreateDetailMn :one
INSERT INTO
    "detailMn" (
        "ProductMnID",
        "ChoiceName",
        "ChoiceValue"
    )
VALUES
    (
        $1,
        $2,
        $3
    ) RETURNING "detailMnId", "ProductMnID", "ChoiceName", "ChoiceValue"
`

type CreateDetailMnParams struct {
	ProductMnId int32
	ChoiceName  string
	ChoiceValue string
}

func (q *Queries) CreateDetailMn(ctx context.Context, arg CreateDetailMnParams) (DetailMn, error) {
	row := q.db.QueryRowContext(ctx, createDetailMn, arg.ProductMnId, arg.ChoiceName, arg.ChoiceValue)
	var i DetailMn
	err := row.Scan(
		&i.DetailMnId,
		&i.ProductMnID,
		&i.ChoiceName,
		&i.ChoiceValue,
	)
	return i, err
}

const deleteDetailMn = `-- name: DeleteDetailMn :exec
DELETE FROM
    "detailMn"
WHERE
    "detailMnId" = $1
`

func (q *Queries) DeleteDetailMn(ctx context.Context, detailmnid int32) error {
	_, err := q.db.ExecContext(ctx, deleteDetailMn, detailmnid)
	return err
}

const findByDetailMnID = `-- name: FindByDetailMnID :one
SELECT
    "detailMnId", "ProductMnID", "ChoiceName", "ChoiceValue"
FROM
    "detailMn"
WHERE
    "detailMnId" = $1
LIMIT
    1
`

func (q *Queries) FindByDetailMnID(ctx context.Context, detailmnid int32) (DetailMn, error) {
	row := q.db.QueryRowContext(ctx, findByDetailMnID, detailmnid)
	var i DetailMn
	err := row.Scan(
		&i.DetailMnId,
		&i.ProductMnID,
		&i.ChoiceName,
		&i.ChoiceValue,
	)
	return i, err
}

const getAllDetailsMn = `-- name: GetAllDetailsMn :many
SELECT
    "detailMnId",
    "ProductMnID",
    "ChoiceName",
    "ChoiceValue"
FROM
    "detailMn"
ORDER BY
    "ProductMnID"
`

func (q *Queries) GetAllDetailsMn(ctx context.Context) ([]DetailMn, error) {
	rows, err := q.db.QueryContext(ctx, getAllDetailsMn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DetailMn
	for rows.Next() {
		var i DetailMn
		if err := rows.Scan(
			&i.DetailMnId,
			&i.ProductMnID,
			&i.ChoiceName,
			&i.ChoiceValue,
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

const getDetailsByProductMnID = `-- name: GetDetailsByProductMnID :many
SELECT
    "detailMnId",
    "ProductMnID",
    "ChoiceName",
    "ChoiceValue"
FROM
    "detailMn"
WHERE
    "ProductMnID" = $1
`

func (q *Queries) GetDetailsByProductMnID(ctx context.Context, productmnid int32) ([]DetailMn, error) {
	rows, err := q.db.QueryContext(ctx, getDetailsByProductMnID, productmnid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DetailMn
	for rows.Next() {
		var i DetailMn
		if err := rows.Scan(
			&i.DetailMnId,
			&i.ProductMnID,
			&i.ChoiceName,
			&i.ChoiceValue,
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

const updateDetailMn = `-- name: UpdateDetailMn :one
UPDATE
    "detailMn"
SET
    "ChoiceName" = $1,
    "ChoiceValue" = $2
WHERE
    "detailMnId" = $3 RETURNING "detailMnId", "ProductMnID", "ChoiceName", "ChoiceValue"
`

type UpdateDetailMnParams struct {
	ChoiceName  string
	ChoiceValue string
	DetailMnId  int32
}

func (q *Queries) UpdateDetailMn(ctx context.Context, arg UpdateDetailMnParams) (DetailMn, error) {
	row := q.db.QueryRowContext(ctx, updateDetailMn, arg.ChoiceName, arg.ChoiceValue, arg.DetailMnId)
	var i DetailMn
	err := row.Scan(
		&i.DetailMnId,
		&i.ProductMnID,
		&i.ChoiceName,
		&i.ChoiceValue,
	)
	return i, err
}
