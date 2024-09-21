-- TABLE "productMn"
BEGIN;

CREATE TABLE
    "productMn" (
        "ProductMnID" SERIAL PRIMARY KEY, -- Unique identifier for each product in Mongolian
        "CategoryMnID" INT NOT NULL, -- Foreign key to the Mongolian category table
        "PriceMn" DECIMAL(10, 2) NOT NULL, -- Product price in Mongolian
        "StockQuantity" INT NOT NULL DEFAULT 0, -- Quantity of stock available
        "Created_At" TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
            "Updated_At" TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Update timestamp
            CONSTRAINT "FK_Product_CategoryMn" FOREIGN KEY ("CategoryMnID") REFERENCES "categoryMn" ("CategoryMnID") ON DELETE CASCADE -- Correct foreign key reference
    ) TABLESPACE pg_default;

COMMIT;