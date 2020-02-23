package cmd

import (
	"api/api"
	"api/auth"
	"api/balance"
	"api/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/markbates/goth/providers/google"
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
	fetcher := balance.NewApiFetcherFromEnv()
	balanceSvc := balance.NewBaseService(fetcher)
	balanceHandler := balance.NewHandler(balanceSvc)

	authCfg := cfg.Google.Auth
	googleAuthProvider := google.New(
		authCfg.ClientID, authCfg.ClientSecret, authCfg.CallbackURL,
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	)

	authHandler := auth.NewHandler(googleAuthProvider)

	srv := api.NewServer()
	srv.Bind(balanceHandler, authHandler)
	return srv.Serve(cfg.Address, cfg.TLS)
}

func Execute() error {
	return rootCmd.Execute()
}
