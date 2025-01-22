-- TABLE "productEn"
BEGIN;

CREATE TABLE "productEn" (
    "ProductEnID" SERIAL PRIMARY KEY,
    "ProductNameEn" TEXT NOT NULL,
    "SubCategoryIDEn" INT NOT NULL,
    "PriceEn" DECIMAL(10, 2) NOT NULL,
    "StockQuantity" INT NOT NULL DEFAULT 0,
    "ImagesPathEn" TEXT NOT NULL DEFAULT '',
    "DescriptionEn" TEXT NOT NULL DEFAULT '',
    "BrandEn" TEXT NOT NULL DEFAULT '',
    "ManufacturedCountryEn" TEXT NOT NULL DEFAULT '',
    "ColorEn" TEXT NOT NULL DEFAULT '',
    "SizeEn" TEXT NOT NULL DEFAULT '',
    "PenOutputEn" TEXT NOT NULL DEFAULT '',
    "FeaturesEn" TEXT NOT NULL DEFAULT '',
    "MaterialEn" TEXT NOT NULL DEFAULT '',
    "StapleSizeEn" TEXT NOT NULL DEFAULT '',
    "CapacityEn" TEXT NOT NULL DEFAULT '',
    "WeightEn" TEXT NOT NULL DEFAULT '',
    "ThicknessEn" TEXT NOT NULL DEFAULT '',
    "PackagingEn" TEXT NOT NULL DEFAULT '',
    "UsageEn" TEXT NOT NULL DEFAULT '',
    "InstructionsEn" TEXT NOT NULL DEFAULT '',
    "ProductCodeEn" TEXT NOT NULL DEFAULT '',
    "CostPriceEn" DECIMAL(10, 2) NOT NULL,
    "RetailPriceEn" DECIMAL(10, 2) NOT NULL,
    "WarehouseStockEn" INT NOT NULL DEFAULT 0,
    "Created_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "Updated_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "FK_Product_SubCategoryIDEn" FOREIGN KEY ("SubCategoryIDEn") REFERENCES "subCategoryEn" ("SubCategoryIDEn") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;