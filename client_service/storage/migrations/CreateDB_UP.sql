CREATE TABLE clients (
    "client_id" UUID DEFAULT UUID_GENERATE() PRIMARY KEY,
    "firstname" VARCHAR(50),
    "lastname" VARCHAR(50),
    "username" VARCHAR(50),
    "phoneNumber" VARCHAR(50),
    "address" VARCHAR(50),
    "type" DEFAULT 'client',
    "password" VARCHAR(20),
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
)