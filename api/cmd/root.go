package cmd

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/psucodervn/go/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"api/db"
	"api/internal/api"
	"api/internal/balance"
	"api/internal/config"
	"api/internal/point"
	"api/pkg/adapter"
	"api/pkg/wakatime"
)

func init() {
	logger.InitFromEnv()
}

var (
	RootCmd = &cobra.Command{
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
	gormConn := adapter.MustNewPostgresGorm(cfg.Postgres)
	conn := db.ConnectDB(cfg.Postgres)

	userRepo := balance.NewGormPostgresUserRepository(gormConn)
	txRepo := balance.NewGormPostgresTransactionRepository(gormConn)
	balanceSvc := balance.NewService(userRepo, txRepo, conn)
	balanceHandler := balance.NewHandler(balanceSvc)

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
	return RootCmd.Execute()
}
