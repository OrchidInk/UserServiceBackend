-- TABLE "UserInfo"
BEGIN;

CREATE TABLE "UserInfo" (
    "UserInfoId" serial not NULL,
    "ID" int DEFAULT(0) not NULL
) TABLESPACE pg_default;

COMMIT;