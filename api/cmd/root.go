package cmd

import (
	"api/api/delivery/http"
	"api/balance/repository"
	"api/balance/usecase"
	"api/config"
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
	fetcher := repository.NewApiFetcherFromEnv()
	balanceUseCase := usecase.NewUseCase(fetcher)

	srv := http.NewServer(balanceUseCase)
	return srv.Serve(cfg.Address, cfg.TLS)
}

func Execute() error {
	return rootCmd.Execute()
}
