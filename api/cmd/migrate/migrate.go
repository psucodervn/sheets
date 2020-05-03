package migrate

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"

	"api/cmd"
	"api/db"
	"api/db/migrations"
	"api/internal/config"
	"api/pkg/migration"
)

var migrateCmd = &cobra.Command{
	Use:  "migrate",
	RunE: runMigration,
}

func runMigration(cmd *cobra.Command, args []string) error {
	var cfg config.MigrationConfig
	envconfig.MustProcess("", &cfg)

	conn := db.ConnectGoPGDB(cfg.Postgres)
	m := migration.NewMigration(conn, migrations.Collection.SetTableName("go_migrations"), "migrate")

	c := "up"
	if len(args) > 0 {
		c = args[0]
	}
	m.Run(c)
	return nil
}

func init() {
	cmd.RootCmd.AddCommand(migrateCmd)
}
