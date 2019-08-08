package handler

import (
	"gopkg.in/telegram-bot-api.v4"
	"strings"
)

func HandleUpdates(b *tgbotapi.BotAPI, u tgbotapi.Update){
	module := strings.Split(u.CallbackQuery.Data, "-")[1]
	switch module {
		case "Medium":
			updates.MediumUpdates(bot, update)


		case "Reddit":
			updates.Reddit_updates(bot, update)


		case "Twitter":
			updates.Twitter_updates(bot, update)
	}
}