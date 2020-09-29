package cmd

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/psucodervn/go/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api/db"
	"api/internal/api"
	"api/internal/auth"
	"api/internal/balance"
	"api/internal/config"
	"api/internal/daysoff"
	"api/internal/handler"
	"api/internal/point"
	"api/internal/telegram"
	"api/internal/user"
	"api/pkg/wakatime"
	sheet "api/proto"
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
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), nil)
	if err != nil {
		log.Panic().Msgf("connect gorm database failed: %v", err)
	}

	notiConn, err := grpc.Dial("bot:8080", grpc.WithInsecure())
	if err != nil {
		log.Panic().Msgf("dial notification server failed: %v", err)
	}
	notiClient := sheet.NewNotificationServiceClient(notiConn)

	jwtMW := api.AuthJWT([]byte(cfg.Auth.JWTSecret))

	balanceSvc := balance.NewService(conn, notiClient)
	balanceHandler := balance.NewHandler(balanceSvc, jwtMW)

	pointSvc := point.NewRestService(cfg.Jira.Username, cfg.Jira.Password, cfg.Jira.Host)
	wakaSvc := wakatime.NewApiFetcher(cfg.Wakatime.ApiKey)
	reportSvc := point.NewBaseReportService(pointSvc, wakaSvc, cfg.Wakatime.Leaderboard, cfg.Wakatime.MapID)
	pointHandler := point.NewHttpHandler(pointSvc, reportSvc, jwtMW)

	ggCfg := cfg.Auth.Google
	ggConf := oauth2.Config{
		Endpoint: google.Endpoint,
		ClientID: ggCfg.ClientID, ClientSecret: ggCfg.ClientSecret, RedirectURL: ggCfg.RedirectURL,
	}
	authSvc := auth.NewService(conn, []byte(cfg.Auth.JWTSecret), ggConf)
	userSvc := user.NewService(conn)
	telegramSvc := telegram.NewService(conn, cfg.TelegramBotName)
	authHandler := auth.NewHandler(authSvc, userSvc, telegramSvc, jwtMW)

	daysOffSvc := daysoff.NewService(gormDB)
	daysOffHandler := handler.NewDaysOffHandler(daysOffSvc, jwtMW)

	srv := api.NewServer()
	srv.Bind(
		balanceHandler,
		pointHandler,
		authHandler,
		daysOffHandler,
	)
	return srv.Serve(cfg.Address, cfg.TLS)
}

func Execute() error {
	return RootCmd.Execute()
}
