-- TABLE "productReviews"
BEGIN;

CREATE TABLE "productReviews" (
    "ReviewID" SERIAL PRIMARY KEY,
    "ProductMnID" INT REFERENCES "productMn" ("ProductMnID") ON DELETE CASCADE,
    "ProductEnID" INT REFERENCES "productEn" ("ProductEnID") ON DELETE CASCADE,
    "UserID" INT REFERENCES "User" ("ID"),
    "Rating" INT CHECK (
        "Rating" >= 1
        AND "Rating" <= 5
    ),
    "ReviewText" TEXT,
    "Created_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
) TABLESPACE pg_default;

COMMIT;