package migrations

func init() {
	up := `
CREATE TABLE IF NOT EXISTS days_off (
	id text PRIMARY KEY,
	created_at timestamptz NOT NULL DEFAULT now(),
	updated_at timestamptz NOT NULL DEFAULT now(),
	user_id text NOT NULL REFERENCES users (id),
	date timestamptz NOT NULL,
	part text NOT NULL,
	note text NOT NULL DEFAULT ''
);
`
	down := `
DROP TABLE IF EXISTS days_off;
`
	Collection.MustRegisterTx(wrap(up), wrap(down))
}
