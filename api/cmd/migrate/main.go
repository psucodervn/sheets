package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/psucodervn/go/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

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

func main() {
	logger.InitFromEnv()

	if err := migrateCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("migration failed")
	}
	log.Info().Msg("migration done")
}
