-- TABLE "productReviews"
BEGIN;

CREATE TABLE "productReviews" (
    "ReviewID" SERIAL PRIMARY KEY,
    "ProductMnID" INT NOT NULL DEFAULT 0,
    "ProductEnID" INT NOT NULL DEFAULT 0,
    "UserID" INT REFERENCES "User" ("ID"),
    "Rating" INT CHECK (
        "Rating" >= 1
        AND "Rating" <= 5
    ),
    "ReviewText" TEXT,
    "Created_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "FK_ProductMN" FOREIGN KEY ("ProductMnID") REFERENCES "productMn" ("ProductMnID") ON DELETE CASCADE,
    CONSTRAINT "Fk_ProductEn" FOREIGN KEY ("ProductEnID") REFERENCES "productEn" ("ProductEnID") ON DELETE CASCADE
) TABLESPACE pg_default;

COMMIT;