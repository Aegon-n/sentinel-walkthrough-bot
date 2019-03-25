package modules

import (
	"github.com/Aegon-n/sentinel-bot/handler/buttons"
	"github.com/Aegon-n/sentinel-bot/handler/messages/en_messages"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strconv"
)


func HandleTMModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update, platform string) {
	if platform == "Linux" {
		handleTMLinuxModules(bot, update)
	}
	if platform == "Windows" {
		handleTMWindowsModules(bot, update)
	}
	if platform == "Mac" {
		handleTMMacModules(bot, update)
	}
}

func handleTMLinuxModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {

	case "TM-Linux-Module0":
		handleTMModuleList(bot, queryID, chatID,"10", "Linux", en_messages.EthWinListOfModulesMsg)

	case "TM-Linux-Module10":
		handleTMModule(bot, queryID, chatID, msgID,"20", "Linux", en_messages.EthWindowsModule10)

	case "TM-Linux-Module20":
		handleTMModule(bot, queryID, chatID, msgID, "30","Linux", en_messages.EthWindowsModule20)

	case "TM-Linux-Module30":
		handleTMModule(bot, queryID, chatID, msgID, "40","Linux", en_messages.EthWindowsModule30)

	case "TM-Linux-Module40":
		handleTMModule(bot, queryID, chatID, msgID, "50","Linux", en_messages.TMWindowsModule40)

	case "TM-Linux-Module50":
		handleTMModule(bot, queryID, chatID, msgID, "60","Linux",  en_messages.EthWindowsModule50)

	case "TM-Linux-Module60":
		handleTMLastModule(bot, queryID, chatID, msgID, "50","Linux",  en_messages.EthWindowsModule60)

	case "TM-Linux-DownloadDoc":
		handleDownload(bot, update, "Linux","Doc")

	case "TM-Linux-VideoSend":
		handleDownload(bot, update, "Linux","Video")

	}
}

func handleTMWindowsModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {

	case "TM-Windows-Module0":
		handleTMModuleList(bot, queryID, chatID,"10", "Windows", en_messages.EthWinListOfModulesMsg)

	case "TM-Windows-Module10":
		handleTMModule(bot, queryID, chatID, msgID, "20","Windows", en_messages.EthWindowsModule10)

	case "TM-Windows-Module20":
		handleTMModule(bot, queryID, chatID, msgID, "30","Windows", en_messages.EthWindowsModule20)

	case "TM-Windows-Module30":
		handleTMModule(bot, queryID, chatID, msgID, "40","Windows", en_messages.EthWindowsModule30)

	case "TM-Windows-Module40":
		handleTMModule(bot, queryID, chatID, msgID, "50","Windows", en_messages.TMWindowsModule40)

	case "TM-Windows-Module50":
		handleTMModule(bot, queryID, chatID, msgID, "60","Windows", en_messages.EthWindowsModule50)

	case "TM-Windows-Module60":
		handleTMLastModule(bot, queryID, chatID, msgID, "50","Windows", en_messages.EthWindowsModule60)

	case "TM-Windows-DownloadDoc":
		handleDownload(bot, update, "Windows","Doc")

	case "TM-Windows-VideoSend":
		handleDownload(bot, update, "Windows","Doc")

	}

}

func handleTMMacModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {

	case "TM-Mac-Module0":
		handleTMModuleList(bot, queryID, chatID,"10", "Mac", en_messages.EthWinListOfModulesMsg)
	case "TM-Mac-Module10":
		handleTMModule(bot, queryID, chatID, msgID,"20","Mac", en_messages.EthWindowsModule10)

	case "TM-Mac-Module20":
		handleTMModule(bot, queryID, chatID, msgID, "30","Mac", en_messages.EthWindowsModule20)

	case "TM-Mac-Module30":
		handleTMModule(bot, queryID, chatID, msgID, "40","Mac", en_messages.EthWindowsModule30)

	case "TM-Mac-Module40":
		handleTMModule(bot, queryID, chatID, msgID, "50","Mac", en_messages.TMWindowsModule40)

	case "TM-Mac-Module50":
		handleTMModule(bot, queryID, chatID, msgID, "60","Mac",  en_messages.EthWindowsModule50)

	case "TM-Mac-Module60":
		handleTMLastModule(bot, queryID, chatID, msgID, "50","Mac",  en_messages.EthWindowsModule60)

	case "TM-Mac-DownloadDoc":
		handleDownload(bot, update, "Mac","Doc")

	case "TM-Mac-VideoSend":
		handleDownload(bot, update, "Mac","Video")

	}
}

func handleTMModuleList(Bot *tgbotapi.BotAPI, queryID string, chatID int64, next, platform, txt string){
	answeredCallback(Bot, queryID)
	msg := tgbotapi.NewMessage(chatID, txt)

	if platform == "Linux"{
		msg.ReplyMarkup =  buttons.GetButtons("LinuxTMModulesButtonList")
	}
	if platform == "Windows" {
		msg.ReplyMarkup =  buttons.GetButtons("WindowsTMModulesButtonList")
	}
	if platform == "Mac" {
		msg.ReplyMarkup =  buttons.GetButtons("MacTMModulesButtonList")
	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
}
func handleTMModule(Bot *tgbotapi.BotAPI, queryID string, chatID int64, msgID int, next, platform, txt string) {
	answeredCallback(Bot, queryID)
	log.Println("in handleTMmodule")
	nextIdx,_ := strconv.Atoi(next)
	prev := nextIdx -20
	prevstr := strconv.Itoa(prev)

	msg := tgbotapi.NewEditMessageText(chatID, msgID, txt)
	btns := tgbotapi.EditMessageReplyMarkupConfig{}

	if platform == "Linux" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.PersistentNavButtons("TM-Linux-Module"+prevstr, "TM-Linux-Module"+next, "TM-Linux-Module"+next))
	}
	if platform == "Windows" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.PersistentNavButtons("TM-Windows-Module"+prevstr, "TM-Windows-Module"+next, "TM-Windows-Module"+next))
	}
	if platform == "Mac" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.PersistentNavButtons("TM-Mac-Module"+prevstr, "TM-Mac-Module"+next, "TM-Mac-Module"+next))
	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)

}
func handleTMLastModule(Bot *tgbotapi.BotAPI, queryID string, chatID int64, msgID int, prev, platform, txt string) {
	answeredCallback(Bot, queryID)
	msg := tgbotapi.NewEditMessageText(chatID, msgID, txt)
	btns := tgbotapi.EditMessageReplyMarkupConfig{}

	if platform == "Linux"{
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.LastModuleButtons("TM-Linux-Module"+prev))
	}
	if platform == "Windows" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.LastModuleButtons("TM-Windows-Module"+prev))
	}
	if platform == "Mac" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.LastModuleButtons("TM-Mac-Module"+prev))
	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)

}