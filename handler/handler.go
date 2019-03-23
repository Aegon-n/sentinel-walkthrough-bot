package handler

import (
	"fmt"
	"github.com/Aegon-n/sentinel-bot/handler/buttons"
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
		handleAppVersion(bot, update, "Desktop")

	case "Sentinel-Mobile":
		handleAppVersion(bot, update, "Mobile")

	case "Android":
		handleOs(bot, update, "Android")

	case "IOS":
		handleOs(bot, update, "IOS")

	case "Linux":
		handleOs(bot, update, "Linux")

	case "Windows":
		handleOs(bot, update, "Windows")

	case "Mac":
		handleOs(bot, update, "Mac")

	default:
		chatID := update.CallbackQuery.Message.Chat.ID
		txt := "Not implemented"
		msg := tgbotapi.NewMessage(chatID, txt)
		bot.Send(msg)
	}
}

func handleAppVersion(Bot *tgbotapi.BotAPI, update *tgbotapi.Update, version string){
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	msg := tgbotapi.NewEditMessageText(chatID, msgID, messages.OSSelectMsg)
	replyMarkup := tgbotapi.InlineKeyboardMarkup{}
	if version == "Desktop" {
		replyMarkup = buttons.DesktopOsButtons("Linux","Windows","Mac")
	}
	if version == "Mobile" {
		replyMarkup = buttons.MobileOsButtons("Android","IOS")
	}
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, replyMarkup)

	Bot.Send(msg)
	Bot.Send(btns)

}
func handleOs(Bot *tgbotapi.BotAPI, update *tgbotapi.Update, os string){
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	if os == "Android"{

		msg := tgbotapi.NewEditMessageText(chatID, msgID, messages.MobileListOfMOdulesMsg)
		btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.ModulesListButton("Android10"))
		Bot.Send(msg)
		Bot.Send(btns)

	}
	if os == "IOS" {

		msg := tgbotapi.NewEditMessageText(chatID, msgID, messages.MobileListOfMOdulesMsg)
		btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.ModulesListButton("IOS10"))
		Bot.Send(msg)
		Bot.Send(btns)

	}
	if os == "Linux" {

		msg := tgbotapi.NewEditMessageText(chatID, msgID, messages.NetworkSelectMsg)
		btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.TestNetButtons("ETH-Linux-Module00", "TM-Linux-Module00"))
		Bot.Send(msg)
		Bot.Send(btns)

	}
	if os == "Windows" {
		msg := tgbotapi.NewEditMessageText(chatID, msgID, messages.NetworkSelectMsg)
		btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.TestNetButtons("ETH-Windows-Module00", "TM-Windows-Module00"))
		Bot.Send(msg)
		Bot.Send(btns)

	}
	if os == "Mac" {
		msg := tgbotapi.NewEditMessageText(chatID, msgID, messages.NetworkSelectMsg)
		btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.TestNetButtons("ETH-Mac-Module00", "TM-Mac-Module00"))
		Bot.Send(msg)
		Bot.Send(btns)
	}
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