-- name: CreateUser :one
INSERT INTO
    "User" (
        "LastName",
        "FirstName",
        "UserName",
        "Email",
        "IsHashedPassword",
        "IsAdmin",
        "IsUser",
        "IsSuperAdmin",
        "IsActive"
    )
VALUES
    (
        sqlc.arg('LastName') :: VARCHAR(100),
        sqlc.arg('FirstName') :: VARCHAR(100),
        sqlc.arg('UserName') :: VARCHAR(100),
        sqlc.arg('Email') :: VARCHAR(100),
        sqlc.arg('IsHashedPassword') :: TEXT,
        sqlc.arg('IsAdmin') :: BOOLEAN,
        sqlc.arg('IsUser') :: BOOLEAN,
        sqlc.arg('IsSuperAdmin') :: BOOLEAN,
        sqlc.arg('IsActive') :: BOOLEAN
    ) RETURNING *;

-- name: FindByAdminId :one
SELECT
    *
FROM
    "User"
WHERE
    "ID" = sqlc.arg('ID')
    AND "IsAdmin" = TRUE;

-- name: FindByUserID :one
SELECT
    *,
    "IsAdmin",
    "IsUser",
    "IsSuperAdmin"
FROM
    "User"
WHERE 
    "ID" = sqlc.arg('ID')
LIMIT 1;

-- name: FindByUser :one
SELECT
    *
FROM
    "User"
WHERE
    "UserName" = sqlc.arg('UserName') :: VARCHAR(100)
LIMIT 1;

-- name: FindByUserName :one
SELECT
    *,
    "IsAdmin",
    "IsSuperAdmin",
    "IsUser"
FROM
    "User"
WHERE
    "UserName" = sqlc.arg('UserName') :: VARCHAR(100)
LIMIT 1;

-- name: FindBySuperAdminAdmin :many
SELECT
    *
FROM
    "User"
WHERE
    "IsSuperAdmin" = TRUE;

-- name: GetListAdmin :many
SELECT
    *
FROM
    "User"
WHERE
    "IsAdmin" = TRUE;

-- name: GetListUser :many
SELECT  
    *
FROM
    "User"
WHERE
    "IsUser" = TRUE;


-- name: CountListUser :many
SELECT
    count(*)
FROM
    "User"
WHERE
    "IsUser" = TRUE;