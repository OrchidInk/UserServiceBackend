-- TABLE "categoryEn"
BEGIN;

CREATE TABLE
    "categoryEn" (
        "CategoryEnID" SERIAL PRIMARY KEY, -- Unique identifier for each category in English
        "CategoryNameEn" VARCHAR(100) NOT NULL -- Category name in English
    ) TABLESPACE pg_default;

COMMIT;