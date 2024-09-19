-- Customer Order Details Table
BEGIN;

CREATE TABLE "CustomerOrderDetail" (
    "CustomerOrderId" SERIAL PRIMARY KEY,            -- Unique order ID for each customer order
    "CustomerId" INT NOT NULL,                       -- Foreign key to the Customer table (you would need a customer table)
    "OrderDate" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Timestamp of the order creation
    "OrderStatus" VARCHAR(50) DEFAULT 'pending',     -- Status of the order (e.g., pending, shipped, delivered)
    "TotalAmount" DECIMAL(10, 2) NOT NULL,           -- Total amount of the order
    "PaymentStatus" VARCHAR(50) DEFAULT 'unpaid',    -- Payment status (e.g., unpaid, paid)
    FOREIGN KEY ("CustomerId") REFERENCES "User"("ID") ON DELETE CASCADE -- Link to User table (Customer data)
) TABLESPACE pg_default;

COMMIT;
