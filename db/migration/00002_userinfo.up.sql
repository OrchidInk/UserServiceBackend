-- TABLE "UserInfo"
BEGIN;

CREATE TABLE "UserInfo" (
    "UserInfoId" SERIAL NOT NULL,
    "UserId" int DEFAULT(0) NOT NULL,
    "LastName" VARCHAR(100) NOT NULL,
    "FirstName" VARCHAR(100) NOT NULL,
    "Email" VARCHAR(100) NOT NULL,
    "BirthDate" DATE NOT NULL,
    "PhoneNumber1" VARCHAR(12) NOT NULL,
    "PhoneNumber2" VARCHAR(12) NOT NULL,
    "Address1" VARCHAR(150) NOT NULL,
    CONSTRAINT "FK_UserId" FOREIGN KEY ("UserId") REFERENCES "User" ("ID"),
    CONSTRAINT "FK_Email" FOREIGN KEY ("Email") REFERENCES "User" ("Email")
) TABLESPACE pg_default;

COMMIT;