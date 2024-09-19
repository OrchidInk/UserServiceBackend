-- TABLE: "delivery"
BEGIN;

CREATE TABLE "delivery" (
    "DeliverId" SERIAL PRIMARY KEY NOT NULL,
    "DeliverName" VARCHAR(100) NOT NULL,
    "OrderId" INT NOT NULL
) TABLESPACE pg_default;

COMMIT;