-- name: CreateBannerInfo :one
INSERT INTO
    "BannerInfo" ("BannerImagePath", "BannerImageUrl")
VALUES
    (
        sqlc.arg('BannerImagePath'),
        sqlc.arg('BannerImageUrl')
    ) RETURNING *;

-- name: GetAllBanners :many
SELECT
    *
FROM
    "BannerInfo";

-- name: UpdateBannerInfo :one
UPDATE
    "BannerInfo"
SET
    "BannerImagePath" = sqlc.arg('BannerImagePath'),
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