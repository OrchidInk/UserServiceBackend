--TABLE "User"
BEGIN;

CREATE TABLE
    "User" (
        "ID" serial PRIMARY KEY,
        "LastName" VARCHAR(100) NOT NULL,
        "FirstName" VARCHAR(100) NOT NULL,
        "UserName" VARCHAR(100) NOT NULL UNIQUE,
        "Email" VARCHAR(100) NOT NULL UNIQUE,
        "IsHashedPassword" TEXT NOT NULL,
        "IsAdmin" BOOLEAN DEFAULT FALSE NOT NULL,
        "IsUser" BOOLEAN DEFAULT FALSE not NULL,
        "IsSuperAdmin" BOOLEAN DEFAULT FALSE NOT NULL,
        "IsActive" BOOLEAN DEFAULT TRUE NOT NULL,
        "Created_At" TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
    ) TABLESPACE pg_default;

COMMIT;