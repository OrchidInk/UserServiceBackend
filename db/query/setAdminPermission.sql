-- name: SetPermissions :exec
INSERT INTO "adminPermission" (
    "AdminID",
    "CanCreate",
    "CanRead",
    "CanUpdate",
    "CanDelete"
) VALUES (
    sqlc.arg('AdminID') :: INT,
    sqlc.arg('CanCreate') :: BOOLEAN,
    sqlc.arg('CanRead') :: BOOLEAN,
    sqlc.arg('CanUpdate') :: BOOLEAN,
    sqlc.arg('CanDelete') :: BOOLEAN
) ON CONFLICT ("AdminID") DO UPDATE SET
    "CanCreate" = EXCLUDED."CanCreate",
    "CanRead" = EXCLUDED."CanRead",
    "CanUpdate" = EXCLUDED."CanUpdate",
    "CanDelete" = EXCLUDED."CanDelete";

-- name: GetPermissionsByAdminID :one
SELECT
    "CanCreate",
    "CanRead",
    "CanUpdate",
    "CanDelete"
FROM
    "adminPermission"
WHERE
    "AdminID" = sqlc.arg('AdminID') :: INT
LIMIT 1;
