package telegram

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gopkg.in/tucnak/telebot.v2"

	"api/model"
)

type BotHandler struct {
	bot *telebot.Bot
	svc *Service
}

func NewBotHandler(bot *telebot.Bot, svc *Service) *BotHandler {
	return &BotHandler{bot: bot, svc: svc}
}

func (h *BotHandler) Start() error {
	if err := h.bot.SetCommands(botCommands); err != nil {
		return fmt.Errorf("set commands: %w", err)
	}

	h.bot.Handle("/top", h.handleTop(false))
	h.bot.Handle("/all", h.handleTop(true))
	h.bot.Handle("/help", h.displayHelp())
	h.bot.Handle("/start", h.start())
	h.bot.Handle("/me", h.getMe())

	log.Info().Msg("bot running...")
	h.bot.Start()
	return nil
}

func (h *BotHandler) handleTop(all bool) interface{} {
	return func(m *telebot.Message) {
		ctx := context.TODO()
		users, err := model.UsersWithBalance().All(ctx, boil.GetContextDB())
		//users, err := h.balanceSvc.Users(ctx, api.Query{})
		if err != nil {
			_, _ = h.bot.Send(m.Chat, "Bot failed: "+err.Error())
			return
		}

		top := math.MaxInt32
		if !all {
			top = 3
			if t, err := strconv.Atoi(strings.TrimSpace(m.Payload)); err == nil {
				top = t
			}
			if top <= 0 {
				top = 1
			}
		}
		str := marshalTopUsers(users, top)
		_, err = h.bot.Send(m.Chat, str)
		if err != nil {
			log.Err(err).Msg("send failed")
		}
	}
}

func (h *BotHandler) displayHelp() interface{} {
	return func(m *telebot.Message) {
		log.Info().Interface("chat", m.Chat).Str("recipient", m.Chat.Recipient()).Msg("")
		_, _ = h.bot.Send(m.Chat, usage)
	}
}

func (h *BotHandler) start() interface{} {
	return func(m *telebot.Message) {
		token := strings.TrimSpace(m.Payload)
		if len(token) == 0 {
			_, _ = h.bot.Send(m.Chat, "Welcome to Verisheet Bot.\n\n"+usage)
			return
		}

		ctx := context.TODO()
		u, err := h.svc.AuthTelegramUser(ctx, token, m.Sender.Recipient())
		if err != nil {
			_, _ = h.bot.Send(m.Chat, "Bot failed: "+err.Error())
			return
		}

		_, _ = h.bot.Send(m.Chat, "Welcome to Verisheet Bot. Your name is "+u.Name+".\n\n"+usage)
	}
}

func (h *BotHandler) getMe() interface{} {
	return func(m *telebot.Message) {
		ctx := context.TODO()
		u, err := h.svc.GetUserWithBalanceByTelegramID(ctx, m.Sender.Recipient())
		if err != nil {
			if err == ErrUserNotFound {
				_, _ = h.bot.Send(m.Chat, "Bot failed: Unauthorized")
				return
			}
			_, _ = h.bot.Send(m.Chat, "Bot failed: "+err.Error())
			return
		}

		_, _ = h.bot.Send(m.Chat, marshalUser(u))
	}
}

type Users []model.UserWithBalance

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].Balance > u[j].Balance
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func marshalUser(user *model.UserWithBalance) string {
	return fmt.Sprintf("%s: %s (vnđ)", user.Name, humanize.Comma(int64(user.Balance)))
}

func marshalTopUsers(users Users, top int) string {
	sort.Sort(users)
	bf := bytes.NewBuffer(nil)
	if top < math.MaxInt32 {
		bf.WriteString(fmt.Sprintf("Top ±%d users:", top))
	} else {
		bf.WriteString("All users:")
	}
	dot := false
	for i, u := range users {
		if i >= top && i < len(users)-top {
			if !dot {
				dot = true
				bf.WriteString("\n......")
			}
			continue
		}
		bf.WriteString(fmt.Sprintf("\n%d. ", i+1))
		bf.WriteString(marshalUser(&u))
	}
	return bf.String()
}
