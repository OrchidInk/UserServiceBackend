-- TABLE "categoryMn"
BEGIN;

CREATE TABLE "categoryMn" (
    "CategoryMnID" SERIAL PRIMARY KEY,
    "CategoryNameMn" VARCHAR(100) NOT NULL
) TABLESPACE pg_default;

COMMIT;