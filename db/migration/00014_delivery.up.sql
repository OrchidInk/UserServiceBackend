-- TABLE: "delivery"
BEGIN;

CREATE TABLE "delivery" (
    "DeliverId" SERIAL PRIMARY KEY NOT NULL,
    "DeliverName" VARCHAR(100) NOT NULL,
    "OrderId" INT NOT NULL,
    "DeliveryAmount" DECIMAL(10, 2) NOT NULL,
    "CreatedAt" timestamp DEFAULT CURRENT_TIMESTAMP
) TABLESPACE pg_default;

COMMIT;