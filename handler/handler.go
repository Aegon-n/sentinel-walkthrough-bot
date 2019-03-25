package handler

import (
	"fmt"
	"github.com/Aegon-n/sentinel-bot/handler/buttons"
	"github.com/Aegon-n/sentinel-bot/handler/constants"
	"github.com/Aegon-n/sentinel-bot/handler/messages"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleGreet(Bot *tgbotapi.BotAPI, update *tgbotapi.Update )  {
	username := update.Message.From.UserName
	chatID := update.Message.Chat.ID
	txt := fmt.Sprintf(messages.WelcomeGreetMsg, username)+"\n"+messages.SelectwalkthroughMsg
	msg := tgbotapi.NewMessage(chatID,txt)
	msg.ParseMode = tgbotapi.ModeMarkdown
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
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
}
func HandleCallbackQuery(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

	switch update.CallbackQuery.Data {

	case "Home":
		handleHome(bot, update, update.CallbackQuery.From.UserName)

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

	case "Exit":
		handleExit(bot, update)

	default:
		chatID := update.CallbackQuery.Message.Chat.ID
		txt := "Not implemented"
		msg := tgbotapi.NewMessage(chatID, txt)
		bot.Send(msg)
	}
}

func handleHome(Bot *tgbotapi.BotAPI, update *tgbotapi.Update ,username string){

	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	msg := tgbotapi.NewEditMessageText(chatID, msgID, fmt.Sprintf(messages.WelcomeGreetMsg, username)+"\n"+messages.AppSelectMsg)
	msg.ParseMode = tgbotapi.ModeMarkdown
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.HomeButtons("Sentinel-Desktop","Sentinel-Mobile"))

	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)
}

func handleAppVersion(Bot *tgbotapi.BotAPI, update *tgbotapi.Update, version string){
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	msg := tgbotapi.EditMessageTextConfig{}

	replyMarkup := tgbotapi.InlineKeyboardMarkup{}
	if version == "Desktop" {
		msg = tgbotapi.NewEditMessageText(chatID, msgID, messages.DesktopOSSelectMsg)
		replyMarkup = buttons.DesktopOsButtons("Linux","Windows","Mac")
	}
	if version == "Mobile" {
		msg = tgbotapi.NewEditMessageText(chatID, msgID, messages.MobileOSSelectMsg)
		replyMarkup = buttons.MobileOsButtons("Android","IOS")
	}
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, replyMarkup)
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)

}
func handleOs(Bot *tgbotapi.BotAPI, update *tgbotapi.Update, os string){
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	msg := tgbotapi.EditMessageTextConfig{}
	btns := tgbotapi.EditMessageReplyMarkupConfig{}
	if os == "Android"{

		msg = tgbotapi.NewEditMessageText(chatID, msgID, messages.MobileListOfMOdulesMsg)
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID, msgID,
			buttons.ModulesListButton("Android10",constants.DownloadUrl ,constants.VideoUrl))

	}
	if os == "IOS" {

		msg = tgbotapi.NewEditMessageText(chatID, msgID, messages.MobileListOfMOdulesMsg)
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID, msgID,
			buttons.ModulesListButton("IOS10",constants.DownloadUrl ,constants.VideoUrl))

	}
	if os == "Linux" {

		msg = tgbotapi.NewEditMessageText(chatID, msgID, messages.LinuxNetworkSelectMsg)
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.TestNetButtons("ETH-Linux-Module0", "TM-Linux-Module0"))

	}
	if os == "Windows" {
		msg = tgbotapi.NewEditMessageText(chatID, msgID, messages.WindowsNetworkSelectMsg)
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.TestNetButtons("ETH-Windows-Module0", "TM-Windows-Module0"))

	}
	if os == "Mac" {
		msg = tgbotapi.NewEditMessageText(chatID, msgID, messages.MacNetworkSelectMsg)
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.TestNetButtons("ETH-Mac-Module0", "TM-Mac-Module0"))

	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)
}

func answeredCallback(Bot *tgbotapi.BotAPI, queryId string){
	config := tgbotapi.CallbackConfig{queryId,"",false,"",0}
	Bot.AnswerCallbackQuery(config)
}
func handleExit(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	msg := tgbotapi.NewEditMessageText(chatID, msgID, messages.ExitMsg)
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)

}