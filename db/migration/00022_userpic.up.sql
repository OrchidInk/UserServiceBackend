--TABLE: "UserPicture"
BEGIN;

CREATE TABLE "UserPicture" (
    "UserPicId" SERIAL NOT NULL PRIMARY KEY,
    "UserImagePath" TEXT NOT NULL,
    "UserImageUrl" TEXT NOT NULL
) TABLESPACE pg_default;

COMMIT;