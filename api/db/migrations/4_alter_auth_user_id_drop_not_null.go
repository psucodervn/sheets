package migrations

func init() {
	up := `ALTER TABLE auth_identities ALTER COLUMN user_id DROP NOT NULL;`
	down := `ALTER TABLE auth_identities ALTER COLUMN user_id SET NOT NULL;`

	Collection.MustRegisterTx(wrap(up), wrap(down))
}
