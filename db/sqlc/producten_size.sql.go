// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: producten_size.sql

package db

import (
	"context"
)

const insertProductEnSize = `-- name: InsertProductEnSize :one
INSERT INTO "productEn_sizes" ("ProductEnID", "SizeId")
VALUES (
    $1, 
    $2) RETURNING "ProductEnID", "SizeId"
`

type InsertProductEnSizeParams struct {
	ProductEnID int32
	SizeId      int32
}

func (q *Queries) InsertProductEnSize(ctx context.Context, arg InsertProductEnSizeParams) (ProductEnSize, error) {
	row := q.db.QueryRowContext(ctx, insertProductEnSize, arg.ProductEnID, arg.SizeId)
	var i ProductEnSize
	err := row.Scan(&i.ProductEnID, &i.SizeId)
	return i, err
}
