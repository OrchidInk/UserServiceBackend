-- Create a separate table for multiple images of Mongolian products
BEGIN;
CREATE TABLE IF NOT EXISTS "productImagesMn" (
    "ImageID" SERIAL PRIMARY KEY,
    "ProductMnID" INT NOT NULL,
    "ImagePath" TEXT NOT NULL,
    "Created_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "FK_ProductImages_ProductMn" 
        FOREIGN KEY ("ProductMnID")
        REFERENCES "productMn" ("ProductMnID") ON DELETE CASCADE
);

COMMIT;