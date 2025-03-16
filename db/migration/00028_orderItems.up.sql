BEGIN;

CREATE TABLE "OrderItems" (
    "OrderItemID" SERIAL PRIMARY KEY,
    "OrderID" INT NOT NULL,
    "ProductMnID" INT,
    "ProductEnID" INT,
    "Quantity" INT NOT NULL,
    "PriceAtOrder" DECIMAL(10, 2) NOT NULL,
    "SelectedColor" TEXT,
    "SelectedSize" TEXT,
    FOREIGN KEY ("OrderID") REFERENCES "Orders" ("OrderID") ON DELETE CASCADE
);

COMMIT;