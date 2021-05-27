package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/go-co-op/gocron"
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
	cfg := config.MustReadBotConfig()
	log.Debug().Interface("config", cfg).Send()

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

	go startCronJob(bot, telegramSvc)

	return botHandler.Start()
}

func startCronJob(bot *tb.Bot, svc *telegram.Service) {
	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Every(1).Day().At("02:25").Do(func() {
		// check for weekday
		wd := time.Now().Weekday()
		if wd < time.Monday || wd > time.Friday {
			log.Info().Str("weekday", wd.String()).Msg("ignore remind checkin")
			return
		}

		ctx := context.Background()
		users, err := svc.ListUserNotCheckedInToday(ctx)
		if err != nil {
			log.Err(err).Send()
			return
		}

		if len(users) == 0 {
			log.Info().Msg("All users have checked in!")
		}

		msg := "You have not checked in today. Do you want to check in now?"
		for _, u := range users {
			_, err = bot.Send(newUser(u.TelegramID.String), msg, &tb.ReplyMarkup{
				InlineKeyboard: [][]tb.InlineButton{
					{{Unique: "checkin", Text: "Yes", Data: u.ID}, {Unique: "checkin", Text: "No", Data: ""}},
				},
			})
			if err != nil {
				log.Err(err).Str("user", u.Name).Send()
			}
		}
	})

	s.StartAsync()
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
