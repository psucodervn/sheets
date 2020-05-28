package main

import (
	"fmt"
	"net"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/psucodervn/go/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc"
	tb "gopkg.in/tucnak/telebot.v2"

	"api/db"
	"api/internal/config"
	"api/internal/telegram"
	sheet "api/proto"
)

func init() {
	logger.InitFromEnv()
}

var (
	botCmd = &cobra.Command{
		Use:  "bot",
		RunE: runBot,
	}
)

func runBot(cmd *cobra.Command, args []string) error {
	var cfg config.BotConfig
	envconfig.MustProcess("", &cfg)

	conn := db.ConnectDB(cfg.Postgres)
	boil.SetDB(conn)

	bot, err := tb.NewBot(tb.Settings{
		Token:  cfg.Telegram.BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return err
	}

	srv := NewNotificationServer(bot, conn)
	go startServer(srv)

	telegramSvc := telegram.NewService(conn, "")
	botHandler := telegram.NewBotHandler(bot, telegramSvc)
	return botHandler.Start()
}

func startServer(srv *NotificationServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Panic().Msgf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	sheet.RegisterNotificationServiceServer(grpcServer, srv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic().Err(err).Msg("serve failed")
	}
}

func main() {
	if err := botCmd.Execute(); err != nil {
		log.Fatal().Msgf("run bot failed: %v", err)
	}
}
