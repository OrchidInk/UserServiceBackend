BEGIN;

CREATE TABLE "productMn_sizes" (
    "ProductMnID" INT NOT NULL,
    "SizeId" INT NOT NULL,
    PRIMARY KEY ("ProductMnID", "SizeId"),
    CONSTRAINT "FK_ProductMnID_in_productEn_sizes"
        FOREIGN KEY ("ProductMnID")
        REFERENCES "productMn" ("ProductMnID")
        ON DELETE CASCADE,
    CONSTRAINT "FK_SizeId_in_productEn_sizes"
        FOREIGN KEY ("SizeId")
        REFERENCES "Size" ("SizeId")
        ON DELETE CASCADE
);

COMMIT;
