-- TABLE "categoryMn"
BEGIN;

CREATE TABLE
    "categoryMn" (
        "CategoryMnID" SERIAL PRIMARY KEY, -- Unique identifier for each category in Mongolian
        "CategoryNameMn" VARCHAR(100) NOT NULL -- Mongolian name of the category
    ) TABLESPACE pg_default;

COMMIT;