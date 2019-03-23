package handler

import (
	"fmt"
	"github.com/Aegon-n/sentinel-bot/handler/messages"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleGreet(Bot *tgbotapi.BotAPI, update *tgbotapi.Update )  {
	username := update.Message.From.UserName
	chatID := update.Message.Chat.ID
	txt := fmt.Sprintf(messages.WelcomeGreetMsg, username)+"\n"+messages.SelectwalkthroughMsg
	msg := tgbotapi.NewMessage(chatID,txt)
	Bot.Send(msg)

}
func HandlerWalkThrough(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

	username := update.Message.From.UserName
	chatID := update.Message.Chat.ID
	txt := fmt.Sprintf(messages.WelcomeGreetMsg, username)+"\n"+messages.AppSelectMsg
	msg := tgbotapi.NewMessage(chatID,txt)

	btn1 := tgbotapi.NewInlineKeyboardButtonData("Sentinel-Desktop App","Sentinel-Desktop")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Sentinel-Mobile App","Sentinel-Mobile")

	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1,btn2}},
	}
	Bot.Send(msg)
}
func HandleCallbackQuery(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

	switch update.CallbackQuery.Data {
	case "Sentinel-Desktop":
		HandlerDesktop(bot, update)

	case "Sentinel-Mobile":
		HandlerMobile(bot, update)

	case "Android":
		HandleAndroid(bot, update)
	case "IOS":
		HandleIOS(bot, update)

	case "Linux":
		HandleLinux(bot, update)

	case "Windows":
		HandleWindows(bot, update)

	case "Mac":
		HandleMac(bot, update)

	case "TMWindows":
		HandleTMWindows(bot, update)

	case "ETHWindows":
		HandleETHWindows(bot, update)
	case "TMLinux":
		HandleTMLinux(bot, update)

	default:
		chatID := update.CallbackQuery.Message.Chat.ID
		txt := "Not implemented"
		msg := tgbotapi.NewMessage(chatID, txt)
		bot.Send(msg)
	}
}
func HandlerDesktop(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	//msgID := update.CallbackQuery.Message.MessageID
	txt := messages.OSSelectMsg
	msg := tgbotapi.NewMessage(chatID,txt)
	btn1 := tgbotapi.NewInlineKeyboardButtonData("☣ Linux","Linux")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("☣ Windows","Windows")
	btn3 := tgbotapi.NewInlineKeyboardButtonData(" ☣ Mac OS", "Mac")
	msg.ReplyMarkup =  tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1}, {btn2}, {btn3}},
	}
	//btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, replyMarkup)

	Bot.Send(msg)
	//Bot.Send(btns)
}

func HandlerMobile(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	//msgID := update.CallbackQuery.Message.MessageID
	txt := messages.OSSelectMsg
	msg := tgbotapi.NewMessage(chatID, txt)
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Android","Android")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("IOS","IOS")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1,btn2}},
	}

	Bot.Send(msg)
	//Bot.Send(btns)
}
func HandleAndroid(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	msg := tgbotapi.NewEditMessageText(chatID, msgID, messages.MobileListOfMOdulesMsg)
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Next","AndroidMobileInstructions10")
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1}},
	})
	Bot.Send(msg)
	Bot.Send(btns)

}
func HandleIOS(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

}
func HandleLinux(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	txt := messages.NetworkSelectMsg
	msg := tgbotapi.NewEditMessageText(chatID,msgID,txt)
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Etherium","ETHLinux")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Tendermint","TMLinux")
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1,btn2}},
	})
	Bot.Send(msg)
	Bot.Send(btns)
}
func HandleWindows(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	txt := messages.NetworkSelectMsg
	msg := tgbotapi.NewEditMessageText(chatID,msgID,txt)
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Etherium","ETHWindows")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Tendermint","TMWindows")
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1,btn2}},
	})
	Bot.Send(msg)
	Bot.Send(btns)
}
func HandleMac(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	txt := messages.NetworkSelectMsg
	msg := tgbotapi.NewEditMessageText(chatID, msgID, txt)
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Etherium","ETHMac")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Tendermint","TMMac")
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID,tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1,btn2}},
	})
	Bot.Send(msg)
	Bot.Send(btns)
}
func HandleETHWindows(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	//msgID := update.CallbackQuery.Message.MessageID
	txt := messages.EthWinListOfModulesMsg
	msg := tgbotapi.NewMessage(chatID, txt)
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Next","ETH-WinModule10")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1}},
	}
	Bot.Send(msg)
}

func HandleTMWindows(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	//msgID := update.CallbackQuery.Message.MessageID

	txt := messages.EthWinListOfModulesMsg
	msg := tgbotapi.NewMessage(chatID, txt)
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Next","TM-WinModule10")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1}},
	}

	Bot.Send(msg)
}

func HandleTMLinux(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	//msgID := update.CallbackQuery.Message.MessageID

	txt := messages.EthWinListOfModulesMsg
	msg := tgbotapi.NewMessage(chatID, txt)
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Next","TM-LinuxModule10")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1}},
	}

	Bot.Send(msg)
}

func HandleTMMac(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	//msgID := update.CallbackQuery.Message.MessageID

	txt := messages.EthWinListOfModulesMsg
	msg := tgbotapi.NewMessage(chatID, txt)
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Next","TM-MacModule10")
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1}},
	}

	Bot.Send(msg)

}

func answeredCallback(Bot *tgbotapi.BotAPI, queryId string){
	config := tgbotapi.CallbackConfig{queryId,"",false,"",0}
	Bot.AnswerCallbackQuery(config)
}