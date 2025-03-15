BEGIN;

CREATE TABLE "productEn_colors" (
    "ProductEnID" INT NOT NULL,
    "ColorId" INT NOT NULL,
    PRIMARY KEY ("ProductEnID", "ColorId"),
    CONSTRAINT "FK_ProductEnID_in_productEn_colors"
        FOREIGN KEY ("ProductEnID")
        REFERENCES "productEn" ("ProductEnID")
        ON DELETE CASCADE,
    CONSTRAINT "FK_ColorId_in_productEn_colors"
        FOREIGN KEY ("ColorId")
        REFERENCES "Color" ("ColorId")
        ON DELETE CASCADE
);

COMMIT;
