package main

import (
	"context"
	"database/sql"

	"gopkg.in/tucnak/telebot.v2"

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
	_, err := s.bot.Send(new(user), req.TransactionId)
	resp := &sheet.NotifyTransactionResponse{
		Success: err == nil,
	}
	if err != nil {
		resp.ErrorCode = err.Error()
	}
	return resp, nil
}

type user struct{}

func (u user) Recipient() string {
	return "371574526"
}
