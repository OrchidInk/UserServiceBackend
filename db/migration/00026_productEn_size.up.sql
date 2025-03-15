BEGIN;

CREATE TABLE "productEn_sizes" (
    "ProductEnID" INT NOT NULL,
    "SizeId" INT NOT NULL,
    PRIMARY KEY ("ProductEnID", "SizeId"),
    CONSTRAINT "FK_ProductEnID_in_productEn_sizes"
        FOREIGN KEY ("ProductEnID")
        REFERENCES "productEn" ("ProductEnID")
        ON DELETE CASCADE,
    CONSTRAINT "FK_SizeId_in_productEn_sizes"
        FOREIGN KEY ("SizeId")
        REFERENCES "Size" ("SizeId")
        ON DELETE CASCADE
);

COMMIT;
