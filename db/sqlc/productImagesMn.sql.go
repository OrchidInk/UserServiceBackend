// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: productImagesMn.sql

package db

import (
	"context"
)

const createImageMn = `-- name: CreateImageMn :one
INSERT INTO
    "imagesMn" ("ProductMnID", "ImagePath")
VALUES
    (
        $1,
        $2
    ) RETURNING "Id", "ProductMnID", "ImagePath", "Created_At"
`

type CreateImageMnParams struct {
	ProductMnID int32
	ImagePath   string
}

func (q *Queries) CreateImageMn(ctx context.Context, arg CreateImageMnParams) (ImagesMn, error) {
	row := q.db.QueryRowContext(ctx, createImageMn, arg.ProductMnID, arg.ImagePath)
	var i ImagesMn
	err := row.Scan(
		&i.Id,
		&i.ProductMnID,
		&i.ImagePath,
		&i.CreatedAt,
	)
	return i, err
}

const deleteImagesMn = `-- name: DeleteImagesMn :exec
DELETE FROM
    "imagesMn"
WHERE
    "Id" = $1
`

func (q *Queries) DeleteImagesMn(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteImagesMn, id)
	return err
}

const getListImagesMn = `-- name: GetListImagesMn :many
SELECT
    "Id", "ProductMnID", "ImagePath", "Created_At"
FROM
    "imagesMn"
`

func (q *Queries) GetListImagesMn(ctx context.Context) ([]ImagesMn, error) {
	rows, err := q.db.QueryContext(ctx, getListImagesMn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ImagesMn
	for rows.Next() {
		var i ImagesMn
		if err := rows.Scan(
			&i.Id,
			&i.ProductMnID,
			&i.ImagePath,
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

const updateImagesMn = `-- name: UpdateImagesMn :one
UPDATE
    "imagesMn"
SET
    "ImagePath" = $1
WHERE
    "Id" = $2 RETURNING "Id", "ProductMnID", "ImagePath", "Created_At"
`

type UpdateImagesMnParams struct {
	ImagePath string
	Id        int32
}

func (q *Queries) UpdateImagesMn(ctx context.Context, arg UpdateImagesMnParams) (ImagesMn, error) {
	row := q.db.QueryRowContext(ctx, updateImagesMn, arg.ImagePath, arg.Id)
	var i ImagesMn
	err := row.Scan(
		&i.Id,
		&i.ProductMnID,
		&i.ImagePath,
		&i.CreatedAt,
	)
	return i, err
}
