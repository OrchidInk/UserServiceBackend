// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: banner.sql

package db

import (
	"context"
)

const createBannerInfo = `-- name: CreateBannerInfo :one
INSERT INTO
    "BannerInfo" ("BannerImageUrl")
VALUES
    ($1) RETURNING "BannerId", "BannerImageUrl"
`

func (q *Queries) CreateBannerInfo(ctx context.Context, bannerimageurl string) (BannerInfo, error) {
	row := q.db.QueryRowContext(ctx, createBannerInfo, bannerimageurl)
	var i BannerInfo
	err := row.Scan(&i.BannerId, &i.BannerImageUrl)
	return i, err
}

const deleteBannerInfo = `-- name: DeleteBannerInfo :exec
DELETE FROM
    "BannerInfo"
WHERE
    "BannerId" = $1
`

func (q *Queries) DeleteBannerInfo(ctx context.Context, bannerid int32) error {
	_, err := q.db.ExecContext(ctx, deleteBannerInfo, bannerid)
	return err
}

const findByBannerId = `-- name: FindByBannerId :one
SELECT
    "BannerId", "BannerImageUrl"
FROM
    "BannerInfo"
WHERE
    "BannerId" = $1
LIMIT
    1
`

func (q *Queries) FindByBannerId(ctx context.Context, bannerid int32) (BannerInfo, error) {
	row := q.db.QueryRowContext(ctx, findByBannerId, bannerid)
	var i BannerInfo
	err := row.Scan(&i.BannerId, &i.BannerImageUrl)
	return i, err
}

const getAllBanners = `-- name: GetAllBanners :many
SELECT
    "BannerId", "BannerImageUrl"
FROM
    "BannerInfo"
`

func (q *Queries) GetAllBanners(ctx context.Context) ([]BannerInfo, error) {
	rows, err := q.db.QueryContext(ctx, getAllBanners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BannerInfo
	for rows.Next() {
		var i BannerInfo
		if err := rows.Scan(&i.BannerId, &i.BannerImageUrl); err != nil {
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

const updateBannerInfo = `-- name: UpdateBannerInfo :one
UPDATE
    "BannerInfo"
SET
    "BannerImageUrl" = $1
WHERE
    "BannerId" = $2 RETURNING "BannerId", "BannerImageUrl"
`

type UpdateBannerInfoParams struct {
	BannerImageUrl string
	BannerId       int32
}

func (q *Queries) UpdateBannerInfo(ctx context.Context, arg UpdateBannerInfoParams) (BannerInfo, error) {
	row := q.db.QueryRowContext(ctx, updateBannerInfo, arg.BannerImageUrl, arg.BannerId)
	var i BannerInfo
	err := row.Scan(&i.BannerId, &i.BannerImageUrl)
	return i, err
}
