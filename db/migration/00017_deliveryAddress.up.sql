-- TABLE: "deliveryAddress"
BEGIN;

CREATE TABLE "deliveryAddress" (
    "AddressId" SERIAL PRIMARY KEY NOT NULL,
    "DeliverId" INT NOT NULL REFERENCES "delivery"("DeliverId") ON DELETE CASCADE,
    "Street" VARCHAR(255) NOT NULL,
    "City" VARCHAR(100) NOT NULL,
    "State" VARCHAR(100),
    "PostalCode" VARCHAR(20) NOT NULL,
    "CreatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) TABLESPACE pg_default;

COMMIT;