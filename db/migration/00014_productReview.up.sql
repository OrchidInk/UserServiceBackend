-- TABLE "productReviews"
BEGIN;

CREATE TABLE
    "productReviews" (
        "ReviewID" SERIAL PRIMARY KEY, -- Unique identifier for the review
        "ProductMnID" INT REFERENCES "productMn" ("ProductMnID") ON DELETE CASCADE, -- Foreign key to Mongolian product
        "ProductEnID" INT REFERENCES "productEn" ("ProductEnID") ON DELETE CASCADE, -- Foreign key to English product
        "UserID" INT REFERENCES "User" ("ID"), -- Foreign key to user
        "Rating" INT CHECK (
            "Rating" >= 1
            AND "Rating" <= 5
        ), -- Rating (1-5 scale)
        "ReviewText" TEXT, -- Review text
        "Created_At" TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP -- Timestamp for review creation
    ) TABLESPACE pg_default;

COMMIT;