package modules

import (
	"github.com/Aegon-n/sentinel-bot/handler/messages"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strconv"
)

func HandleEthModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {
	case "ETH-WinModule10":
		HandleETHModule(bot, queryID, chatID, msgID,"20", messages.EthWindowsModule10)

	case "ETH-WinModule20":
		HandleETHModule(bot, queryID, chatID, msgID, "30",messages.EthWindowsModule20)

	case "ETH-WinModule30":
		HandleETHModule(bot, queryID, chatID, msgID, "40",messages.EthWindowsModule30)

	case "ETH-WinModule40":
		HandleETHModule(bot, queryID, chatID, msgID, "50",messages.EthWindowsModule40)

	case "ETH-WinModule50":
		HandleETHModule(bot, queryID, chatID, msgID, "60", messages.EthWindowsModule50)

	case "ETH-WinModule60":
		HandleETHModule(bot, queryID, chatID, msgID, "70", messages.EthWindowsModule60)

	case "WinModule30":

	case "WinModule31":

	case "WinModule32":

	case "WinModule40":

	}
}

func HandleEthModule(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	txt := messages.EthWindowsModule10
	msg := tgbotapi.NewEditMessageText(chatID,msgID, txt)

	btn1 := tgbotapi.NewInlineKeyboardButtonData("Prev","EthWinModule10")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Next","EthWinModule12")
	markup := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1,btn2}},
	}
	msg2 := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, markup)
	Bot.Send(msg)
	Bot.Send(msg2)
}
func HandleSKip(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	txt := messages.EthWindowsModule20
	msg := tgbotapi.NewEditMessageText(chatID,msgID, txt)
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Prev","EthWinModule12")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Prev","EthWinModule30")
	btn3 := tgbotapi.NewInlineKeyboardButtonData("Next","EthWinModule21")
	markup := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1,btn2,btn3}},
	}
	msg2 := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, markup)
	Bot.Send(msg)
	Bot.Send(msg2)
}
func answeredCallback(Bot *tgbotapi.BotAPI, queryId string){
	config := tgbotapi.CallbackConfig{queryId,"",false,"",0}
	Bot.AnswerCallbackQuery(config)
}

func HandleETHModule(Bot *tgbotapi.BotAPI, queryID string, chatID int64, msgID int, id, txt string) {
	answeredCallback(Bot, queryID)
	log.Println("in handleEthmodule")
	present,_ := strconv.Atoi(id)
	prev := present -20
	prevstr := strconv.Itoa(prev)

	msg := tgbotapi.NewEditMessageText(chatID, msgID, txt)
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID,
		msgID,persistentButtons("ETH-WinModule"+prevstr,"ETH-WinModule"+id,"ETH-WinModule"+id))

	Bot.Send(msg)
	Bot.Send(btns)

}

func persistentButtons(data1, data2, data3 string) tgbotapi.InlineKeyboardMarkup {
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Prev",data1)
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Skip",data2)
	btn3 := tgbotapi.NewInlineKeyboardButtonData("Next",data3)
	btns := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1,btn2,btn3}},
	}
	return btns
}