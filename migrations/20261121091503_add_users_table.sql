-- Create enum type "status"
CREATE TYPE "status" AS ENUM ('active', 'inactive', 'pending');
-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "username" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "phone" character varying(255) NOT NULL,
  "first_name" character varying(255) NOT NULL,
  "last_name" character varying(255) NOT NULL,
  "password" text NOT NULL,
  "status" "status" NOT NULL DEFAULT 'active',
  "last_login_at" bigint NULL,
  "created_at" bigint NOT NULL DEFAULT (EXTRACT(epoch FROM now()))::bigint,
  "updated_at" bigint NOT NULL DEFAULT (EXTRACT(epoch FROM now()))::bigint,
  "deleted_at" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "users_email_key" UNIQUE ("email"),
  CONSTRAINT "users_phone_key" UNIQUE ("phone"),
  CONSTRAINT "users_username_key" UNIQUE ("username")
);
