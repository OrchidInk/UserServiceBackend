-- TABLE "productImages"
BEGIN;

CREATE TABLE "productImages" (
    "ImageID" SERIAL PRIMARY KEY,
     "ProductMnID" INT REFERENCES "productMn" ("ProductMnID") ON DELETE CASCADE,
   "ProductEnID" INT REFERENCES "productEn" ("ProductEnID") ON DELETE CASCADE,
    "ImageURLEn" TEXT NOT NULL,
    "ImageURLMn" TEXT NOT NULL
) TABLESPACE pg_default;

COMMIT;
