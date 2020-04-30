package migrations

import (
	"github.com/go-pg/migrations/v7"
)

// Collection contains migrations collection
var Collection = migrations.NewCollection()

func wrap(query string) func(db migrations.DB) error {
	return func(db migrations.DB) error {
		_, err := db.Exec(query)
		return err
	}
}
