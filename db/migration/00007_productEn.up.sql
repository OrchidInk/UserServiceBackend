-- TABLE "productEn"
BEGIN;

CREATE TABLE "productEn" (
    "ProductEnID" SERIAL PRIMARY KEY,
    "ProductNameEn" TEXT NOT NULL,
    "subCategoryIDEn" INT NOT NULL,
    "PriceEn" DECIMAL(10, 2) NOT NULL,
    "StockQuantity" INT NOT NULL DEFAULT 0,
    "ImagesPathEn" TEXT NOT NULL DEFAULT '',
    "Created_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "Updated_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "FK_Product_subCategoryIDEn" FOREIGN KEY ("subCategoryIDEn") REFERENCES "subCategoryEn" ("subCategoryIDEn") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;