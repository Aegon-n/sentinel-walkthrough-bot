package handler

import (
	"fmt"
	"strings"

	"github.com/Aegon-n/sentinel-bot/sno/buttons"
	"github.com/Aegon-n/sentinel-bot/sno/helper"
	"github.com/Aegon-n/sentinel-bot/updates/messages"
	updates "github.com/Aegon-n/sentinel-bot/updates/services"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleUpdatesHome(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	chatID := helper.GetchatID(u)
	txt := messages.HomeMsg + "\n\n" + messages.ChooseOpt
	btns := buttons.GetButtons("UpdatesHomeBtns")
	if u.Message != nil {
		msg := tgbotapi.NewMessage(chatID, txt)
		fmt.Println(btns)
		msg.ReplyMarkup = btns
		msg.ParseMode = tgbotapi.ModeMarkdown
		b.Send(msg)
		return
	}
	helper.Send(b, u, txt, &btns)
	return
}
func HandleUpdates(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	module := strings.Split(u.CallbackQuery.Data, "-")[1]
	switch module {
	case "Home":
		HandleUpdatesHome(b, u)
	case "Medium":
		updates.MediumUpdates(b, u)

	case "Reddit":
		updates.Reddit_updates(b, u)

	case "Twitter":
		updates.Twitter_updates(b, u)
	}
}
