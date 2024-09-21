BEGIN;

CREATE TABLE
    "adminPermission" (
        "PermissionID" SERIAL PRIMARY KEY, -- Unique identifier for the permission
        "AdminID" INT REFERENCES "User" ("ID") ON DELETE CASCADE, -- Foreign key to admin
        "CanCreate" BOOLEAN DEFAULT FALSE, -- Can create resources
        "CanRead" BOOLEAN DEFAULT TRUE, -- Can read resources (default to true)
        "CanUpdate" BOOLEAN DEFAULT FALSE, -- Can update resources
        "CanDelete" BOOLEAN DEFAULT FALSE, -- Can delete resources
        "AssignedAt" TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- When the permission was assigned
            CONSTRAINT "FK_ADMINID" FOREIGN KEY ("AdminID") REFERENCES "User" ("ID")
    ) TABLESPACE pg_default;

COMMIT;