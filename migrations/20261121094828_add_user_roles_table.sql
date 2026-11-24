-- Create "user_roles" table
CREATE TABLE "user_roles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "role_id" uuid NOT NULL,
  "assigned_by" uuid NOT NULL,
  "assigned_at" bigint NOT NULL DEFAULT (EXTRACT(epoch FROM now()))::bigint,
  "created_at" bigint NOT NULL DEFAULT (EXTRACT(epoch FROM now()))::bigint,
  "updated_at" bigint NOT NULL DEFAULT (EXTRACT(epoch FROM now()))::bigint,
  "deleted_at" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_roles_assigned_by_fkey" FOREIGN KEY ("assigned_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "user_roles_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "user_roles_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_user_roles_role_id" to table: "user_roles"
CREATE INDEX "idx_user_roles_role_id" ON "user_roles" ("role_id");
-- Create index "idx_user_roles_user_id" to table: "user_roles"
CREATE INDEX "idx_user_roles_user_id" ON "user_roles" ("user_id");
