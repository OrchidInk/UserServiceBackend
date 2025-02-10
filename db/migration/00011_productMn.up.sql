-- TABLE "productMn"
BEGIN;

CREATE TABLE "productMn" (
    "ProductMnID" SERIAL PRIMARY KEY,
    "ProductNameMn" TEXT NOT NULL,
    "sCategoryIdMn" INT NOT NULL,
    "PriceMn" DECIMAL(10, 2) NOT NULL,
    "StockQuantity" INT NOT NULL DEFAULT 0,
    "ImagesPathMn" TEXT NOT NULL DEFAULT '',
    "DescriptionMn" TEXT NOT NULL DEFAULT '',
    "BrandMn" TEXT NOT NULL DEFAULT '',
    "ManufacturedCountryMn" TEXT NOT NULL DEFAULT '',
    "ColorMn" TEXT NOT NULL DEFAULT '',
    "SizeMn" TEXT NOT NULL DEFAULT '',
    "PenOutputMn" TEXT NOT NULL DEFAULT '',
    "FeaturesMn" TEXT NOT NULL DEFAULT '',
    "MaterialMn" TEXT NOT NULL DEFAULT '',
    "StapleSizeMn" TEXT NOT NULL DEFAULT '',
    "CapacityMn" TEXT NOT NULL DEFAULT '',
    "WeightMn" TEXT NOT NULL DEFAULT '',
    "ThicknessMn" TEXT NOT NULL DEFAULT '',
    "PackagingMn" TEXT NOT NULL DEFAULT '',
    "UsageMn" TEXT NOT NULL DEFAULT '',
    "InstructionsMn" TEXT NOT NULL DEFAULT '',
    "ProductCodeMn" TEXT NOT NULL DEFAULT '',
    "CostPriceMn" DECIMAL(10, 2) NOT NULL,
    "RetailPriceMn" DECIMAL(10, 2) NOT NULL,
    "WarehouseStockMn" INT NOT NULL DEFAULT 0,
    "Created_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "Updated_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "FK_Product_sCategoryIDMn" FOREIGN KEY ("sCategoryIdMn") REFERENCES "sCategoryMn" ("sCategoryIdMn") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;