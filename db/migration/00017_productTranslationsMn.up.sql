-- TABLE "productTranslationsMn"
BEGIN;

CREATE TABLE
    "productTranslationsMn" (
        "ProductTranslationMnID" SERIAL PRIMARY KEY NOT NULL,
        "ProductMnID" INT REFERENCES "productMn" ("ProductMnID") ON DELETE CASCADE,
        "LanguageCode" VARCHAR(2) NOT NULL,
        "ProductMnName" VARCHAR(100) NOT NULL,
        "ProductDescription" TEXT,
        UNIQUE ("ProductMnID", "LanguageCode")
    ) TABLESPACE pg_default;

COMMIT;