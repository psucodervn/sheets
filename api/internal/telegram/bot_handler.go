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
	"gopkg.in/tucnak/telebot.v2"

	"api/internal/api"
	"api/internal/balance"
	"api/model"
)

type BotHandler struct {
	bot        *telebot.Bot
	balanceSvc balance.Service
}

func NewBotHandler(bot *telebot.Bot, balanceSvc balance.Service) *BotHandler {
	return &BotHandler{bot: bot, balanceSvc: balanceSvc}
}

func (h *BotHandler) Start() error {
	if err := h.bot.SetCommands(botCommands); err != nil {
		return fmt.Errorf("set commands: %w", err)
	}

	h.bot.Handle("/top", h.handleTop(false))
	h.bot.Handle("/all", h.handleTop(true))
	h.bot.Handle("/help", h.displayHelp())

	log.Info().Msg("bot running...")
	h.bot.Start()
	return nil
}

func (h *BotHandler) handleTop(all bool) interface{} {
	return func(m *telebot.Message) {
		ctx := context.TODO()
		users, err := h.balanceSvc.Users(ctx, api.Query{})
		if err != nil {
			_, _ = h.bot.Send(m.Chat, "Bot failed: "+err.Error())
			return
		}
		log.Info().Msgf("payload: `%s`", m.Payload)
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
		_, _ = h.bot.Send(m.Chat, usage)
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

func marshalUser(user model.UserWithBalance) string {
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
		bf.WriteString(marshalUser(u))
	}
	return bf.String()
}
