-- Order Items Table
BEGIN;

CREATE TABLE "OrderItems" (
    "OrderItemId" SERIAL PRIMARY KEY,
    "CustomerOrderId" INT ,
    "ProductMnID" INT ,
    "ProductEnID" INT ,
    "UserId" INT NOT NULL DEFAULT 0,
    "PhoneNumber" varchar(100) not NULL,
    "Quantity" INT NOT NULL,
    "PriceAtOrder" DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY ("ProductMnID") REFERENCES "productMn" ("ProductMnID") ON DELETE CASCADE,
    FOREIGN KEY ("ProductEnID") REFERENCES "productEn" ("ProductEnID") ON DELETE CASCADE,
    FOREIGN KEY ("UserId") REFERENCES "User" ("ID") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;