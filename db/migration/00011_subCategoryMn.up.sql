-- TABLE "subCategoryMn"
BEGIN;

CREATE TABLE "subCategoryMn" (
    "SubCategoryIDMn" SERIAL PRIMARY KEY NOT NULL,
    "subCategoryNameMn" VARCHAR(100) NOT NULL,
    "CategoryMnID" INT NOT NULL,
    CONSTRAINT "FK_Product_CategoryMn" FOREIGN KEY ("CategoryMnID") REFERENCES "categoryMn" ("CategoryMnID") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;