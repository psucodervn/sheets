package cmd

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/psucodervn/go/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"api/api"
	"api/balance"
	"api/config"
	"api/pkg/database"
	"api/point"
	"api/wakatime"
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

	importer := balance.NewImporter(fetcher, userRepo, txRepo)
	importFn := func() {
		if err := importer.Run(); err != nil {
			log.Err(err).Msg("import failed")
		}
	}
	go func() {
		importFn()
		for range time.Tick(1 * time.Minute) {
			importFn()
		}
	}()

	pointSvc := point.NewRestService(cfg.Jira.Username, cfg.Jira.Password, cfg.Jira.Host)
	wakaSvc := wakatime.NewApiFetcher(cfg.Wakatime.ApiKey)
	reportSvc := point.NewBaseReportService(pointSvc, wakaSvc, cfg.Wakatime.Leaderboard, cfg.Wakatime.MapID)
	pointHandler := point.NewHttpHandler(pointSvc, reportSvc)

	srv := api.NewServer()
	srv.Bind(
		balanceHandler,
		pointHandler,
	)
	return srv.Serve(cfg.Address, cfg.TLS)
}

func Execute() error {
	return rootCmd.Execute()
}
