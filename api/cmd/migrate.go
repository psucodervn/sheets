package cmd

import (
	"api/config"
	"api/model"
	"api/pkg/database"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:  "migrate",
	RunE: runMigration,
}

func runMigration(cmd *cobra.Command, args []string) error {
	var cfg config.MigrationConfig
	envconfig.MustProcess("", &cfg)

	db := database.MustNewPostgresGorm(cfg.Postgres)
	err := db.AutoMigrate(
		&model.User{},
		&model.Transaction{},
	).Error
	return err
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
