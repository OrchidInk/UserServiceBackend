-- TABLE "productTranslationsEn"
BEGIN;

CREATE TABLE "productTranslationsEn" (
    "ProductTranslationEnID" SERIAL PRIMARY KEY NOT NULL,
    "ProductEnID" INT REFERENCES "productEn" ("ProductEnID") ON DELETE CASCADE,
    "LanguageCode" VARCHAR(2) NOT NULL,
    "ProductEnName" VARCHAR(100) NOT NULL,
    "ProductDescription" TEXT,
    UNIQUE ("ProductEnId", "LanguageCode")
) TABLESPACE pg_default;

COMMIT;