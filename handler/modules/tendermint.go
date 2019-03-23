package modules

import (
	"github.com/Aegon-n/sentinel-bot/handler/messages"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strconv"
)

func HandleTMModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {
	case "TM-WinModule10":
		HandleTMModule(bot, queryID, chatID, msgID,"20", messages.EthWindowsModule10)

	case "TM-WinModule20":
		HandleTMModule(bot, queryID, chatID, msgID, "30",messages.EthWindowsModule20)

	case "TM-WinModule30":
		HandleTMModule(bot, queryID, chatID, msgID, "40",messages.EthWindowsModule30)

	case "TM-WinModule40":
		HandleTMModule(bot, queryID, chatID, msgID, "50",messages.TMWindowsModule40)

	case "TM-WinModule50":
		HandleTMModule(bot, queryID, chatID, msgID, "60", messages.EthWindowsModule50)

	case "TM-WinModule60":
		HandleTMModule(bot, queryID, chatID, msgID, "70", messages.EthWindowsModule60)
	}
}

func HandleTMModule(Bot *tgbotapi.BotAPI, queryID string, chatID int64, msgID int, id, txt string) {
	answeredCallback(Bot, queryID)
	log.Println("in handleTMmodule")
	present,_ := strconv.Atoi(id)
	prev := present -20
	prevstr := strconv.Itoa(prev)

	msg := tgbotapi.NewEditMessageText(chatID, msgID, txt)
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID,
		msgID,persistentButtons("TM-WinModule"+prevstr,"TM-WinModule"+id,"TM-WinModule"+id))

	Bot.Send(msg)
	Bot.Send(btns)

}