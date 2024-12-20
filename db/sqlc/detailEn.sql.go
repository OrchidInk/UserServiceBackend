// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: detailEn.sql

package db

import (
	"context"
)

const createDetailEn = `-- name: CreateDetailEn :one
INSERT INTO
    "detailEn" (
        "ProductEnID",
        "ChoiceName",
        "ChoiceValue"
    )
VALUES
    (
        $1,
        $2,
        $3
    ) RETURNING "detailEnId", "ProductEnID", "ChoiceName", "ChoiceValue"
`

type CreateDetailEnParams struct {
	ProductEnID int32
	ChoiceName  string
	ChoiceValue string
}

func (q *Queries) CreateDetailEn(ctx context.Context, arg CreateDetailEnParams) (DetailEn, error) {
	row := q.db.QueryRowContext(ctx, createDetailEn, arg.ProductEnID, arg.ChoiceName, arg.ChoiceValue)
	var i DetailEn
	err := row.Scan(
		&i.DetailEnId,
		&i.ProductEnID,
		&i.ChoiceName,
		&i.ChoiceValue,
	)
	return i, err
}

const deleteDetailEn = `-- name: DeleteDetailEn :exec
DELETE FROM
    "detailEn"
WHERE
    "detailEnId" = $1
`

func (q *Queries) DeleteDetailEn(ctx context.Context, detailenid int32) error {
	_, err := q.db.ExecContext(ctx, deleteDetailEn, detailenid)
	return err
}

const findByDetailEn = `-- name: FindByDetailEn :one
SELECT
    "detailEnId", "ProductEnID", "ChoiceName", "ChoiceValue"
FROM
    "detailEn"
WHERE
    "detailEnId" = $1
LIMIT
    1
`

func (q *Queries) FindByDetailEn(ctx context.Context, detailenid int32) (DetailEn, error) {
	row := q.db.QueryRowContext(ctx, findByDetailEn, detailenid)
	var i DetailEn
	err := row.Scan(
		&i.DetailEnId,
		&i.ProductEnID,
		&i.ChoiceName,
		&i.ChoiceValue,
	)
	return i, err
}

const getAllDetailsEn = `-- name: GetAllDetailsEn :many
SELECT
    "detailEnId",
    "ProductEnID",
    "ChoiceName",
    "ChoiceValue"
FROM
    "detailEn"
ORDER BY
    "ProductEnID"
`

func (q *Queries) GetAllDetailsEn(ctx context.Context) ([]DetailEn, error) {
	rows, err := q.db.QueryContext(ctx, getAllDetailsEn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DetailEn
	for rows.Next() {
		var i DetailEn
		if err := rows.Scan(
			&i.DetailEnId,
			&i.ProductEnID,
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

const getDetailsByProductEnID = `-- name: GetDetailsByProductEnID :many
SELECT
    "detailEnId",
    "ProductEnID",
    "ChoiceName",
    "ChoiceValue"
FROM
    "detailEn"
WHERE
    "ProductEnID" = $1
`

func (q *Queries) GetDetailsByProductEnID(ctx context.Context, productenid int32) ([]DetailEn, error) {
	rows, err := q.db.QueryContext(ctx, getDetailsByProductEnID, productenid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DetailEn
	for rows.Next() {
		var i DetailEn
		if err := rows.Scan(
			&i.DetailEnId,
			&i.ProductEnID,
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

const updateDetailEn = `-- name: UpdateDetailEn :one
UPDATE
    "detailEn"
SET
    "ChoiceName" = $1,
    "ChoiceValue" = $2
WHERE
    "detailEnId" = $3 RETURNING "detailEnId", "ProductEnID", "ChoiceName", "ChoiceValue"
`

type UpdateDetailEnParams struct {
	ChoiceName  string
	ChoiceValue string
	DetailEnId  int32
}

func (q *Queries) UpdateDetailEn(ctx context.Context, arg UpdateDetailEnParams) (DetailEn, error) {
	row := q.db.QueryRowContext(ctx, updateDetailEn, arg.ChoiceName, arg.ChoiceValue, arg.DetailEnId)
	var i DetailEn
	err := row.Scan(
		&i.DetailEnId,
		&i.ProductEnID,
		&i.ChoiceName,
		&i.ChoiceValue,
	)
	return i, err
}
