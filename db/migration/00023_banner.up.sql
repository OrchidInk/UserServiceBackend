--TABLE "BannerInfo"
BEGIN;

CREATE TABLE "BannerInfo" (
    "BannerId" SERIAL NOT NULL PRIMARY KEY,
    "BannerImageUrl" TEXT NOT NULL
) TABLESPACE pg_default;

COMMIT;