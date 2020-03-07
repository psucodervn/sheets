package importer

import (
	"api/balance"
	"api/config"
	"api/pkg/database"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:  "import",
	RunE: runImporter,
}

func runImporter(cmd *cobra.Command, args []string) error {
	var cfg config.ImporterConfig
	envconfig.MustProcess("", &cfg)

	fetcher := balance.NewApiFetcherFromEnv()
	db := database.MustNewPostgresGorm(cfg.Postgres)
	userRepo := balance.NewPostgresUserRepository(db)
	txRepo := balance.NewPostgresTransactionRepository(db)

	importer := NewImporter(fetcher, userRepo, txRepo)
	return importer.Run()
}

func Command() *cobra.Command {
	return importCmd
}
