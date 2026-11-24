-- Create "role_permissions" table
CREATE TABLE "role_permissions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "role_id" uuid NOT NULL,
  "permission_id" uuid NOT NULL,
  "assigned_by" uuid NOT NULL,
  "assigned_at" bigint NOT NULL DEFAULT (EXTRACT(epoch FROM now()))::bigint,
  "created_at" bigint NOT NULL DEFAULT (EXTRACT(epoch FROM now()))::bigint,
  "updated_at" bigint NOT NULL DEFAULT (EXTRACT(epoch FROM now()))::bigint,
  "deleted_at" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "role_permissions_assigned_by_fkey" FOREIGN KEY ("assigned_by") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "role_permissions_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_role_permissions_permission_id" to table: "role_permissions"
CREATE INDEX "idx_role_permissions_permission_id" ON "role_permissions" ("permission_id");
-- Create index "idx_role_permissions_role_id" to table: "role_permissions"
CREATE INDEX "idx_role_permissions_role_id" ON "role_permissions" ("role_id");
