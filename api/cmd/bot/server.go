package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"math"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/sync/errgroup"
	"gopkg.in/tucnak/telebot.v2"

	"api/model"
	"api/proto"
)

type NotificationServer struct {
	bot *telebot.Bot
	db  *sql.DB
}

func NewNotificationServer(bot *telebot.Bot, db *sql.DB) *NotificationServer {
	return &NotificationServer{bot: bot, db: db}
}

func (s *NotificationServer) NotifyTransaction(ctx context.Context, req *sheet.NotifyTransactionRequest) (*sheet.NotifyTransactionResponse, error) {
	resp := &sheet.NotifyTransactionResponse{Success: false}
	txLog, err := s.getTransactionLog(ctx, req.TransactionLogId)
	if err != nil {
		if err == sql.ErrNoRows {
			resp.ErrorCode = "NOT_FOUND"
		} else {
			log.Ctx(ctx).Err(err).Msg("FindTransactionLog failed")
			resp.ErrorCode = "DATABASE_ERROR"
		}
		return resp, nil
	}

	userChanges := s.calcUserChanges(ctx, txLog)
	log.Ctx(ctx).Info().Interface("changes", userChanges).Msg("")
	g, ctx := errgroup.WithContext(ctx)
	for userID := range userChanges {
		userID := userID
		g.Go(func() error {
			return s.notifyTransactionLog(ctx, userID, userChanges[userID], txLog)
		})
	}
	err = g.Wait()
	if err != nil {
		resp.ErrorCode = err.Error()
	} else {
		resp.Success = true
	}
	return resp, nil
}

func (s *NotificationServer) getTransactionLog(ctx context.Context, id string) (*model.TransactionLog, error) {
	return model.TransactionLogs(
		model.TransactionLogWhere.ID.EQ(id),
		qm.Load(model.TransactionLogRels.Actor),
		qm.Load(model.TransactionLogRels.Transaction, qm.Or(model.TransactionColumns.ID+" IN ($1) AND "+model.TransactionColumns.DeletedAt+" IS NOT NULL")),
	).One(ctx, s.db)
}

func (s *NotificationServer) calcUserChanges(ctx context.Context, txLog *model.TransactionLog) map[string]float64 {
	m := make(map[string]float64)
	switch txLog.Action {
	case model.ActionCreate:
		adjustUserChanges(m, new(transaction).FromModel(txLog.R.Transaction), 1)
	case model.ActionRemove:
		adjustUserChanges(m, new(transaction).FromModel(txLog.R.Transaction), -1)
	case model.ActionUpdate:
		log.Info().Interface("meta", txLog.Meta).Msg("")
		var tmp struct {
			OldTX transaction `json:"oldTx"`
			NewTX transaction `json:"newTx"`
		}
		_ = mapstructure.Decode(txLog.Meta, &tmp)
		adjustUserChanges(m, &tmp.OldTX, -1)
		adjustUserChanges(m, &tmp.NewTX, 1)
	}
	return m
}

func (s *NotificationServer) notifyTransactionLog(ctx context.Context, userID string, change float64, txLog *model.TransactionLog) error {
	u, err := model.UsersWithBalance(model.UserWhere.ID.EQ(userID)).One(ctx, s.db)
	if err != nil {
		return err
	}
	if u.TelegramID.IsZero() {
		log.Ctx(ctx).Warn().Msg(u.Name + " has no telegram_id")
		// ignore
		return nil
	}

	bf := bytes.NewBuffer(nil)
	bf.WriteString(fmt.Sprintf(
		"%s has %sd transaction '%s'",
		txLog.Meta["username"], strings.ToLower(string(txLog.Action)), txLog.R.Transaction.Summary),
	)

	abs := math.Abs(change)
	if abs < 1e-3 {
		bf.WriteString(", but your balance has not changed.")
	} else {
		s := humanize.Comma(int64(abs))
		if change < 0 {
			bf.WriteString(", so your balance has decreased by " + s + " (vnđ).")
		} else {
			bf.WriteString(", so your balance has increased by " + s + " (vnđ).")
		}
	}
	bf.WriteString("\nYour current balance is " + humanize.Comma(int64(u.Balance)) + " (vnđ).")

	log.Info().Str("message", bf.String()).Msg("send to " + u.Name)
	_, err = s.bot.Send(newUser(u.TelegramID.String), bf.String())
	return err
}

func adjustUserChanges(m map[string]float64, tx *transaction, base float64) {
	for _, v := range tx.Payers {
		m[v.ID] += v.Value * base
	}
	for _, v := range tx.Participants {
		m[v.ID] -= v.Value * base
	}
}

type user struct {
	ID string
}

func newUser(id string) *user {
	return &user{ID: id}
}

func (u user) Recipient() string {
	return u.ID
}

type transaction struct {
	ID           string        `json:"id"`
	Payers       model.Changes `json:"payers"`
	Participants model.Changes `json:"participants"`
	Value        float64       `json:"value"`
	Summary      string        `json:"summary"`
}

func (t *transaction) FromModel(tx *model.Transaction) *transaction {
	return &transaction{
		ID:           tx.ID,
		Payers:       tx.Payers,
		Participants: tx.Participants,
		Value:        tx.Value,
		Summary:      tx.Summary,
	}
}
