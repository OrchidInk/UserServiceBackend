-- Color Table:
BEGIN;

CREATE TABLE "Color" (
"ColorId" SERIAL PRIMARY KEY,
"Color" TEXT not NULL
) TABLESPACE pg_default;

COMMIT;