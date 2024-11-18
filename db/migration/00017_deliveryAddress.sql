BEGIN;

CREATE TABLE deliveryAddress (
    "AddressId" SERIAL PRIMARY KEY,
    "DeliverId" INT NOT NULL,
    "Street" VARCHAR(255) NOT NULL,
    "City" VARCHAR(100) NOT NULL,
    "State" VARCHAR(100),
    "PostalCode" VARCHAR(20) NOT NULL,
    "Created_At" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "FK_Delivery_ID" FOREIGN KEY ("DeliverId") REFERENCES "delivery"("DeliverId") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;