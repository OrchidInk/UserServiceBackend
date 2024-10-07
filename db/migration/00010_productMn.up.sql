-- TABLE "productMn"
BEGIN;

CREATE TABLE "productMn" (
    "ProductMnID" SERIAL PRIMARY KEY,
    "ProductNameMn" TEXT NOT NULL,
    "subCategoryIDMn" INT NOT NULL,
    "PriceMn" DECIMAL(10, 2) NOT NULL,
    "StockQuantity" INT NOT NULL DEFAULT 0,
    "ImagesPathMn" TEXT NOT NULL DEFAULT '',
    "Created_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "Updated_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "FK_Product_subCategoryIDMn" FOREIGN KEY ("subCategoryIDMn") REFERENCES "subCategoryMn" ("subCategoryIDMn") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;