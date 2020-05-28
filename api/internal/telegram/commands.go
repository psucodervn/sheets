package telegram

import (
	"gopkg.in/tucnak/telebot.v2"
)

var (
	botCommands = []telebot.Command{
		{
			Text:        "top",
			Description: "list top users by balance",
		},
		{
			Text:        "all",
			Description: "list all users with balance",
		},
		{
			Text:        "help",
			Description: "display help",
		},
		{
			Text:        "me",
			Description: "fetch my information",
		},
	}

	usage = `Verisheet Bot commands:
- /all: list all users
- /top <n>: list top n users with highest/lowest balance
- /help: display help
- /me: fetch my information`
)
