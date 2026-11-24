-- Create "roles" table
CREATE TABLE "roles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "created_at" bigint NOT NULL DEFAULT (EXTRACT(epoch FROM now()))::bigint,
  "updated_at" bigint NOT NULL DEFAULT (EXTRACT(epoch FROM now()))::bigint,
  "deleted_at" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "roles_name_key" UNIQUE ("name")
);
