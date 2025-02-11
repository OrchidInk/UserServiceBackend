BEGIN;

CREATE TABLE "subCategoryEn" (
    "SubCategoryIDEn" SERIAL PRIMARY KEY NOT NULL,
    "subCategoryNameEn" VARCHAR(100) NOT NULL,
    "CategoryEnID" INT NOT NULL,
    CONSTRAINT "FK_Product_CategoryEn" 
        FOREIGN KEY ("CategoryEnID") 
        REFERENCES "categoryEn" ("CategoryEnID") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;
