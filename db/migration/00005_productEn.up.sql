-- TABLE "productEn"
BEGIN;

CREATE TABLE "productEn" (
    "ProductEnID" SERIAL PRIMARY KEY,                  -- Unique identifier for each product in English
    "CategoryEnID" INT NOT NULL,                       -- Foreign key to the English category table
    "PriceEn" DECIMAL(10, 2) NOT NULL,                 -- Product price in English
    "StockQuantity" INT NOT NULL DEFAULT 0,            -- Quantity of stock available
    "Created_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,  -- Creation timestamp
    "Updated_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,  -- Update timestamp
    CONSTRAINT "FK_Product_CategoryEn" FOREIGN KEY ("CategoryEnID") REFERENCES "categoryEn"("CategoryEnID") ON DELETE CASCADE -- Foreign key reference to the categoryEn table
) TABLESPACE pg_default;

COMMIT;
