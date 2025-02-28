-- Size Product Table
BEGIN;

CREATE TABLE "Size" (
    "SizeId" SERIAL PRIMARY KEY,
    "Size" text not NULL
) TABLESPACE pg_default;

COMMIT;
