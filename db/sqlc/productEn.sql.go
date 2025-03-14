// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: productEn.sql

package db

import (
	"context"
	"database/sql"
)

const createProductEn = `-- name: CreateProductEn :one
INSERT INTO
    "productEn" (
        "ProductNameEn",
        "sCategoryIdEn",
        "PriceEn",
        "StockQuantity",
        "ImagesPathEn",
        "DescriptionEn",
        "BrandEn",
        "ManufacturedCountryEn",
        "ColorEn",
        "SizeEn",
        "PenOutputEn",
        "FeaturesEn",
        "MaterialEn",
        "StapleSizeEn",
        "CapacityEn",
        "WeightEn",
        "ThicknessEn",
        "PackagingEn",
        "ProductCodeEn",
        "CostPriceEn",
        "RetailPriceEn",
        "WarehouseStockEn"
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20,
        $21,
        $22
    ) RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "ColorEn", "SizeEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
`

type CreateProductEnParams struct {
	ProductNameEn         string
	SCategoryIdEn         int32
	PriceEn               string
	StockQuantity         int32
	ImagesPathEn          string
	DescriptionEn         string
	BrandEn               string
	ManufacturedCountryEn string
	ColorEn               string
	SizeEn                string
	PenOutputEn           string
	FeaturesEn            string
	MaterialEn            string
	StapleSizeEn          string
	CapacityEn            string
	WeightEn              string
	ThicknessEn           string
	PackagingEn           string
	ProductCodeEn         string
	CostPriceEn           string
	RetailPriceEn         string
	WarehouseStockEn      int32
}

func (q *Queries) CreateProductEn(ctx context.Context, arg CreateProductEnParams) (ProductEn, error) {
	row := q.db.QueryRowContext(ctx, createProductEn,
		arg.ProductNameEn,
		arg.SCategoryIdEn,
		arg.PriceEn,
		arg.StockQuantity,
		arg.ImagesPathEn,
		arg.DescriptionEn,
		arg.BrandEn,
		arg.ManufacturedCountryEn,
		arg.ColorEn,
		arg.SizeEn,
		arg.PenOutputEn,
		arg.FeaturesEn,
		arg.MaterialEn,
		arg.StapleSizeEn,
		arg.CapacityEn,
		arg.WeightEn,
		arg.ThicknessEn,
		arg.PackagingEn,
		arg.ProductCodeEn,
		arg.CostPriceEn,
		arg.RetailPriceEn,
		arg.WarehouseStockEn,
	)
	var i ProductEn
	err := row.Scan(
		&i.ProductEnID,
		&i.ProductNameEn,
		&i.SCategoryIdEn,
		&i.PriceEn,
		&i.StockQuantity,
		&i.ImagesPathEn,
		&i.DescriptionEn,
		&i.BrandEn,
		&i.ManufacturedCountryEn,
		&i.ColorEn,
		&i.SizeEn,
		&i.PenOutputEn,
		&i.FeaturesEn,
		&i.MaterialEn,
		&i.StapleSizeEn,
		&i.CapacityEn,
		&i.WeightEn,
		&i.ThicknessEn,
		&i.PackagingEn,
		&i.UsageEn,
		&i.InstructionsEn,
		&i.ProductCodeEn,
		&i.CostPriceEn,
		&i.RetailPriceEn,
		&i.WarehouseStockEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deductSockQuantityByProductEnID = `-- name: DeductSockQuantityByProductEnID :one
UPDATE
    "productEn"
SET
    "StockQuantity" = "StockQuantity" - $1
WHERE
    "ProductEnID" = $2
    AND "StockQuantity" >= $1 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "ColorEn", "SizeEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
`

type DeductSockQuantityByProductEnIDParams struct {
	QuantityPurchased int32
	ProductEnID       int32
}

func (q *Queries) DeductSockQuantityByProductEnID(ctx context.Context, arg DeductSockQuantityByProductEnIDParams) (ProductEn, error) {
	row := q.db.QueryRowContext(ctx, deductSockQuantityByProductEnID, arg.QuantityPurchased, arg.ProductEnID)
	var i ProductEn
	err := row.Scan(
		&i.ProductEnID,
		&i.ProductNameEn,
		&i.SCategoryIdEn,
		&i.PriceEn,
		&i.StockQuantity,
		&i.ImagesPathEn,
		&i.DescriptionEn,
		&i.BrandEn,
		&i.ManufacturedCountryEn,
		&i.ColorEn,
		&i.SizeEn,
		&i.PenOutputEn,
		&i.FeaturesEn,
		&i.MaterialEn,
		&i.StapleSizeEn,
		&i.CapacityEn,
		&i.WeightEn,
		&i.ThicknessEn,
		&i.PackagingEn,
		&i.UsageEn,
		&i.InstructionsEn,
		&i.ProductCodeEn,
		&i.CostPriceEn,
		&i.RetailPriceEn,
		&i.WarehouseStockEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteByProductEnId = `-- name: DeleteByProductEnId :exec
DELETE FROM
    "productEn"
WHERE
    "ProductEnID" = $1
`

func (q *Queries) DeleteByProductEnId(ctx context.Context, productenid int32) error {
	_, err := q.db.ExecContext(ctx, deleteByProductEnId, productenid)
	return err
}

const filterByProductEnName = `-- name: FilterByProductEnName :many
SELECT
    "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "ColorEn", "SizeEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
FROM
    "productEn"
WHERE
    "ProductNameEn" ILIKE '%' || $1 || '%' -- Case-insensitive search for partial match
ORDER BY
    "Created_At" DESC
`

func (q *Queries) FilterByProductEnName(ctx context.Context, productnameen sql.NullString) ([]ProductEn, error) {
	rows, err := q.db.QueryContext(ctx, filterByProductEnName, productnameen)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductEn
	for rows.Next() {
		var i ProductEn
		if err := rows.Scan(
			&i.ProductEnID,
			&i.ProductNameEn,
			&i.SCategoryIdEn,
			&i.PriceEn,
			&i.StockQuantity,
			&i.ImagesPathEn,
			&i.DescriptionEn,
			&i.BrandEn,
			&i.ManufacturedCountryEn,
			&i.ColorEn,
			&i.SizeEn,
			&i.PenOutputEn,
			&i.FeaturesEn,
			&i.MaterialEn,
			&i.StapleSizeEn,
			&i.CapacityEn,
			&i.WeightEn,
			&i.ThicknessEn,
			&i.PackagingEn,
			&i.UsageEn,
			&i.InstructionsEn,
			&i.ProductCodeEn,
			&i.CostPriceEn,
			&i.RetailPriceEn,
			&i.WarehouseStockEn,
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

const findByProductIdEn = `-- name: FindByProductIdEn :one
SELECT
    "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "ColorEn", "SizeEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
FROM
    "productEn"
WHERE
    "ProductEnID" = $1
LIMIT
    1
`

func (q *Queries) FindByProductIdEn(ctx context.Context, productenid int32) (ProductEn, error) {
	row := q.db.QueryRowContext(ctx, findByProductIdEn, productenid)
	var i ProductEn
	err := row.Scan(
		&i.ProductEnID,
		&i.ProductNameEn,
		&i.SCategoryIdEn,
		&i.PriceEn,
		&i.StockQuantity,
		&i.ImagesPathEn,
		&i.DescriptionEn,
		&i.BrandEn,
		&i.ManufacturedCountryEn,
		&i.ColorEn,
		&i.SizeEn,
		&i.PenOutputEn,
		&i.FeaturesEn,
		&i.MaterialEn,
		&i.StapleSizeEn,
		&i.CapacityEn,
		&i.WeightEn,
		&i.ThicknessEn,
		&i.PackagingEn,
		&i.UsageEn,
		&i.InstructionsEn,
		&i.ProductCodeEn,
		&i.CostPriceEn,
		&i.RetailPriceEn,
		&i.WarehouseStockEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getListProductEn = `-- name: GetListProductEn :many
SELECT
    "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "ColorEn", "SizeEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
FROM
    "productEn"
ORDER BY
    "Created_At" DESC
`

func (q *Queries) GetListProductEn(ctx context.Context) ([]ProductEn, error) {
	rows, err := q.db.QueryContext(ctx, getListProductEn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductEn
	for rows.Next() {
		var i ProductEn
		if err := rows.Scan(
			&i.ProductEnID,
			&i.ProductNameEn,
			&i.SCategoryIdEn,
			&i.PriceEn,
			&i.StockQuantity,
			&i.ImagesPathEn,
			&i.DescriptionEn,
			&i.BrandEn,
			&i.ManufacturedCountryEn,
			&i.ColorEn,
			&i.SizeEn,
			&i.PenOutputEn,
			&i.FeaturesEn,
			&i.MaterialEn,
			&i.StapleSizeEn,
			&i.CapacityEn,
			&i.WeightEn,
			&i.ThicknessEn,
			&i.PackagingEn,
			&i.UsageEn,
			&i.InstructionsEn,
			&i.ProductCodeEn,
			&i.CostPriceEn,
			&i.RetailPriceEn,
			&i.WarehouseStockEn,
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

const updateByEnImagePath = `-- name: UpdateByEnImagePath :one
UPDATE
    "productEn"
SET
    "ImagesPathEn" = $1
WHERE
    "ProductEnID" = $2 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "ColorEn", "SizeEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
`

type UpdateByEnImagePathParams struct {
	ImagesPathEn string
	ProductEnID  int32
}

func (q *Queries) UpdateByEnImagePath(ctx context.Context, arg UpdateByEnImagePathParams) (ProductEn, error) {
	row := q.db.QueryRowContext(ctx, updateByEnImagePath, arg.ImagesPathEn, arg.ProductEnID)
	var i ProductEn
	err := row.Scan(
		&i.ProductEnID,
		&i.ProductNameEn,
		&i.SCategoryIdEn,
		&i.PriceEn,
		&i.StockQuantity,
		&i.ImagesPathEn,
		&i.DescriptionEn,
		&i.BrandEn,
		&i.ManufacturedCountryEn,
		&i.ColorEn,
		&i.SizeEn,
		&i.PenOutputEn,
		&i.FeaturesEn,
		&i.MaterialEn,
		&i.StapleSizeEn,
		&i.CapacityEn,
		&i.WeightEn,
		&i.ThicknessEn,
		&i.PackagingEn,
		&i.UsageEn,
		&i.InstructionsEn,
		&i.ProductCodeEn,
		&i.CostPriceEn,
		&i.RetailPriceEn,
		&i.WarehouseStockEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateByProductEnPrice = `-- name: UpdateByProductEnPrice :one
UPDATE
    "productEn"
SET
    "PriceEn" = $1
WHERE
    "ProductEnID" = $2 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "ColorEn", "SizeEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
`

type UpdateByProductEnPriceParams struct {
	PriceEn     string
	ProductEnID int32
}

func (q *Queries) UpdateByProductEnPrice(ctx context.Context, arg UpdateByProductEnPriceParams) (ProductEn, error) {
	row := q.db.QueryRowContext(ctx, updateByProductEnPrice, arg.PriceEn, arg.ProductEnID)
	var i ProductEn
	err := row.Scan(
		&i.ProductEnID,
		&i.ProductNameEn,
		&i.SCategoryIdEn,
		&i.PriceEn,
		&i.StockQuantity,
		&i.ImagesPathEn,
		&i.DescriptionEn,
		&i.BrandEn,
		&i.ManufacturedCountryEn,
		&i.ColorEn,
		&i.SizeEn,
		&i.PenOutputEn,
		&i.FeaturesEn,
		&i.MaterialEn,
		&i.StapleSizeEn,
		&i.CapacityEn,
		&i.WeightEn,
		&i.ThicknessEn,
		&i.PackagingEn,
		&i.UsageEn,
		&i.InstructionsEn,
		&i.ProductCodeEn,
		&i.CostPriceEn,
		&i.RetailPriceEn,
		&i.WarehouseStockEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateByProductEnStockQuantity = `-- name: UpdateByProductEnStockQuantity :one
UPDATE
    "productEn"
SET
    "StockQuantity" = $1
WHERE
    "ProductEnID" = $2 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "ColorEn", "SizeEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
`

type UpdateByProductEnStockQuantityParams struct {
	StockQuantity int32
	ProductEnID   int32
}

func (q *Queries) UpdateByProductEnStockQuantity(ctx context.Context, arg UpdateByProductEnStockQuantityParams) (ProductEn, error) {
	row := q.db.QueryRowContext(ctx, updateByProductEnStockQuantity, arg.StockQuantity, arg.ProductEnID)
	var i ProductEn
	err := row.Scan(
		&i.ProductEnID,
		&i.ProductNameEn,
		&i.SCategoryIdEn,
		&i.PriceEn,
		&i.StockQuantity,
		&i.ImagesPathEn,
		&i.DescriptionEn,
		&i.BrandEn,
		&i.ManufacturedCountryEn,
		&i.ColorEn,
		&i.SizeEn,
		&i.PenOutputEn,
		&i.FeaturesEn,
		&i.MaterialEn,
		&i.StapleSizeEn,
		&i.CapacityEn,
		&i.WeightEn,
		&i.ThicknessEn,
		&i.PackagingEn,
		&i.UsageEn,
		&i.InstructionsEn,
		&i.ProductCodeEn,
		&i.CostPriceEn,
		&i.RetailPriceEn,
		&i.WarehouseStockEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProductEn = `-- name: UpdateProductEn :one
UPDATE 
    "productEn"
SET
    "ProductNameEn" = $1,
    "sCategoryIdEn" = $2,
    "PriceEn" = $3,
    "StockQuantity" = $4,
    "ImagesPathEn" = $5,
    "DescriptionEn" = $6,
    "BrandEn" = $7,
    "ManufacturedCountryEn" = $8,
    "ColorEn" = $9,
    "SizeEn" = $10,
    "PenOutputEn" = $11,
    "FeaturesEn" = $12,
    "MaterialEn" = $13,
    "StapleSizeEn" = $14,
    "CapacityEn" = $15,
    "WeightEn" = $16,
    "ThicknessEn" = $17,
    "PackagingEn" = $18,
    "UsageEn" = $19,
    "InstructionsEn" = $20,
    "ProductCodeEn" = $21,
    "CostPriceEn" = $22,
    "RetailPriceEn" = $23,
    "WarehouseStockEn" = $24
    WHERE
    "ProductEnID" = $25 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "ColorEn", "SizeEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
`

type UpdateProductEnParams struct {
	ProductNameEn         string
	SCategoryIdEn         int32
	PriceEn               string
	StockQuantity         int32
	ImagesPathEn          string
	DescriptionEn         string
	BrandEn               string
	ManufacturedCountryEn string
	ColorEn               string
	SizeEn                string
	PenOutputEn           string
	FeaturesEn            string
	MaterialEn            string
	StapleSizeEn          string
	CapacityEn            string
	WeightEn              string
	ThicknessEn           string
	PackagingEn           string
	UsageEn               string
	InstructionsEn        string
	ProductCodeEn         string
	CostPriceEn           string
	RetailPriceEn         string
	WarehouseStockEn      int32
	ProductEnID           int32
}

func (q *Queries) UpdateProductEn(ctx context.Context, arg UpdateProductEnParams) (ProductEn, error) {
	row := q.db.QueryRowContext(ctx, updateProductEn,
		arg.ProductNameEn,
		arg.SCategoryIdEn,
		arg.PriceEn,
		arg.StockQuantity,
		arg.ImagesPathEn,
		arg.DescriptionEn,
		arg.BrandEn,
		arg.ManufacturedCountryEn,
		arg.ColorEn,
		arg.SizeEn,
		arg.PenOutputEn,
		arg.FeaturesEn,
		arg.MaterialEn,
		arg.StapleSizeEn,
		arg.CapacityEn,
		arg.WeightEn,
		arg.ThicknessEn,
		arg.PackagingEn,
		arg.UsageEn,
		arg.InstructionsEn,
		arg.ProductCodeEn,
		arg.CostPriceEn,
		arg.RetailPriceEn,
		arg.WarehouseStockEn,
		arg.ProductEnID,
	)
	var i ProductEn
	err := row.Scan(
		&i.ProductEnID,
		&i.ProductNameEn,
		&i.SCategoryIdEn,
		&i.PriceEn,
		&i.StockQuantity,
		&i.ImagesPathEn,
		&i.DescriptionEn,
		&i.BrandEn,
		&i.ManufacturedCountryEn,
		&i.ColorEn,
		&i.SizeEn,
		&i.PenOutputEn,
		&i.FeaturesEn,
		&i.MaterialEn,
		&i.StapleSizeEn,
		&i.CapacityEn,
		&i.WeightEn,
		&i.ThicknessEn,
		&i.PackagingEn,
		&i.UsageEn,
		&i.InstructionsEn,
		&i.ProductCodeEn,
		&i.CostPriceEn,
		&i.RetailPriceEn,
		&i.WarehouseStockEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProductEnSubCategory = `-- name: UpdateProductEnSubCategory :one
UPDATE
    "productEn"
SET
    "sCategoryIdEn" = $1
WHERE
    "ProductEnID" = $2 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "ColorEn", "SizeEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
`

type UpdateProductEnSubCategoryParams struct {
	SCategoryIdEn int32
	ProductEnID   int32
}

func (q *Queries) UpdateProductEnSubCategory(ctx context.Context, arg UpdateProductEnSubCategoryParams) (ProductEn, error) {
	row := q.db.QueryRowContext(ctx, updateProductEnSubCategory, arg.SCategoryIdEn, arg.ProductEnID)
	var i ProductEn
	err := row.Scan(
		&i.ProductEnID,
		&i.ProductNameEn,
		&i.SCategoryIdEn,
		&i.PriceEn,
		&i.StockQuantity,
		&i.ImagesPathEn,
		&i.DescriptionEn,
		&i.BrandEn,
		&i.ManufacturedCountryEn,
		&i.ColorEn,
		&i.SizeEn,
		&i.PenOutputEn,
		&i.FeaturesEn,
		&i.MaterialEn,
		&i.StapleSizeEn,
		&i.CapacityEn,
		&i.WeightEn,
		&i.ThicknessEn,
		&i.PackagingEn,
		&i.UsageEn,
		&i.InstructionsEn,
		&i.ProductCodeEn,
		&i.CostPriceEn,
		&i.RetailPriceEn,
		&i.WarehouseStockEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSProductEn = `-- name: UpdateSProductEn :one
UPDATE
    "productEn"
SET
    "StockQuantity" = $1
WHERE
    "ProductEnID" = $2 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "ColorEn", "SizeEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
`

type UpdateSProductEnParams struct {
	StockQuantity int32
	ProductEnID   int32
}

func (q *Queries) UpdateSProductEn(ctx context.Context, arg UpdateSProductEnParams) (ProductEn, error) {
	row := q.db.QueryRowContext(ctx, updateSProductEn, arg.StockQuantity, arg.ProductEnID)
	var i ProductEn
	err := row.Scan(
		&i.ProductEnID,
		&i.ProductNameEn,
		&i.SCategoryIdEn,
		&i.PriceEn,
		&i.StockQuantity,
		&i.ImagesPathEn,
		&i.DescriptionEn,
		&i.BrandEn,
		&i.ManufacturedCountryEn,
		&i.ColorEn,
		&i.SizeEn,
		&i.PenOutputEn,
		&i.FeaturesEn,
		&i.MaterialEn,
		&i.StapleSizeEn,
		&i.CapacityEn,
		&i.WeightEn,
		&i.ThicknessEn,
		&i.PackagingEn,
		&i.UsageEn,
		&i.InstructionsEn,
		&i.ProductCodeEn,
		&i.CostPriceEn,
		&i.RetailPriceEn,
		&i.WarehouseStockEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
