-- TABLE "Customer"
BEGIN;

CREATE TABLE
    "Customer" (
        "CustomerId" serial PRIMARY KEY,
        "CustomerName" VARCHAR(100) NOT NULL,
        "ContractDate" INT NOT NULL,
        "IsActive" BOOLEAN DEFAULT FALSE NOT NULL,
        "Created_At" TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
            "Updated_At" TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
    ) TABLESPACE pg_default;

COMMIT;