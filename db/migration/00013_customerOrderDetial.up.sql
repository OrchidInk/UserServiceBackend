-- Customer Order Details Table
BEGIN;

CREATE TABLE "CustomerOrderDetail" (
    "CustomerOrderId" SERIAL PRIMARY KEY,
    "CustomerId" INT NOT NULL,
    "OrderDate" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "OrderStatus" VARCHAR(50) DEFAULT 'pending',
    "TotalAmount" DECIMAL(10, 2) NOT NULL,
    "PaymentStatus" VARCHAR(50) DEFAULT 'unpaid',
    FOREIGN KEY ("CustomerId") REFERENCES "User" ("ID") ON DELETE CASCADE -- Link to User table (Customer data)
) TABLESPACE pg_default;

COMMIT;