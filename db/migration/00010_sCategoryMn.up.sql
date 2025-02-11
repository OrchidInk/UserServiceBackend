BEGIN;

CREATE TABLE "sCategoryMn" (
    "sCategoryIdMn" SERIAL PRIMARY KEY,
    "sCategoryName" VARCHAR(100) NOT NULL,
    "SubCategoryIDMn" INT NOT NULL,
    CONSTRAINT "FK_sCAtegory_ID" FOREIGN KEY ("SubCategoryIDMn")
        REFERENCES "subCategoryMn" ("SubCategoryIDMn") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;
