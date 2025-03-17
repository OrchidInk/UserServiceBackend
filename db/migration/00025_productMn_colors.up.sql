BEGIN;

CREATE TABLE "productMn_colors" (
    "ProductMnID" INT NOT NULL,
    "ColorId" INT NOT NULL,
    PRIMARY KEY ("ProductMnID", "ColorId"),
    CONSTRAINT "FK_ProductMnID_in_productEn_colors"
        FOREIGN KEY ("ProductMnID")
        REFERENCES "productMn" ("ProductMnID")
        ON DELETE CASCADE,
    CONSTRAINT "FK_ColorId_in_productEn_colors"
        FOREIGN KEY ("ColorId")
        REFERENCES "Color" ("ColorId")
        ON DELETE CASCADE
);

COMMIT;
