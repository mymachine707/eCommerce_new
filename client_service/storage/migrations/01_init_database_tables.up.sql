CREATE TABLE "clients" (
	"id" CHAR(36) PRIMARY KEY,
	"firstname" VARCHAR(255) UNIQUE NOT NULL,
	"lastname" TEXT NOT NULL,
	author_id CHAR(36),
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);
