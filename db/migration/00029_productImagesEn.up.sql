-- Create a separate table for multiple images of English products
BEGIN;

CREATE TABLE IF NOT EXISTS "productImagesEn" (
    "ImageID" SERIAL PRIMARY KEY,
    "ProductEnID" INT NOT NULL,
    "ImagePath" TEXT NOT NULL,
    "Created_At" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "FK_ProductImages_ProductEn" 
        FOREIGN KEY ("ProductEnID")
        REFERENCES "productEn" ("ProductEnID") ON DELETE CASCADE
);

COMMIT;