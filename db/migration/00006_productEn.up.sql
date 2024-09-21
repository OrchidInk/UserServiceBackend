-- TABLE "productEn"
BEGIN;

CREATE TABLE
    "productEn" (
        "ProductEnID" SERIAL PRIMARY KEY, -- Unique identifier for each product in English
        "CategoryEnID" INT NOT NULL, -- Foreign key to the English category table
        "PriceEn" DECIMAL(10, 2) NOT NULL, -- Product price in English
        "StockQuantity" INT NOT NULL DEFAULT 0,
        "ImagesID" INT NOT NULL,
        "Created_At" TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            "Updated_At" TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            CONSTRAINT "FK_Product_CategoryEn" FOREIGN KEY ("CategoryEnID") REFERENCES "categoryEn" ("CategoryEnID") ON DELETE CASCADE,
            CONSTRAINT "FK_Product_ImagesID" FOREIGN KEY ("ImagesID") REFERENCES "productImages" ("ImageID") ON DELETE CASCADE
    ) TABLESPACE pg_default;

COMMIT;