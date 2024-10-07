-- TABLE "categoryEn"
BEGIN;

CREATE TABLE "categoryEn" (
    "CategoryEnID" SERIAL PRIMARY KEY,
    "CategoryNameEn" VARCHAR(100) NOT NULL
) TABLESPACE pg_default;

COMMIT;