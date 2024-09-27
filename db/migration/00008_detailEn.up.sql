-- TABLE "detailEn"
BEGIN;

CREATE TABLE "detailEn" (
    "detailEnId" SERIAL PRIMARY KEY,
    "ProductEnID" INT NOT NULL,
    "ChoiceName" VARCHAR(100) NOT NULL,
    "ChoiceValue" VARCHAR(100) NOT NULL,
    FOREIGN KEY ("ProductEnID") REFERENCES "productEn" ("ProductEnID") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;