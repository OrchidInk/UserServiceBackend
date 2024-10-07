-- TABLE "productImages"
BEGIN;

CREATE TABLE "productImages" (
  "ImageID" SERIAL PRIMARY KEY,
  "ImagePathEn" TEXT NOT NULL DEFAULT '',
  "ImagePathMn" TEXT NOT NULL DEFAULT ''
) TABLESPACE pg_default;

COMMIT;