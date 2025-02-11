-- File: 002_create_sCategoryEn.sql
BEGIN;

CREATE TABLE "sCategoryEn" (
    "sCategoryIdEn" SERIAL PRIMARY KEY NOT NULL,
    "sCategoryNameEn" VARCHAR(100) NOT NULL,
    "SubCategoryIDEn" INT NOT NULL,
    CONSTRAINT "Fk_sCategory" 
        FOREIGN KEY ("SubCategoryIDEn")
        REFERENCES "subCategoryEn" ("SubCategoryIDEn") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;
