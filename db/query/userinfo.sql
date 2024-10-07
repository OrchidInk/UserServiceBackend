-- name: CreateUserInfo :one
INSERT INTO
    "UserInfo" (
        "UserId",
        "UserImagePath",
        "LastName",
        "FirstName",
        "Email",
        "BirthDate",
        "PhoneNumber1",
        "PhoneNumber2",
        "Address1"
    )
VALUES
    (
        sqlc.arg ('UserId'),
        sqlc.arg ('UserImagePath'),
        sqlc.arg ('LastName'),
        sqlc.arg ('FirstName'),
        sqlc.arg ('Email'),
        sqlc.arg ('BirthDate'),
        sqlc.arg ('PhoneNumber1'),
        sqlc.arg ('PhoneNumber2'),
        sqlc.arg ('Address1')
    ) RETURNING *;

-- name: FindUserInfoByUserId :one
SELECT
    *
FROM
    "UserInfo"
WHERE
    "UserId" = sqlc.arg ('UserId')
LIMIT
    1;

-- name: FindUserInfoByEmail :one
SELECT
    *
FROM
    "UserInfo"
WHERE
    "Email" = sqlc.arg ('Email')
LIMIT
    1;

-- name: UpdateUserInfo :exec
UPDATE "UserInfo"
SET
    "LastName" = sqlc.arg ('LastName'),
    "FirstName" = sqlc.arg ('FirstName'),
    "Email" = sqlc.arg ('Email'),
    "BirthDate" = sqlc.arg ('BirthDate'),
    "PhoneNumber1" = sqlc.arg ('PhoneNumber1'),
    "PhoneNumber2" = sqlc.arg ('PhoneNumber2'),
    "Address1" = sqlc.arg ('Address1')
WHERE
    "UserId" = sqlc.arg ('UserId');

-- name: DeleteUserInfo :exec
DELETE FROM "UserInfo"
WHERE
    "UserId" = sqlc.arg ('UserId');

-- name: UpdateByUserImagePath :one
UPDATE "UserInfo"
SET
    "UserImagePath" = sqlc.arg ('UserImagePath')
WHERE
    "UserId" = sqlc.arg ('UserId') RETURNING *;