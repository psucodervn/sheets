package telegram

import (
	"gopkg.in/tucnak/telebot.v2"
)

var (
	botCommands = []telebot.Command{
		{
			Text:        "checkin",
			Description: "Checkin",
		},
		{
			Text:        "top",
			Description: "list top users by balance",
		},
		{
			Text:        "all",
			Description: "list all users with balance",
		},
		{
			Text:        "me",
			Description: "fetch my information",
		},
		{
			Text:        "help",
			Description: "display help",
		},
	}

	usage = `Verisheet Bot commands:
- /checkin: check-in
  /checkin out: check-out
  /checkin 9h25: check in at 9:25 AM
- /all: list all users
- /top <n>: list top n users with highest/lowest balance
- /me: fetch my information
- /help: display help`
)
