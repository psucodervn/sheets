package migrations

func init() {
	up := `
CREATE TABLE checkins (
	time timestamptz NOT NULL DEFAULT now(),
	user_id text NOT NULL,
	date text NOT NULL,
	on_time boolean NOT NULL DEFAULT FALSE,
	star_earned integer NOT NULL DEFAULT 0,
	PRIMARY KEY (user_id, date)
);

ALTER TABLE checkins ADD FOREIGN KEY (user_id) REFERENCES users (id);
`
	down := `
DROP TABLE IF EXISTS checkins;
`
	Collection.MustRegisterTx(wrap(up), wrap(down))
}
