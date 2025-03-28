// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: image.sql

package db

import (
	"context"
)

const createProductImageEn = `-- name: CreateProductImageEn :one
INSERT INTO "productImagesEn" (
    "ProductEnID",
    "ImagePath"
)
VALUES (
    $1,
    $2
)
RETURNING "ImageID", "ProductEnID", "ImagePath", "Created_At"
`

type CreateProductImageEnParams struct {
	ProductEnID int32
	ImagePath   string
}

func (q *Queries) CreateProductImageEn(ctx context.Context, arg CreateProductImageEnParams) (ProductImagesEn, error) {
	row := q.db.QueryRowContext(ctx, createProductImageEn, arg.ProductEnID, arg.ImagePath)
	var i ProductImagesEn
	err := row.Scan(
		&i.ImageID,
		&i.ProductEnID,
		&i.ImagePath,
		&i.CreatedAt,
	)
	return i, err
}

const createProductImageMn = `-- name: CreateProductImageMn :one
INSERT INTO "productImagesMn" (
    "ProductMnID",
    "ImagePath"
)
VALUES (
    $1,
    $2
)
RETURNING "ImageID", "ProductMnID", "ImagePath", "Created_At"
`

type CreateProductImageMnParams struct {
	ProductMnID int32
	ImagePath   string
}

func (q *Queries) CreateProductImageMn(ctx context.Context, arg CreateProductImageMnParams) (ProductImagesMn, error) {
	row := q.db.QueryRowContext(ctx, createProductImageMn, arg.ProductMnID, arg.ImagePath)
	var i ProductImagesMn
	err := row.Scan(
		&i.ImageID,
		&i.ProductMnID,
		&i.ImagePath,
		&i.CreatedAt,
	)
	return i, err
}

const deleteProductImageEn = `-- name: DeleteProductImageEn :exec
DELETE FROM "productImagesEn"
WHERE "ImageID" = $1
`

func (q *Queries) DeleteProductImageEn(ctx context.Context, imageid int32) error {
	_, err := q.db.ExecContext(ctx, deleteProductImageEn, imageid)
	return err
}

const deleteProductImageMn = `-- name: DeleteProductImageMn :exec
DELETE FROM "productImagesMn"
WHERE "ImageID" = $1
`

func (q *Queries) DeleteProductImageMn(ctx context.Context, imageid int32) error {
	_, err := q.db.ExecContext(ctx, deleteProductImageMn, imageid)
	return err
}

const getAllProductImagesEn = `-- name: GetAllProductImagesEn :many
SELECT
    "ImageID", "ProductEnID", "ImagePath", "Created_At"
FROM
    "productImagesEn"
`

func (q *Queries) GetAllProductImagesEn(ctx context.Context) ([]ProductImagesEn, error) {
	rows, err := q.db.QueryContext(ctx, getAllProductImagesEn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductImagesEn
	for rows.Next() {
		var i ProductImagesEn
		if err := rows.Scan(
			&i.ImageID,
			&i.ProductEnID,
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

const getProductImagesEnByProductID = `-- name: GetProductImagesEnByProductID :many
SELECT "ImageID", "ProductEnID", "ImagePath", "Created_At"
FROM "productImagesEn"
WHERE "ProductEnID" = $1
ORDER BY "Created_At" ASC
`

func (q *Queries) GetProductImagesEnByProductID(ctx context.Context, productenid int32) ([]ProductImagesEn, error) {
	rows, err := q.db.QueryContext(ctx, getProductImagesEnByProductID, productenid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductImagesEn
	for rows.Next() {
		var i ProductImagesEn
		if err := rows.Scan(
			&i.ImageID,
			&i.ProductEnID,
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

const getProductImagesMn = `-- name: GetProductImagesMn :many
SELECT
    "ImageID", "ProductMnID", "ImagePath", "Created_At"
FROM
    "productImagesMn"
`

func (q *Queries) GetProductImagesMn(ctx context.Context) ([]ProductImagesMn, error) {
	rows, err := q.db.QueryContext(ctx, getProductImagesMn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductImagesMn
	for rows.Next() {
		var i ProductImagesMn
		if err := rows.Scan(
			&i.ImageID,
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

const getProductImagesMnByProductID = `-- name: GetProductImagesMnByProductID :many
SELECT "ImageID", "ProductMnID", "ImagePath", "Created_At"
FROM "productImagesMn"
WHERE "ProductMnID" = $1
ORDER BY "Created_At" ASC
`

func (q *Queries) GetProductImagesMnByProductID(ctx context.Context, productmnid int32) ([]ProductImagesMn, error) {
	rows, err := q.db.QueryContext(ctx, getProductImagesMnByProductID, productmnid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductImagesMn
	for rows.Next() {
		var i ProductImagesMn
		if err := rows.Scan(
			&i.ImageID,
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

const updateProductImageEn = `-- name: UpdateProductImageEn :one
UPDATE "productImagesEn"
SET "ImagePath" = $1
WHERE "ImageID" = $2
RETURNING "ImageID", "ProductEnID", "ImagePath", "Created_At"
`

type UpdateProductImageEnParams struct {
	ImagePath string
	ImageID   int32
}

func (q *Queries) UpdateProductImageEn(ctx context.Context, arg UpdateProductImageEnParams) (ProductImagesEn, error) {
	row := q.db.QueryRowContext(ctx, updateProductImageEn, arg.ImagePath, arg.ImageID)
	var i ProductImagesEn
	err := row.Scan(
		&i.ImageID,
		&i.ProductEnID,
		&i.ImagePath,
		&i.CreatedAt,
	)
	return i, err
}

const updateProductImageMn = `-- name: UpdateProductImageMn :one
UPDATE "productImagesMn"
SET "ImagePath" = $1
WHERE "ImageID" = $2
RETURNING "ImageID", "ProductMnID", "ImagePath", "Created_At"
`

type UpdateProductImageMnParams struct {
	ImagePath string
	ImageID   int32
}

func (q *Queries) UpdateProductImageMn(ctx context.Context, arg UpdateProductImageMnParams) (ProductImagesMn, error) {
	row := q.db.QueryRowContext(ctx, updateProductImageMn, arg.ImagePath, arg.ImageID)
	var i ProductImagesMn
	err := row.Scan(
		&i.ImageID,
		&i.ProductMnID,
		&i.ImagePath,
		&i.CreatedAt,
	)
	return i, err
}
