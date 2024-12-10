-- TABLE "detailMn"
BEGIN;

CREATE TABLE "detailMn" (
    "detailMnId" SERIAL PRIMARY KEY,
    "ProductMnID" INT NOT NULL,
    "ChoiceName" VARCHAR(100) NOT NULL,
    "ChoiceValue" VARCHAR(100) NOT NULL,
    FOREIGN KEY ("ProductMnID") REFERENCES "productMn" ("ProductMnID") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;