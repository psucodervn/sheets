CREATE TABLE "auth_identities" (
  "id" text NOT NULL,
  "provider" text NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz,
  "user_id" text NOT NULL,
  PRIMARY KEY ("id", "provider")
);

CREATE TABLE "users" (
  "id" text PRIMARY KEY,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz,
  "name" text NOT NULL,
  "email" text UNIQUE,
  "hashed_password" text,
  "sheet_name" text UNIQUE,
  "jira_name" text UNIQUE
);

CREATE TABLE "transactions" (
  "id" text PRIMARY KEY,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz,
  "time" timestamptz NOT NULL,
  "value" float8 NOT NULL,
  "summary" text NOT NULL,
  "description" text,
  "senders" jsonb NOT NULL DEFAULT '[]',
  "receivers" jsonb NOT NULL DEFAULT '[]'
);

ALTER TABLE "auth_identities" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "auth_identities" ("user_id");
