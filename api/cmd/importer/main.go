package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/psucodervn/go/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"api/db"
	"api/internal/balance"
	"api/internal/config"
	"api/pkg/adapter"
)

var importCmd = &cobra.Command{
	Use:  "import",
	RunE: runImporter,
}

func runImporter(cmd *cobra.Command, args []string) error {
	var cfg config.ImporterConfig
	envconfig.MustProcess("", &cfg)

	fetcher := balance.NewApiFetcherFromEnv()
	conn := db.ConnectDB(cfg.Postgres)
	imp := balance.NewImporter(fetcher, conn)
	if err := imp.Run(); err != nil {
		log.Fatal().Err(err).Msg("import failed")
	}

	log.Info().Msg("import done")
	return nil
}

func runOldImporter(cmd *cobra.Command, args []string) error {
	var cfg config.ImporterConfig
	envconfig.MustProcess("", &cfg)

	fetcher := balance.NewApiFetcherFromEnv()
	conn := adapter.MustNewPostgresGorm(cfg.Postgres)
	userRepo := balance.NewGormPostgresUserRepository(conn)
	txRepo := balance.NewGormPostgresTransactionRepository(conn)
	importer := balance.NewOldImporter(fetcher, userRepo, txRepo)
	return importer.Run()
}

func main() {
	logger.InitFromEnv()

	if err := importCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("import failed")
	}
}
