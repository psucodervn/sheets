package cmd

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/psucodervn/go/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"api/db"
	"api/internal/api"
	"api/internal/auth"
	"api/internal/balance"
	"api/internal/config"
	"api/internal/point"
	"api/internal/user"
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
	conn := db.ConnectDB(cfg.Postgres)
	boil.SetDB(conn)

	jwtMW := api.AuthJWT([]byte(cfg.Auth.JWTSecret))

	balanceSvc := balance.NewService(conn)
	balanceHandler := balance.NewHandler(balanceSvc, jwtMW)

	pointSvc := point.NewRestService(cfg.Jira.Username, cfg.Jira.Password, cfg.Jira.Host)
	wakaSvc := wakatime.NewApiFetcher(cfg.Wakatime.ApiKey)
	reportSvc := point.NewBaseReportService(pointSvc, wakaSvc, cfg.Wakatime.Leaderboard, cfg.Wakatime.MapID)
	pointHandler := point.NewHttpHandler(pointSvc, reportSvc)

	ggCfg := cfg.Auth.Google
	ggConf := oauth2.Config{
		Endpoint: google.Endpoint,
		ClientID: ggCfg.ClientID, ClientSecret: ggCfg.ClientSecret, RedirectURL: ggCfg.RedirectURL,
	}
	authSvc := auth.NewService([]byte(cfg.Auth.JWTSecret), ggConf)
	userSvc := user.NewService(conn)
	authHandler := auth.NewHandler(authSvc, userSvc, jwtMW)

	srv := api.NewServer()
	srv.Bind(
		balanceHandler,
		pointHandler,
		authHandler,
	)
	return srv.Serve(cfg.Address, cfg.TLS)
}

func Execute() error {
	return RootCmd.Execute()
}
