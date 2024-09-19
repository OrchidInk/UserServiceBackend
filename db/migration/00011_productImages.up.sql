-- TABLE: "productImages"
BEGIN;

CREATE TABLE "productImages" (
    "ImageID" SERIAL PRIMARY KEY,
    "ProductID" INT REFERENCES "product" ("ProductID") ON DELETE CASCADE,
    "ImageURL" TEXT NOT NULL
) TABLESPACE pg_default;

COMMIT;
