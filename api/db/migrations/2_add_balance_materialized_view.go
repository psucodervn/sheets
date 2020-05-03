package migrations

func init() {
	up := `
CREATE MATERIALIZED VIEW balance
AS
	SELECT u.id as user_id, COALESCE(SUM(tu.value), 0) as value
	FROM
		users u LEFT JOIN
			(SELECT items->>'id' as user_id, -(items->>'value')::FLOAT as value
				FROM transactions tx, jsonb_array_elements(tx.participants) items
				WHERE tx.deleted_at IS NULL
			UNION
			SELECT items->>'id' as user_id, (items->>'value')::FLOAT as value
				FROM transactions tx, jsonb_array_elements(tx.payers) items
				WHERE tx.deleted_at IS NULL
			) AS tu
			ON u.id = tu.user_id
	GROUP BY u.id
	ORDER BY value DESC, u.id
WITH DATA;

CREATE UNIQUE INDEX balance_user_id ON balance (user_id);
`
	down := `DROP MATERIALIZED VIEW balance;`

	Collection.MustRegisterTx(wrap(up), wrap(down))
}
