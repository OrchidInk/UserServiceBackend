-- name: CreateBannerInfo :one
INSERT INTO
    "BannerInfo" ("BannerImageUrl")
VALUES
    (sqlc.arg('BannerImageUrl')) RETURNING *;

-- name: GetAllBanners :many
SELECT
    *
FROM
    "BannerInfo";

-- name: UpdateBannerInfo :one
UPDATE
    "BannerInfo"
SET
    "BannerImageUrl" = sqlc.arg('BannerImageUrl')
WHERE
    "BannerId" = sqlc.arg('BannerId') RETURNING *;

-- name: DeleteBannerInfo :exec
DELETE FROM
    "BannerInfo"
WHERE
    "BannerId" = sqlc.arg('BannerId');

-- name: FindByBannerId :one
SELECT
    *
FROM
    "BannerInfo"
WHERE
    "BannerId" = sqlc.arg('BannerId')
LIMIT
    1;