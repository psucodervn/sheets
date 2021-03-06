package telegram

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

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
	h.bot.Handle("/checkin", h.checkIn())
	h.bot.Handle("/checkin_list", h.listCheckins())
	h.bot.Handle("/stars", h.listStars())
	h.bot.Handle("/remind", h.remind())

	h.bot.Handle(&telebot.InlineButton{Unique: "checkin"}, func(cb *telebot.Callback) {
		uID := cb.Data
		_ = h.bot.Delete(cb.Message)
		if len(uID) == 0 || uID == "Cancel" {
			return
		}
		h.doCheckin(cb.Message)
	})

	//go func() {
	//	for a := range h.bot.Updates {
	//		fmt.Println(a.Message.Sticker)
	//	}
	//}()

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

func (h *BotHandler) checkIn() interface{} {
	return h.doCheckin
}

func (h *BotHandler) doCheckin(m *telebot.Message) {
	ctx := boil.WithDebug(context.TODO(), true)
	u, err := h.svc.GetUserWithBalanceByTelegramID(ctx, getUserID(m))
	if err != nil {
		if err == ErrUserNotFound {
			_, _ = h.bot.Send(m.Chat, "Bot failed: Unauthorized")
			return
		}
		_, _ = h.bot.Send(m.Chat, "Bot failed: "+err.Error())
		return
	}

	payload := strings.ToLower(strings.TrimSpace(m.Payload))
	if payload == "out" {
		h.checkOut(ctx, m, u)
		return
	}

	t := time.Now()
	if len(payload) > 0 {
		t, err = parseTime(payload)
		if err != nil {
			_, _ = h.bot.Send(m.Chat, fmt.Sprintf("Invalid time `"+payload+"`. Hint: use 9h25 or 9:25 format for 9h25 AM."))
			return
		}
	}

	if isWeekend(t) {
		_, _ = h.bot.Send(m.Chat, fmt.Sprintf("You can't check-in on weekend 🤨🤨"))
		return
	}
	if t.After(time.Now()) {
		_, _ = h.bot.Send(m.Chat, fmt.Sprintf("Impossible 😡😡"))
		_ = h.sendSticker(m.Chat, StickerBoNgheMay)
		return
	}

	ci, err := h.svc.CheckUserIn(ctx, u, t)
	if err != nil {
		if err == ErrAlreadyCheckedIn {
			_, _ = h.bot.Send(m.Chat, "You have already checked in today!")
			return
		}
		_, _ = h.bot.Send(m.Chat, "Check-in failed: "+err.Error())
		_ = h.sendSticker(m.Chat, StickerDoAnO)
		return
	}

	var msg string
	st := checkInTime(t)
	if ci.OnTime {
		msg = fmt.Sprintf("Congratz, you earned %d star for checked in on time (%s) 🥳🥳", ci.StarEarned, st)
		_, _ = h.bot.Send(m.Chat, msg)
		//_ = h.sendSticker(m.Chat, StickerThugLife)
		//msg = fmt.Sprintf("➡️ <b><i>%s</i></b> checked in at <code>%s</code> and earned <b>%d</b> star 🥳🥳", u.Name, st, ci.StarEarned)
		//_ = h.sendToAll(ctx, u, msg)
	} else {
		msg = fmt.Sprintf("Too late! It's %s. You earned nothing 🤣🤣", st)
		_, _ = h.bot.Send(m.Chat, msg)
		_ = h.sendSticker(m.Chat, StickerDenThoi)
		//msg = fmt.Sprintf("➡️ <b><i>%s</i></b> checked in at <code>%s</code> but it's too late 🤣🤣", u.Name, st)
		//_ = h.sendToAll(ctx, u, msg)
	}
}

func getUserID(m *telebot.Message) string {
	uid := strconv.Itoa(m.Sender.ID)
	if m.Sender.IsBot {
		uid = strconv.Itoa(int(m.Chat.ID))
	}
	return uid
}

func isWeekend(t time.Time) bool {
	wd := t.In(LocalZone).Weekday()
	return wd == time.Sunday || wd == time.Saturday
}

func parseTime(payload string) (time.Time, error) {
	payload = strings.Replace(payload, ":", "h", -1)
	t, err := time.ParseInLocation("15h04", payload, LocalZone)
	if err != nil {
		return time.Time{}, err
	}
	y, m, d := time.Now().In(LocalZone).Date()
	t = t.AddDate(y, int(m)-1, d-1)
	return t, nil
}

func checkInTime(t time.Time) string {
	return t.In(LocalZone).Format("03:04 PM")
}

func (h *BotHandler) sendSticker(to telebot.Recipient, sticker Sticker) error {
	_, err := h.bot.Send(to, &telebot.Sticker{
		File: telebot.File{FileID: string(sticker)},
	})
	return err
}

func (h *BotHandler) checkOut(ctx context.Context, m *telebot.Message, u *model.UserWithBalance) {
	ci, err := h.svc.CheckUserOut(ctx, u, time.Now())
	if err == ErrNotCheckedIn {
		_, _ = h.bot.Send(m.Chat, "You haven't checked in yet!")
		_ = h.sendSticker(m.Chat, StickerBoNgheMay)
		return
	} else if err != nil {
		_, _ = h.bot.Send(m.Chat, "Check-out failed: "+err.Error())
		_ = h.sendSticker(m.Chat, StickerDoAnO)
		return
	}

	var msg string
	if ci.StarEarned > 0 {
		msg = fmt.Sprintf("You checked out and lost %d star 🤔🤔", ci.StarEarned)
		_, _ = h.bot.Send(m.Chat, msg)
		//msg = fmt.Sprintf("↩️ <b><i>%s</i></b> checked out and lost <b><i>%d</i></b> star 🤔🤔", u.Name, ci.StarEarned)
		//_ = h.sendToAll(ctx, u, msg)
	} else {
		msg = fmt.Sprintf("Checked out 😏😏")
		_, _ = h.bot.Send(m.Chat, msg)
		//msg = fmt.Sprintf("↩️ <b><i>%s</i></b> checked out 😏😏", u.Name)
		//_ = h.sendToAll(ctx, u, msg)
	}
}

func (h *BotHandler) sendToAll(ctx context.Context, user *model.UserWithBalance, msg string) error {
	users, err := h.svc.ListUserHasTelegramID(ctx)
	if err != nil {
		return ErrDatabase
	}
	for _, u := range users {
		if u.TelegramID.IsZero() || u.TelegramID == user.TelegramID {
			continue
		}
		if _, err = h.bot.Send(newUser(u.TelegramID.String), msg, &telebot.SendOptions{
			ParseMode: telebot.ModeHTML,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (h *BotHandler) listCheckins() interface{} {
	return func(m *telebot.Message) {
		today := time.Now().In(LocalZone)
		ctx := context.TODO()
		cis, err := h.svc.ListCheckins(ctx, today)
		if err != nil {
			_, _ = h.bot.Send(m.Chat, "List check-in failed: "+err.Error())
			_ = h.sendSticker(m.Chat, StickerDoAnO)
			return
		}

		if len(cis) == 0 {
			_, _ = h.bot.Send(m.Chat, "No one have checked in yet!")
			return
		}
		bf := bytes.NewBuffer(nil)
		for _, ci := range cis {
			bf.WriteString(fmt.Sprintf("<code>%s</code>: %s\n", checkInTime(ci.Time), ci.R.User.Name))
		}
		_, _ = h.bot.Send(m.Chat, bf.String(), &telebot.SendOptions{
			ParseMode: telebot.ModeHTML,
		})
	}
}

func (h *BotHandler) listStars() interface{} {
	return func(m *telebot.Message) {
		ctx := context.TODO()
		res, err := h.svc.ListStarsInCurrentMonth(ctx)
		if err != nil {
			_, _ = h.bot.Send(m.Chat, "List failed: "+err.Error())
			return
		}

		if len(res) == 0 {
			_, _ = h.bot.Send(m.Chat, "No one have earned stars yet.")
			return
		}

		sb := &strings.Builder{}
		for _, u := range res {
			sb.WriteString(fmt.Sprintf("%s: %d\n", u.Name, int(u.Stars)))
		}
		_, _ = h.bot.Send(m.Chat, sb.String())
	}
}

func (h *BotHandler) remind() interface{} {
	return func(m *telebot.Message) {
		ctx := context.TODO()
		if err := h.remindCheckin(ctx, m); err != nil {
			_, _ = h.bot.Send(m.Chat, err.Error())
		}
	}
}

func (h *BotHandler) remindCheckin(ctx context.Context, m *telebot.Message) error {
	users, err := h.svc.ListUserNotCheckedInToday(ctx)
	if err != nil {
		return err
	}

	if len(users) == 0 {
		_, _ = h.bot.Send(m.Chat, "All users have checked in!")
	}

	msg := "You have not checked in today. Do you want to check in now?"
	for _, u := range users {
		_, err = h.bot.Send(newUser(u.TelegramID.String), msg, &telebot.ReplyMarkup{
			InlineKeyboard: [][]telebot.InlineButton{
				{{Unique: "checkin", Text: "Yes", Data: u.ID}, {Unique: "checkin", Text: "No", Data: ""}},
			},
		})
		if err != nil {
			log.Err(err).Str("user", u.Name).Send()
		}
	}
	return nil
}
