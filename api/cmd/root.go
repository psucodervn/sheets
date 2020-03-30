package cmd

import (
	"api/api"
	"api/balance"
	"api/cmd/importer"
	"api/config"
	"api/pkg/database"
	"api/point"
	"github.com/kelseyhightower/envconfig"
	"github.com/psucodervn/go/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "api",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var cfg config.LogConfig
			envconfig.MustProcess("LOG", &cfg)
			logger.Init(cfg.Debug, cfg.Pretty)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var cfg config.ApiConfig
			envconfig.MustProcess("", &cfg)
			log.Debug().Interface("config", cfg).Msg("api starting")

			return runApiServer(cfg)
		},
	}
)

func runApiServer(cfg config.ApiConfig) error {
	db := database.MustNewPostgresGorm(cfg.Postgres)

	fetcher := balance.NewApiFetcherFromEnv()
	userRepo := balance.NewPostgresUserRepository(db)
	txRepo := balance.NewPostgresTransactionRepository(db)
	balanceSvc := balance.NewBaseService(fetcher, userRepo, txRepo)
	balanceHandler := balance.NewHandler(balanceSvc)

	pointSvc := point.NewRestService(cfg.Jira.Username, cfg.Jira.Password, cfg.Jira.Host)
	pointHandler := point.NewHttpHandler(pointSvc)

	srv := api.NewServer()
	srv.Bind(
		balanceHandler,
		pointHandler,
	)
	return srv.Serve(cfg.Address, cfg.TLS)
}

func init() {
	rootCmd.AddCommand(importer.Command())
}

func Execute() error {
	return rootCmd.Execute()
}
