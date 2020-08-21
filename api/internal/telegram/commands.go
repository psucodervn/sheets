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
			Text:        "checkin_list",
			Description: "List Checkin",
		},
		{
			Text:        "top",
			Description: "List top users by balance",
		},
		{
			Text:        "all",
			Description: "List all users with balance",
		},
		{
			Text:        "me",
			Description: "Fetch my information",
		},
		{
			Text:        "help",
			Description: "Display help",
		},
	}

	usage = `Verisheet Bot commands:
- /checkin:      check-in
  /checkin out:  check-out
  /checkin 9h25: check-in at 9:25 AM
- /checkin_list: list check-in today
- /all: list all users
- /top <n>: list top n users with highest/lowest balance
- /me: fetch my information
- /help: display help`
)
