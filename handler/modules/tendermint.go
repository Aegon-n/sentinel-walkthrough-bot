package modules

import (
	"github.com/Aegon-n/sentinel-bot/handler/buttons"
	"github.com/Aegon-n/sentinel-bot/handler/constants"
	"github.com/Aegon-n/sentinel-bot/handler/messages"
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
		handleTMModuleList(bot, queryID, chatID, msgID,"10", "Linux", messages.EthWinListOfModulesMsg)

	case "TM-Linux-Module10":
		handleTMModule(bot, queryID, chatID, msgID,"20", "Linux", messages.EthWindowsModule10)

	case "TM-Linux-Module20":
		handleTMModule(bot, queryID, chatID, msgID, "30","Linux", messages.EthWindowsModule20)

	case "TM-Linux-Module30":
		handleTMModule(bot, queryID, chatID, msgID, "40","Linux", messages.EthWindowsModule30)

	case "TM-Linux-Module40":
		handleTMModule(bot, queryID, chatID, msgID, "50","Linux", messages.TMWindowsModule40)

	case "TM-Linux-Module50":
		handleTMModule(bot, queryID, chatID, msgID, "60","Linux",  messages.EthWindowsModule50)

	case "TM-Linux-Module60":
		handleTMModule(bot, queryID, chatID, msgID, "70","Linux",  messages.EthWindowsModule60)

	case "TM-Linux-Module70":
		handleTMLastModule(bot, queryID, chatID, msgID, "60","Linux",messages.LastModuleMsg)

	case "TM-Linux-DownloadDoc":
		handleDownload(bot, update, "Linux")
	}
}

func handleTMWindowsModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {

	case "TM-Windows-Module00":
		handleTMModuleList(bot, queryID, chatID, msgID,"10", "Windows", messages.EthWinListOfModulesMsg)

	case "TM-Windows-Module10":
		handleTMModule(bot, queryID, chatID, msgID, "20","Windows", messages.EthWindowsModule10)

	case "TM-Windows-Module20":
		handleTMModule(bot, queryID, chatID, msgID, "30","Windows", messages.EthWindowsModule20)

	case "TM-Windows-Module30":
		handleTMModule(bot, queryID, chatID, msgID, "40","Windows", messages.EthWindowsModule30)

	case "TM-Windows-Module40":
		handleTMModule(bot, queryID, chatID, msgID, "50","Windows", messages.TMWindowsModule40)

	case "TM-Windows-Module50":
		handleTMModule(bot, queryID, chatID, msgID, "60","Windows", messages.EthWindowsModule50)

	case "TM-Windows-Module60":
		handleTMModule(bot, queryID, chatID, msgID, "70","Windows", messages.EthWindowsModule60)

	case "TM-Windows-Module70":
		handleTMLastModule(bot, queryID, chatID, msgID, "60","Windows",messages.LastModuleMsg)

	case "TM-Windows-DownloadDoc":
		handleDownload(bot, update, "Windows")
	}
}

func handleTMMacModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {

	case "TM-Mac-Module0":
		handleTMModuleList(bot, queryID, chatID, msgID,"10", "Mac", messages.EthWinListOfModulesMsg)
	case "TM-Mac-Module10":
		handleTMModule(bot, queryID, chatID, msgID,"20","Mac", messages.EthWindowsModule10)

	case "TM-Mac-Module20":
		handleTMModule(bot, queryID, chatID, msgID, "30","Mac", messages.EthWindowsModule20)

	case "TM-Mac-Module30":
		handleTMModule(bot, queryID, chatID, msgID, "40","Mac", messages.EthWindowsModule30)

	case "TM-Mac-Module40":
		handleTMModule(bot, queryID, chatID, msgID, "50","Mac", messages.TMWindowsModule40)

	case "TM-Mac-Module50":
		handleTMModule(bot, queryID, chatID, msgID, "60","Mac",  messages.EthWindowsModule50)

	case "TM-Mac-Module60":
		handleTMModule(bot, queryID, chatID, msgID, "70","Mac",  messages.EthWindowsModule60)

	case "TM-Mac-Module70":
		handleTMLastModule(bot, queryID, chatID, msgID, "60","Mac",messages.LastModuleMsg)

	case "TM-Mac-DownloadDoc":
		handleDownload(bot, update, "Mac")
	}
}

func handleTMModuleList(Bot *tgbotapi.BotAPI, queryID string, chatID int64, msgID int, next, platform, txt string){
	answeredCallback(Bot, queryID)
	msg := tgbotapi.NewMessage(chatID, txt)

	if platform == "Linux"{
		msg.ReplyMarkup = buttons.ModulesListButton("TM-Linux-Module"+next, "TM-Linux-DownloadDoc" ,constants.VideoUrl)
	}
	if platform == "Windows" {
		msg.ReplyMarkup = buttons.ModulesListButton("TM-Windows-Module"+next, "TM-Windows-DownloadDoc" ,constants.VideoUrl)
	}
	if platform == "Mac" {
		msg.ReplyMarkup = buttons.ModulesListButton("TM-Mac-Module"+next, "TM-Mac-DownloadDoc" ,constants.VideoUrl)
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