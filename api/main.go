package main

import (
	"github.com/rs/zerolog/log"

	"api/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("run failed")
	}
}
