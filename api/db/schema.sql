CREATE TABLE "auth_identities" (
  "id" text NOT NULL,
  "provider" text NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz,
  "user_id" text,
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
  "creator_id" text,
  "time" timestamptz NOT NULL,
  "value" float8 NOT NULL,
  "summary" text NOT NULL,
  "description" text,
  "payers" jsonb NOT NULL DEFAULT '[]'::jsonb,
  "participants" jsonb NOT NULL DEFAULT '[]'::jsonb
);

CREATE TABLE "transaction_logs" (
  "id" text PRIMARY KEY,
  "transaction_id" text NOT NULL,
  "actor_id" text NOT NULL,
  "action" text NOT NULL,
  "time" timestamptz NOT NULL,
  "meta" jsonb NOT NULL DEFAULT '{}'::jsonb
);

ALTER TABLE "auth_identities" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("id");

ALTER TABLE "transaction_logs" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "transaction_logs" ADD FOREIGN KEY ("actor_id") REFERENCES "users" ("id");

CREATE INDEX ON "auth_identities" ("user_id");
