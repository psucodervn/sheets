package migrations

func init() {
	up := `
ALTER TABLE "users" ADD COLUMN telegram_id text UNIQUE;

CREATE TABLE "telegram_tokens" (
	"token" text PRIMARY KEY,
	"created_at" timestamptz NOT NULL,
	"expire_at" timestamptz NOT NULL,
	"user_id" text NOT NULL REFERENCES "users" ("id")
)
`
	down := `
DROP TABLE "telegram_tokens";

ALTER TABLE "users" DROP COLUMN telegram_id;
`
	Collection.MustRegisterTx(wrap(up), wrap(down))
}
