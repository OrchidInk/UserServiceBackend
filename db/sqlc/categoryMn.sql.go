// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: categoryMn.sql

package db

import (
	"context"
	"database/sql"
)

const createCategoryMn = `-- name: CreateCategoryMn :one
INSERT INTO
    "categoryMn" ("CategoryNameMn")
VALUES
    (
        $1 :: VARCHAR(100)
    ) RETURNING "CategoryMnID", "CategoryNameMn"
`

func (q *Queries) CreateCategoryMn(ctx context.Context, categorynamemn string) (CategoryMn, error) {
	row := q.db.QueryRowContext(ctx, createCategoryMn, categorynamemn)
	var i CategoryMn
	err := row.Scan(&i.CategoryMnID, &i.CategoryNameMn)
	return i, err
}

const deleteCategoryByMnId = `-- name: DeleteCategoryByMnId :exec
DELETE FROM
    "categoryMn"
WHERE
    "CategoryMnID" = $1 :: INT
`

func (q *Queries) DeleteCategoryByMnId(ctx context.Context, categorymnid int32) error {
	_, err := q.db.ExecContext(ctx, deleteCategoryByMnId, categorymnid)
	return err
}

const findByCategoryMn = `-- name: FindByCategoryMn :one
SELECT
    "CategoryMnID", "CategoryNameMn"
FROM
    "categoryMn"
WHERE
    "CategoryMnID" = $1
LIMIT
    1
`

func (q *Queries) FindByCategoryMn(ctx context.Context, categorymnid int32) (CategoryMn, error) {
	row := q.db.QueryRowContext(ctx, findByCategoryMn, categorymnid)
	var i CategoryMn
	err := row.Scan(&i.CategoryMnID, &i.CategoryNameMn)
	return i, err
}

const findByCategoryMnId = `-- name: FindByCategoryMnId :one
SELECT
    "CategoryMnID", "CategoryNameMn"
FROM
    "categoryMn"
WHERE
    "CategoryMnID" = $1
LIMIT
    1
`

func (q *Queries) FindByCategoryMnId(ctx context.Context, categorymnid int32) (CategoryMn, error) {
	row := q.db.QueryRowContext(ctx, findByCategoryMnId, categorymnid)
	var i CategoryMn
	err := row.Scan(&i.CategoryMnID, &i.CategoryNameMn)
	return i, err
}

const findByNameMnCategoryMn = `-- name: FindByNameMnCategoryMn :one
SELECT
    "CategoryMnID", "CategoryNameMn"
FROM
    "categoryMn"
WHERE
    "CategoryNameMn" = $1 :: VARCHAR(100)
LIMIT
    1
`

func (q *Queries) FindByNameMnCategoryMn(ctx context.Context, categorynamemn string) (CategoryMn, error) {
	row := q.db.QueryRowContext(ctx, findByNameMnCategoryMn, categorynamemn)
	var i CategoryMn
	err := row.Scan(&i.CategoryMnID, &i.CategoryNameMn)
	return i, err
}

const findSubCategoriesAndProductsByCategoryIDMn = `-- name: FindSubCategoriesAndProductsByCategoryIDMn :many
SELECT
    c."CategoryMnID",
    c."CategoryNameMn",
    sc."SubCategoryIDMn",
    sc."subCategoryNameMn",
    scc."sCategoryIdMn",
    scc."sCategoryNameMn",
    p."ProductMnID",
    p."ProductNameMn",
    p."PriceMn",
    p."StockQuantity",
    p."ImagesPathMn"
FROM
    "categoryMn" c
    LEFT JOIN "subCategoryMn" sc ON c."CategoryMnID" = sc."CategoryMnID"
    LEFT JOIN "sCategoryMn" scc ON scc."SubCategoryIDMn" = sc."SubCategoryIDMn"
    LEFT JOIN "productMn" p ON p."sCategoryIdMn" = scc."sCategoryIdMn"
WHERE
    c."CategoryMnID" = $1
ORDER BY
    scc."sCategoryIdMn",
    p."ProductMnID"
`

type FindSubCategoriesAndProductsByCategoryIDMnRow struct {
	CategoryMnID      int32
	CategoryNameMn    string
	SubCategoryIDMn   sql.NullInt32
	SubCategoryNameMn sql.NullString
	SCategoryIdMn     sql.NullInt32
	SCategoryNameMn   sql.NullString
	ProductMnID       sql.NullInt32
	ProductNameMn     sql.NullString
	PriceMn           sql.NullString
	StockQuantity     sql.NullInt32
	ImagesPathMn      sql.NullString
}

func (q *Queries) FindSubCategoriesAndProductsByCategoryIDMn(ctx context.Context, categorymnid int32) ([]FindSubCategoriesAndProductsByCategoryIDMnRow, error) {
	rows, err := q.db.QueryContext(ctx, findSubCategoriesAndProductsByCategoryIDMn, categorymnid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindSubCategoriesAndProductsByCategoryIDMnRow
	for rows.Next() {
		var i FindSubCategoriesAndProductsByCategoryIDMnRow
		if err := rows.Scan(
			&i.CategoryMnID,
			&i.CategoryNameMn,
			&i.SubCategoryIDMn,
			&i.SubCategoryNameMn,
			&i.SCategoryIdMn,
			&i.SCategoryNameMn,
			&i.ProductMnID,
			&i.ProductNameMn,
			&i.PriceMn,
			&i.StockQuantity,
			&i.ImagesPathMn,
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

const getCategoriesWithSubCategoriesAndProductMn = `-- name: GetCategoriesWithSubCategoriesAndProductMn :many
SELECT
    c."CategoryMnID",
    c."CategoryNameMn",
    sc."SubCategoryIDMn",
    sc."subCategoryNameMn",
    scc."sCategoryIdMn",
    scc."sCategoryNameMn",
    p."ProductMnID",
    p."ProductNameMn",
    p."PriceMn",
    p."StockQuantity",
    p."ImagesPathMn"
FROM
    "categoryMn" c
    LEFT JOIN "subCategoryMn" sc ON c."CategoryMnID" = sc."CategoryMnID"
    LEFT JOIN "sCategoryMn" scc ON scc."SubCategoryIDMn" = sc."SubCategoryIDMn"
    LEFT JOIN "productMn" p ON scc."sCategoryIdMn" = p."sCategoryIdMn"
ORDER BY
    c."CategoryMnID",
    scc."sCategoryIdMn",
    p."ProductMnID"
`

type GetCategoriesWithSubCategoriesAndProductMnRow struct {
	CategoryMnID      int32
	CategoryNameMn    string
	SubCategoryIDMn   sql.NullInt32
	SubCategoryNameMn sql.NullString
	SCategoryIdMn     sql.NullInt32
	SCategoryNameMn   sql.NullString
	ProductMnID       sql.NullInt32
	ProductNameMn     sql.NullString
	PriceMn           sql.NullString
	StockQuantity     sql.NullInt32
	ImagesPathMn      sql.NullString
}

func (q *Queries) GetCategoriesWithSubCategoriesAndProductMn(ctx context.Context) ([]GetCategoriesWithSubCategoriesAndProductMnRow, error) {
	rows, err := q.db.QueryContext(ctx, getCategoriesWithSubCategoriesAndProductMn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCategoriesWithSubCategoriesAndProductMnRow
	for rows.Next() {
		var i GetCategoriesWithSubCategoriesAndProductMnRow
		if err := rows.Scan(
			&i.CategoryMnID,
			&i.CategoryNameMn,
			&i.SubCategoryIDMn,
			&i.SubCategoryNameMn,
			&i.SCategoryIdMn,
			&i.SCategoryNameMn,
			&i.ProductMnID,
			&i.ProductNameMn,
			&i.PriceMn,
			&i.StockQuantity,
			&i.ImagesPathMn,
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

const getCategoriesWithSubCategoriesMn = `-- name: GetCategoriesWithSubCategoriesMn :many
SELECT
    c."CategoryMnID",
    c."CategoryNameMn",
    sc."SubCategoryIDMn",
    sc."subCategoryNameMn",
    scc."sCategoryIdMn",
    scc."sCategoryNameMn"
FROM
    "categoryMn" c
    LEFT JOIN "subCategoryMn" sc 
        ON c."CategoryMnID" = sc."CategoryMnID"
    LEFT JOIN "sCategoryMn" scc 
        ON scc."SubCategoryIDMn" = sc."SubCategoryIDMn"
ORDER BY
    c."CategoryMnID"
`

type GetCategoriesWithSubCategoriesMnRow struct {
	CategoryMnID      int32
	CategoryNameMn    string
	SubCategoryIDMn   sql.NullInt32
	SubCategoryNameMn sql.NullString
	SCategoryIdMn     sql.NullInt32
	SCategoryNameMn   sql.NullString
}

func (q *Queries) GetCategoriesWithSubCategoriesMn(ctx context.Context) ([]GetCategoriesWithSubCategoriesMnRow, error) {
	rows, err := q.db.QueryContext(ctx, getCategoriesWithSubCategoriesMn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCategoriesWithSubCategoriesMnRow
	for rows.Next() {
		var i GetCategoriesWithSubCategoriesMnRow
		if err := rows.Scan(
			&i.CategoryMnID,
			&i.CategoryNameMn,
			&i.SubCategoryIDMn,
			&i.SubCategoryNameMn,
			&i.SCategoryIdMn,
			&i.SCategoryNameMn,
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

const getListByAllCategoryMn = `-- name: GetListByAllCategoryMn :many
SELECT
    "CategoryMnID", "CategoryNameMn"
FROM
    "categoryMn"
`

func (q *Queries) GetListByAllCategoryMn(ctx context.Context) ([]CategoryMn, error) {
	rows, err := q.db.QueryContext(ctx, getListByAllCategoryMn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CategoryMn
	for rows.Next() {
		var i CategoryMn
		if err := rows.Scan(&i.CategoryMnID, &i.CategoryNameMn); err != nil {
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

const updateCategoryMn = `-- name: UpdateCategoryMn :one
UPDATE
    "categoryMn"
SET
    "CategoryNameMn" = $1 :: VARCHAR(100)
WHERE
    "CategoryMnID" = $2 :: INT RETURNING "CategoryMnID", "CategoryNameMn"
`

type UpdateCategoryMnParams struct {
	CategoryNameMn string
	CategoryMnID   int32
}

func (q *Queries) UpdateCategoryMn(ctx context.Context, arg UpdateCategoryMnParams) (CategoryMn, error) {
	row := q.db.QueryRowContext(ctx, updateCategoryMn, arg.CategoryNameMn, arg.CategoryMnID)
	var i CategoryMn
	err := row.Scan(&i.CategoryMnID, &i.CategoryNameMn)
	return i, err
}
