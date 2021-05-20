package migrations

func init() {
	up := `
ALTER TABLE users ADD COLUMN remind_checkin bool DEFAULT FALSE;
`
	down := `
ALTER TABLE users DROP COLUMN remind_checkin;
`
	Collection.MustRegisterTx(wrap(up), wrap(down))
}
