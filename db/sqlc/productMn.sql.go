// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: productMn.sql

package db

import (
	"context"
	"database/sql"
)

const createProductMn = `-- name: CreateProductMn :one
INSERT INTO
    "productMn" (
        "ProductNameMn",
        "subCategoryIDMn",
        "PriceMn",
        "StockQuantity",
        "ImagesPathMn"
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5
    ) RETURNING "ProductMnID", "ProductNameMn", "subCategoryIDMn", "PriceMn", "StockQuantity", "ImagesPathMn", "Created_At", "Updated_At"
`

type CreateProductMnParams struct {
	ProductNameMn   string
	SubCategoryIDMn int32
	PriceMn         string
	StockQuantity   int32
	ImagesPathMn    string
}

func (q *Queries) CreateProductMn(ctx context.Context, arg CreateProductMnParams) (ProductMn, error) {
	row := q.db.QueryRowContext(ctx, createProductMn,
		arg.ProductNameMn,
		arg.SubCategoryIDMn,
		arg.PriceMn,
		arg.StockQuantity,
		arg.ImagesPathMn,
	)
	var i ProductMn
	err := row.Scan(
		&i.ProductMnID,
		&i.ProductNameMn,
		&i.SubCategoryIDMn,
		&i.PriceMn,
		&i.StockQuantity,
		&i.ImagesPathMn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deductSockQuantityByProductMnID = `-- name: DeductSockQuantityByProductMnID :one
UPDATE
    "productMn"
SET
    "StockQuantity" = "StockQuantity" - $1
WHERE
    "ProductMnID" = $2
    AND "StockQuantity" >= $1 RETURNING "ProductMnID", "ProductNameMn", "subCategoryIDMn", "PriceMn", "StockQuantity", "ImagesPathMn", "Created_At", "Updated_At"
`

type DeductSockQuantityByProductMnIDParams struct {
	QuantityPurchased int32
	ProductMnID       int32
}

func (q *Queries) DeductSockQuantityByProductMnID(ctx context.Context, arg DeductSockQuantityByProductMnIDParams) (ProductMn, error) {
	row := q.db.QueryRowContext(ctx, deductSockQuantityByProductMnID, arg.QuantityPurchased, arg.ProductMnID)
	var i ProductMn
	err := row.Scan(
		&i.ProductMnID,
		&i.ProductNameMn,
		&i.SubCategoryIDMn,
		&i.PriceMn,
		&i.StockQuantity,
		&i.ImagesPathMn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteByProductMnId = `-- name: DeleteByProductMnId :exec
DELETE FROM
    "productMn"
WHERE
    "ProductMnID" = $1
`

func (q *Queries) DeleteByProductMnId(ctx context.Context, productmnid int32) error {
	_, err := q.db.ExecContext(ctx, deleteByProductMnId, productmnid)
	return err
}

const filterByProductMnName = `-- name: FilterByProductMnName :many
SELECT
    "ProductMnID", "ProductNameMn", "subCategoryIDMn", "PriceMn", "StockQuantity", "ImagesPathMn", "Created_At", "Updated_At"
FROM
    "productMn"
WHERE
    "ProductMnName" ILIKE '%' || $1 || '%' -- Case-insensitive search for partial match
ORDER BY
    "Created_At" DESC
`

func (q *Queries) FilterByProductMnName(ctx context.Context, productmnname sql.NullString) ([]ProductMn, error) {
	rows, err := q.db.QueryContext(ctx, filterByProductMnName, productmnname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductMn
	for rows.Next() {
		var i ProductMn
		if err := rows.Scan(
			&i.ProductMnID,
			&i.ProductNameMn,
			&i.SubCategoryIDMn,
			&i.PriceMn,
			&i.StockQuantity,
			&i.ImagesPathMn,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const findByProductIdMn = `-- name: FindByProductIdMn :one
SELECT
    "ProductMnID", "ProductNameMn", "subCategoryIDMn", "PriceMn", "StockQuantity", "ImagesPathMn", "Created_At", "Updated_At"
FROM
    "productMn"
WHERE
    "ProductMnID" = $1
LIMIT
    1
`

func (q *Queries) FindByProductIdMn(ctx context.Context, productmnid int32) (ProductMn, error) {
	row := q.db.QueryRowContext(ctx, findByProductIdMn, productmnid)
	var i ProductMn
	err := row.Scan(
		&i.ProductMnID,
		&i.ProductNameMn,
		&i.SubCategoryIDMn,
		&i.PriceMn,
		&i.StockQuantity,
		&i.ImagesPathMn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getListProductMn = `-- name: GetListProductMn :many
SELECT
    "ProductMnID", "ProductNameMn", "subCategoryIDMn", "PriceMn", "StockQuantity", "ImagesPathMn", "Created_At", "Updated_At"
FROM
    "productMn"
ORDER BY
    "Created_At" DESC
`

func (q *Queries) GetListProductMn(ctx context.Context) ([]ProductMn, error) {
	rows, err := q.db.QueryContext(ctx, getListProductMn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductMn
	for rows.Next() {
		var i ProductMn
		if err := rows.Scan(
			&i.ProductMnID,
			&i.ProductNameMn,
			&i.SubCategoryIDMn,
			&i.PriceMn,
			&i.StockQuantity,
			&i.ImagesPathMn,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateByMnImagePath = `-- name: UpdateByMnImagePath :one
UPDATE
    "productMn"
SET
    "ImagesPathMn" = $1
WHERE
    "ProductMnID" = $2 RETURNING "ProductMnID", "ProductNameMn", "subCategoryIDMn", "PriceMn", "StockQuantity", "ImagesPathMn", "Created_At", "Updated_At"
`

type UpdateByMnImagePathParams struct {
	ImagesPathMn string
	ProductMnID  int32
}

func (q *Queries) UpdateByMnImagePath(ctx context.Context, arg UpdateByMnImagePathParams) (ProductMn, error) {
	row := q.db.QueryRowContext(ctx, updateByMnImagePath, arg.ImagesPathMn, arg.ProductMnID)
	var i ProductMn
	err := row.Scan(
		&i.ProductMnID,
		&i.ProductNameMn,
		&i.SubCategoryIDMn,
		&i.PriceMn,
		&i.StockQuantity,
		&i.ImagesPathMn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateByProductMnPrice = `-- name: UpdateByProductMnPrice :one
UPDATE
    "productMn"
SET
    "PriceMn" = $1
WHERE
    "ProductMnID" = $2 RETURNING "ProductMnID", "ProductNameMn", "subCategoryIDMn", "PriceMn", "StockQuantity", "ImagesPathMn", "Created_At", "Updated_At"
`

type UpdateByProductMnPriceParams struct {
	PriceMn     string
	ProductMnID int32
}

func (q *Queries) UpdateByProductMnPrice(ctx context.Context, arg UpdateByProductMnPriceParams) (ProductMn, error) {
	row := q.db.QueryRowContext(ctx, updateByProductMnPrice, arg.PriceMn, arg.ProductMnID)
	var i ProductMn
	err := row.Scan(
		&i.ProductMnID,
		&i.ProductNameMn,
		&i.SubCategoryIDMn,
		&i.PriceMn,
		&i.StockQuantity,
		&i.ImagesPathMn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateByProductMnStockQuantity = `-- name: UpdateByProductMnStockQuantity :one
UPDATE
    "productMn"
SET
    "StockQuantity" = $1
WHERE
    "ProductMnID" = $2 RETURNING "ProductMnID", "ProductNameMn", "subCategoryIDMn", "PriceMn", "StockQuantity", "ImagesPathMn", "Created_At", "Updated_At"
`

type UpdateByProductMnStockQuantityParams struct {
	StockQuantity int32
	ProductMnID   int32
}

func (q *Queries) UpdateByProductMnStockQuantity(ctx context.Context, arg UpdateByProductMnStockQuantityParams) (ProductMn, error) {
	row := q.db.QueryRowContext(ctx, updateByProductMnStockQuantity, arg.StockQuantity, arg.ProductMnID)
	var i ProductMn
	err := row.Scan(
		&i.ProductMnID,
		&i.ProductNameMn,
		&i.SubCategoryIDMn,
		&i.PriceMn,
		&i.StockQuantity,
		&i.ImagesPathMn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}