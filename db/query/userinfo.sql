-- name: CreateUserInfo :one
INSERT INTO
    "UserInfo" (
        "UserId",
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
        sqlc.arg('UserId') :: INT,
        sqlc.arg('LastName') :: VARCHAR(100),
        sqlc.arg('FirstName') :: VARCHAR(100),
        sqlc.arg('Email') :: VARCHAR(100),
        sqlc.arg('BirthDate') :: DATE,
        sqlc.arg('PhoneNumber1') :: VARCHAR(12),
        sqlc.arg('PhoneNumber2') :: VARCHAR(12),
        sqlc.arg('Address1') :: VARCHAR(150)
    ) RETURNING *;

-- name: FindUserInfoByUserId :one
SELECT
    *
FROM
    "UserInfo"
WHERE
    "UserId" = sqlc.arg('UserId') :: INT
LIMIT 1;


-- name: FindUserInfoByEmail :one
SELECT
    *
FROM
    "UserInfo"
WHERE
    "Email" = sqlc.arg('Email') :: VARCHAR(100)
LIMIT 1;


-- name: UpdateUserInfo :exec
UPDATE
    "UserInfo"
SET
    "LastName" = sqlc.arg('LastName') :: VARCHAR(100),
    "FirstName" = sqlc.arg('FirstName') :: VARCHAR(100),
    "Email" = sqlc.arg('Email') :: VARCHAR(100),
    "BirthDate" = sqlc.arg('BirthDate') :: DATE,
    "PhoneNumber1" = sqlc.arg('PhoneNumber1') :: VARCHAR(12),
    "PhoneNumber2" = sqlc.arg('PhoneNumber2') :: VARCHAR(12),
    "Address1" = sqlc.arg('Address1') :: VARCHAR(150)
WHERE
    "UserId" = sqlc.arg('UserId') :: INT;

-- name: DeleteUserInfo :exec
DELETE FROM
    "UserInfo"
WHERE
    "UserId" = sqlc.arg('UserId') :: INT;
