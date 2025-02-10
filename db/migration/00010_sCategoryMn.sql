-- TABLE "sCategoryMn"

BEGIN;

CREATE TABLE "sCategoryMn" (
    "sCategoryIdMn" SERIAL NOT NULL,
    "sCategoryName" VARCHAR(100) NOT NULL,
    "SubCategoryIDMn" INT NOT NULL,
    CONSTRAINT "FK_sCAtegory_ID" FOREIGN KEY ("SubCategoryIDMn") REFERENCES "subCategoryMn" ("SubCategoryIDMn") on DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;