CREATE TABLE clients (
    "client_id" SERIAL NOT NULL, 
    "firstname" VARCHAR(50),
    "lastname" VARCHAR(50),
    "phoneNumber" VARCHAR(50),
    "address" VARCHAR(50),
    "type" VARCHAR(10) DEFAULT 'client',
    "username" VARCHAR(50),
    "password" VARCHAR(20),
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);
