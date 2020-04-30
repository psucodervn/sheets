package cmd

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/psucodervn/go/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"api/internal/api"
	"api/internal/balance"
	"api/internal/config"
	"api/internal/point"
	"api/pkg/adapter"
	"api/pkg/wakatime"
)

var (
	rootCmd = &cobra.Command{
		Use: "api",
		RunE: func(cmd *cobra.Command, args []string) error {
			var cfg config.ApiConfig
			envconfig.MustProcess("", &cfg)
			log.Debug().Interface("config", cfg).Msg("api starting")

			return runApiServer(cfg)
		},
	}
)

func runApiServer(cfg config.ApiConfig) error {
	db := adapter.MustNewPostgresGorm(cfg.Postgres)

	fetcher := balance.NewApiFetcherFromEnv()
	userRepo := balance.NewGormPostgresUserRepository(db)
	txRepo := balance.NewGormPostgresTransactionRepository(db)
	balanceSvc := balance.NewBaseService(fetcher, userRepo, txRepo)
	balanceHandler := balance.NewHandler(balanceSvc)

	importer := balance.NewOldImporter(fetcher, userRepo, txRepo)
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

func init() {
	logger.InitFromEnv()
}

func Execute() error {
	return rootCmd.Execute()
}
