package migrations

func init() {
	up := `
CREATE TABLE "transaction_logs" (
  "id" text PRIMARY KEY,
  "transaction_id" text NOT NULL,
  "actor_id" text NOT NULL,
  "action" text NOT NULL,
  "time" timestamptz NOT NULL,
  "meta" jsonb NOT NULL DEFAULT '{}'::jsonb
);

ALTER TABLE "transaction_logs" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "transaction_logs" ADD FOREIGN KEY ("actor_id") REFERENCES "users" ("id");
`
	down := `
DROP TABLE "transaction_logs";
`
	Collection.MustRegisterTx(wrap(up), wrap(down))
}
