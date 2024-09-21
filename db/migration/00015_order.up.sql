-- Order Items Table
BEGIN;

CREATE TABLE
    "OrderItems" (
        "OrderItemId" SERIAL PRIMARY KEY, -- Unique ID for each order item
        "CustomerOrderId" INT NOT NULL, -- Foreign key to the CustomerOrderDetail table
        "ProductMnID" INT NOT NULL DEFAULT 0, -- Foreign key to the Product table
        "ProductEnID" INT NOT NULL DEFAULT 0,
        "Quantity" INT NOT NULL, -- Quantity of the product ordered
        "PriceAtOrder" DECIMAL(10, 2) NOT NULL, -- Price of the product at the time of order
        FOREIGN KEY ("CustomerOrderId") REFERENCES "CustomerOrderDetail" ("CustomerOrderId") ON DELETE CASCADE,
        FOREIGN KEY ("ProductMnID") REFERENCES "productMn" ("ProductMnID") ON DELETE CASCADE, -- Link to product table
        FOREIGN KEY ("ProductEnID") REFERENCES "productEn" ("ProductEnID") ON DELETE CASCADE
    ) TABLESPACE pg_default;

COMMIT;