package modules

import (
	"github.com/Aegon-n/sentinel-bot/handler/buttons"
	"github.com/Aegon-n/sentinel-bot/handler/messages/en_messages"
	"gopkg.in/telegram-bot-api.v4"
	"strconv"
)
func HandleMobileModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update, platform string) {
	if platform == "Android" {
		handleAndroidModules(bot, update)
	}
	if platform == "IOS" {
		handleIOSModules(bot, update)
	}
}
func handleAndroidModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	switch module {
	case "Mobile-Android-Module0":
		handleMobileModuleList(bot, queryID, chatID, en_messages.AndroidMobileListOfModulesMsg, "Android")
	case "Mobile-Android-Module10":
		handleMobileModule(bot, queryID, chatID, msgID, "20",en_messages.AndroidModule10, "Android")
	case "Mobile-Android-Module20":
		handleMobileModule(bot, queryID, chatID, msgID, "30",en_messages.AndroidModule20, "Android")
	case "Mobile-Android-Module30":
		handleMobileModule(bot, queryID, chatID, msgID, "40",en_messages.AndroidModule30, "Android")
	case "Mobile-Android-Module40":
		handleMobileModule(bot, queryID, chatID, msgID, "50",en_messages.AndroidModule40, "Android")
	case "Mobile-Android-Module50":
		handleMobileModule(bot, queryID, chatID, msgID, "60",en_messages.AndroidModule50, "Android")
	case "Mobile-Android-Module60":
		handleMobileLastModule(bot, queryID, chatID, msgID, "50", en_messages.AndroidModule60, "Android")
	case "Mobile-Android-DownloadDoc":
		handleDownload(bot, update, "Android","Doc")
	case "Mobile-Android-VideoSend":
		handleDownload(bot, update, "Android","Video")
	}

}
func handleIOSModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	//msgID := update.CallbackQuery.Message.MessageID
	switch module {
	case "Mobile-IOS-Module0":
		handleMobileModuleList(bot, queryID, chatID, en_messages.IOSMobileListOfModulesMsg, "IOS")
	/*case "Mobile-IOS-Module10":
		handleMobileModule(bot, queryID, chatID, msgID, "20", en_messages.IOSModule10, "IOS")
	case "Mobile-IOS-Module20":
		handleMobileModule(bot, queryID, chatID, msgID, "30", en_messages.IOSModule20, "IOS")
	case "Mobile-IOS-Module30":
		handleMobileModule(bot, queryID, chatID, msgID, "40", en_messages.IOSModule30, "IOS")
	case "Mobile-IOS-Module40":
		handleMobileModule(bot, queryID, chatID, msgID, "50", en_messages.IOSModule40, "IOS")
	case "Mobile-IOS-Module50":
		handleMobileModule(bot, queryID, chatID, msgID, "60", en_messages.IOSModule50, "IOS")
	case "Mobile-IOS-Module60":
		handleMobileLastModule(bot, queryID, chatID, msgID, "50", en_messages.IOSModule60, "IOS")
	case "Mobile-IOS-DownloadDoc":
		handleDownload(bot, update, "IOS", "Doc")
	case "Mobile-IOS-VideoSend":
		handleDownload(bot, update, "IOS", "Video")*/
	}
}
func handleMobileModule(Bot *tgbotapi.BotAPI, queryID string, chatID int64, msgID int, next, message, typ string) {
	answeredCallback(Bot, queryID)
	nextIdx, _ := strconv.Atoi(next)
	prev := nextIdx - 20
	prevstr := strconv.Itoa(prev)

	msg := tgbotapi.NewEditMessageText(chatID, msgID, message)
	btns := tgbotapi.EditMessageReplyMarkupConfig{}
	if typ == "Android"{
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.PersistentNavButtons("Mobile-Android-Module"+prevstr, "Mobile-Android-Module"+next, "Mobile-Android-Module"+next))
	}
	if typ == "IOS" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.PersistentNavButtons("Mobile-IOS-Module"+prevstr, "Mobile-IOS-Module"+next, "Mobile-IOS-Module"+next))
	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)
}
func handleMobileModuleList(Bot *tgbotapi.BotAPI, queryID string, chatID int64, message, typ string) {
	answeredCallback(Bot, queryID)
	msg := tgbotapi.MessageConfig{}
	if typ == "Android" {
		msg = tgbotapi.NewMessage(chatID, message)
		msg.ReplyMarkup = buttons.GetButtons("AndroidModulesButtonList")
	}
	if typ == "IOS" {
		msg = tgbotapi.NewMessage(chatID, message)
		//msg.ReplyMarkup = buttons.GetButtons("IOSModulesButtonList")
	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
}

func handleMobileLastModule(Bot *tgbotapi.BotAPI, queryID string, chatID int64, msgID int, prev, message, typ string) {
	answeredCallback(Bot, queryID)
	msg := tgbotapi.NewEditMessageText(chatID, msgID, message)
	btns := tgbotapi.EditMessageReplyMarkupConfig{}

	if typ == "Android"{
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.LastModuleButtons("Mobile-Android-Module"+prev))
	}
	if typ == "IOS" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.LastModuleButtons("Mobile-IOS-Module"+prev))
	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)
}
