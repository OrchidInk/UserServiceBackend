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
        $20
    ) RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
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
    AND "StockQuantity" >= $1 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
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
    "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
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
    "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
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
    "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
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

const getProductEnWithAllColorsAndSizes = `-- name: GetProductEnWithAllColorsAndSizes :many


SELECT 
    p."ProductEnID",
    p."ProductNameEn",
    p."sCategoryIdEn",
    p."PriceEn",
    p."StockQuantity",
    p."ImagesPathEn",
    p."DescriptionEn",
    p."BrandEn",
    p."ManufacturedCountryEn",
    p."PenOutputEn",
    p."FeaturesEn",
    p."MaterialEn",
    p."StapleSizeEn",
    p."CapacityEn",
    p."WeightEn",
    p."ThicknessEn",
    p."PackagingEn",
    p."UsageEn",
    p."InstructionsEn",
    p."ProductCodeEn",
    p."CostPriceEn",
    p."RetailPriceEn",
    p."WarehouseStockEn",
    p."Created_At",
    p."Updated_At",
    COALESCE(
        ARRAY_AGG(DISTINCT c."ColorId") FILTER (WHERE c."ColorId" IS NOT NULL),
        '{}'
    ) AS "ColorIds",
    COALESCE(
        ARRAY_AGG(DISTINCT c."Color") FILTER (WHERE c."Color" IS NOT NULL),
        '{}'
    ) AS "ColorNames",
    COALESCE(
        ARRAY_AGG(DISTINCT s."SizeId") FILTER (WHERE s."SizeId" IS NOT NULL),
        '{}'
    ) AS "SizeIds",
    COALESCE(
        ARRAY_AGG(DISTINCT s."Size") FILTER (WHERE s."Size" IS NOT NULL),
        '{}'
    ) AS "SizeNames",
    COALESCE(
        ARRAY_AGG(DISTINCT pi."ImagePath") FILTER (WHERE pi."ImagePath" IS NOT NULL),
        '{}'
    ) AS "ImagePaths"
FROM "productEn" p
LEFT JOIN "productEn_colors" pc 
       ON p."ProductEnID" = pc."ProductEnID"
LEFT JOIN "productEn_sizes" ps 
       ON p."ProductEnID" = ps."ProductEnID"
LEFT JOIN "Color" c 
       ON pc."ColorId" = c."ColorId"
LEFT JOIN "Size" s 
       ON ps."SizeId" = s."SizeId"
LEFT JOIN "productImagesEn" pi 
       ON p."ProductEnID" = pi."ProductEnID"
GROUP BY p."ProductEnID"
`

type GetProductEnWithAllColorsAndSizesRow struct {
	ProductEnID           int32
	ProductNameEn         string
	SCategoryIdEn         int32
	PriceEn               string
	StockQuantity         int32
	ImagesPathEn          string
	DescriptionEn         string
	BrandEn               string
	ManufacturedCountryEn string
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
	CreatedAt             sql.NullTime
	UpdatedAt             sql.NullTime
	ColorIds              interface{}
	ColorNames            interface{}
	SizeIds               interface{}
	SizeNames             interface{}
	ImagePaths            interface{}
}

// -- name: GetProductEnWithDetails :many
// SELECT
//
//	p."ProductEnID",
//	p."ProductNameEn",
//	p."sCategoryIdEn",
//	p."PriceEn",
//	p."StockQuantity",
//	p."ImagesPathEn",
//	p."DescriptionEn",
//	p."BrandEn",
//	p."ManufacturedCountryEn",
//	p."ColorId",
//	c."Color" AS colorName,
//	p."SizeId",
//	s."Size" AS sizeName,
//	p."PenOutputEn",
//	p."FeaturesEn",
//	p."MaterialEn",
//	p."StapleSizeEn",
//	p."CapacityEn",
//	p."WeightEn",
//	p."ThicknessEn",
//	p."PackagingEn",
//	p."UsageEn",
//	p."InstructionsEn",
//	p."ProductCodeEn",
//	p."CostPriceEn",
//	p."RetailPriceEn",
//	p."WarehouseStockEn",
//	p."Created_At",
//	p."Updated_At"
//
// FROM "productEn" p
// LEFT JOIN "Color" c ON p."ColorId" = c."ColorId"
// LEFT JOIN "Size" s ON p."SizeId" = s."SizeId";
func (q *Queries) GetProductEnWithAllColorsAndSizes(ctx context.Context) ([]GetProductEnWithAllColorsAndSizesRow, error) {
	rows, err := q.db.QueryContext(ctx, getProductEnWithAllColorsAndSizes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductEnWithAllColorsAndSizesRow
	for rows.Next() {
		var i GetProductEnWithAllColorsAndSizesRow
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
			&i.ColorIds,
			&i.ColorNames,
			&i.SizeIds,
			&i.SizeNames,
			&i.ImagePaths,
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

const getProductEnWithAllColorsAndSizesByID = `-- name: GetProductEnWithAllColorsAndSizesByID :one
SELECT 
    p."ProductEnID",
    p."ProductNameEn",
    p."sCategoryIdEn",
    p."PriceEn",
    p."StockQuantity",
    p."ImagesPathEn",
    p."DescriptionEn",
    p."BrandEn",
    p."ManufacturedCountryEn",
    p."PenOutputEn",
    p."FeaturesEn",
    p."MaterialEn",
    p."StapleSizeEn",
    p."CapacityEn",
    p."WeightEn",
    p."ThicknessEn",
    p."PackagingEn",
    p."UsageEn",
    p."InstructionsEn",
    p."ProductCodeEn",
    p."CostPriceEn",
    p."RetailPriceEn",
    p."WarehouseStockEn",
    p."Created_At",
    p."Updated_At",
    COALESCE(
        ARRAY_AGG(DISTINCT c."ColorId") FILTER (WHERE c."ColorId" IS NOT NULL),
        '{}'
    ) AS "ColorIds",
    COALESCE(
        ARRAY_AGG(DISTINCT c."Color") FILTER (WHERE c."Color" IS NOT NULL),
        '{}'
    ) AS "ColorNames",
    COALESCE(
        ARRAY_AGG(DISTINCT s."SizeId") FILTER (WHERE s."SizeId" IS NOT NULL),
        '{}'
    ) AS "SizeIds",
    COALESCE(
        ARRAY_AGG(DISTINCT s."Size") FILTER (WHERE s."Size" IS NOT NULL),
        '{}'
    ) AS "SizeNames",
    COALESCE(
        ARRAY_AGG(DISTINCT pi."ImagePath") FILTER (WHERE pi."ImagePath" IS NOT NULL),
        '{}'
    ) AS "ImagePaths"
FROM "productEn" p
LEFT JOIN "productEn_colors" pc 
       ON p."ProductEnID" = pc."ProductEnID"
LEFT JOIN "productEn_sizes" ps 
       ON p."ProductEnID" = ps."ProductEnID"
LEFT JOIN "Color" c 
       ON pc."ColorId" = c."ColorId"
LEFT JOIN "Size" s 
       ON ps."SizeId" = s."SizeId"
LEFT JOIN "productImagesEn" pi 
       ON p."ProductEnID" = pi."ProductEnID"
WHERE p."ProductEnID" = $1
GROUP BY p."ProductEnID"
`

type GetProductEnWithAllColorsAndSizesByIDRow struct {
	ProductEnID           int32
	ProductNameEn         string
	SCategoryIdEn         int32
	PriceEn               string
	StockQuantity         int32
	ImagesPathEn          string
	DescriptionEn         string
	BrandEn               string
	ManufacturedCountryEn string
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
	CreatedAt             sql.NullTime
	UpdatedAt             sql.NullTime
	ColorIds              interface{}
	ColorNames            interface{}
	SizeIds               interface{}
	SizeNames             interface{}
	ImagePaths            interface{}
}

func (q *Queries) GetProductEnWithAllColorsAndSizesByID(ctx context.Context, productenid int32) (GetProductEnWithAllColorsAndSizesByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getProductEnWithAllColorsAndSizesByID, productenid)
	var i GetProductEnWithAllColorsAndSizesByIDRow
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
		&i.ColorIds,
		&i.ColorNames,
		&i.SizeIds,
		&i.SizeNames,
		&i.ImagePaths,
	)
	return i, err
}

const updateByEnImagePath = `-- name: UpdateByEnImagePath :one
UPDATE
    "productEn"
SET
    "ImagesPathEn" = $1
WHERE
    "ProductEnID" = $2 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
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
    "ProductEnID" = $2 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
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
    "ProductEnID" = $2 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
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
    "PenOutputEn" = $9,
    "FeaturesEn" = $10,
    "MaterialEn" = $11,
    "StapleSizeEn" = $12,
    "CapacityEn" = $13,
    "WeightEn" = $14,
    "ThicknessEn" = $15,
    "PackagingEn" = $16,
    "UsageEn" = $17,
    "InstructionsEn" = $18,
    "ProductCodeEn" = $19,
    "CostPriceEn" = $20,
    "RetailPriceEn" = $21,
    "WarehouseStockEn" = $22
    WHERE
    "ProductEnID" = $23 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
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
    "ProductEnID" = $2 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
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
    "ProductEnID" = $2 RETURNING "ProductEnID", "ProductNameEn", "sCategoryIdEn", "PriceEn", "StockQuantity", "ImagesPathEn", "DescriptionEn", "BrandEn", "ManufacturedCountryEn", "PenOutputEn", "FeaturesEn", "MaterialEn", "StapleSizeEn", "CapacityEn", "WeightEn", "ThicknessEn", "PackagingEn", "UsageEn", "InstructionsEn", "ProductCodeEn", "CostPriceEn", "RetailPriceEn", "WarehouseStockEn", "Created_At", "Updated_At"
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
