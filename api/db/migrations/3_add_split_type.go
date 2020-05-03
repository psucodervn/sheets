package migrations

func init() {
	up := `
alter table transactions add split_type int default 0 not null;
`
	down := `
alter table transactions drop column split_type;
`

	Collection.MustRegisterTx(wrap(up), wrap(down))
}
