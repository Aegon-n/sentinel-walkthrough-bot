package modules

import (
	"github.com/Aegon-n/sentinel-bot/handler/buttons"
	"github.com/Aegon-n/sentinel-bot/handler/constants"
	"github.com/Aegon-n/sentinel-bot/handler/messages/en_messages"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strconv"
)

func HandleEthModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update, platform string) {
	if platform == "Linux" {
		handleETHLinuxModules(bot, update)
	}
	if platform == "Windows" {
		handleETHWindowsModules(bot, update)
	}
	if platform == "Mac" {
		handleETHMacModules(bot, update)
	}
}

func handleETHWindowsModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {

	case "ETH-Windows-Module0":
		handleETHModuleList(bot, queryID, chatID, "10","Windows", en_messages.EthWinListOfModulesMsg)

	case "ETH-Windows-Module10":
		//log.Println(en_messages.NewEthWindowsModule10)
		handleETHModule(bot, queryID, chatID, msgID,"20","Windows", en_messages.EthWindowsModule10)

	case "ETH-Windows-Module20":
		handleETHModule(bot, queryID, chatID, msgID, "30","Windows", en_messages.EthWindowsModule20)

	case "ETH-Windows-Module30":
		handleETHModule(bot, queryID, chatID, msgID, "40","Windows", en_messages.EthWindowsModule30)

	case "ETH-Windows-Module40":
		handleETHModule(bot, queryID, chatID, msgID, "50","Windows", en_messages.EthWindowsModule40)

	case "ETH-Windows-Module50":
		handleETHModule(bot, queryID, chatID, msgID, "60","Windows", en_messages.EthWindowsModule50)

	case "ETH-Windows-Module60":
		handleETHModule(bot, queryID, chatID, msgID, "70","Windows", en_messages.EthWindowsModule60)

	case "ETH-Windows-Module70":
		handleEthLastModule(bot, queryID, chatID, msgID, "60","Windows", en_messages.LastModuleMsg)

	case "ETH-Windows-DownloadDoc":
		handleDownload(bot, update, "Windows","Doc")

	case "ETH-Windows-VideoSend":
		handleDownload(bot, update, "Windows","Video")

	}

}

func handleETHLinuxModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {

	case "ETH-Linux-Module0":
		handleETHModuleList(bot, queryID, chatID, "10","Linux", en_messages.EthWinListOfModulesMsg)

	case "ETH-Linux-Module10":
		log.Println("here")
		handleETHModule(bot, queryID, chatID, msgID,"20","Linux", en_messages.EthWindowsModule10)

	case "ETH-Linux-Module20":
		handleETHModule(bot, queryID, chatID, msgID, "30","Linux", en_messages.EthWindowsModule20)

	case "ETH-Linux-Module30":
		handleETHModule(bot, queryID, chatID, msgID, "40","Linux", en_messages.EthWindowsModule30)

	case "ETH-Linux-Module40":
		handleETHModule(bot, queryID, chatID, msgID, "50","Linux", en_messages.EthWindowsModule40)

	case "ETH-Linux-Module50":
		handleETHModule(bot, queryID, chatID, msgID, "60","Linux", en_messages.EthWindowsModule50)

	case "ETH-Linux-Module60":
		handleEthLastModule(bot, queryID, chatID, msgID, "50","Linux", en_messages.EthWindowsModule60)

	case "ETH-Linux-DownloadDoc":
		handleDownload(bot, update, "Linux","Doc")

	case "ETH-Linux-VideoSend":
		handleDownload(bot, update, "Linux","Video")
	}
}

func handleETHMacModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {

	case "ETH-Mac-Module0":
		handleETHModuleList(bot, queryID, chatID, "10","Mac", en_messages.EthWinListOfModulesMsg)

	case "ETH-Mac-Module10":
		handleETHModule(bot, queryID, chatID, msgID,"20","Mac", en_messages.EthWindowsModule10)

	case "ETH-Mac-Module20":
		handleETHModule(bot, queryID, chatID, msgID, "30","Mac", en_messages.EthWindowsModule20)

	case "ETH-Mac-Module30":
		handleETHModule(bot, queryID, chatID, msgID, "40","Mac", en_messages.EthWindowsModule30)

	case "ETH-Mac-Module40":
		handleETHModule(bot, queryID, chatID, msgID, "50","Mac", en_messages.EthWindowsModule40)

	case "ETH-Mac-Module50":
		handleETHModule(bot, queryID, chatID, msgID, "60","Mac", en_messages.EthWindowsModule50)

	case "ETH-Mac-Module60":
		handleEthLastModule(bot, queryID, chatID, msgID,"50","Mac",en_messages.EthWindowsModule60)

	case "ETH-Mac-DownloadDoc":
		handleDownload(bot, update, "Mac","Doc")

	case "ETH-Mac-VideoSend":
		handleDownload(bot, update, "Mac","Video")
	}
}
func handleETHModuleList(Bot *tgbotapi.BotAPI, queryID string, chatID int64, next, platform, txt string){
	answeredCallback(Bot, queryID)
	msg := tgbotapi.NewMessage(chatID, txt)

	if platform == "Linux"{
		msg.ReplyMarkup =  buttons.GetButtons("LinuxEthModulesButtonList")
	}
	if platform == "Windows" {
		msg.ReplyMarkup =  buttons.GetButtons("WindowsEthModulesButtonList")
	}
	if platform == "Mac" {
		msg.ReplyMarkup =  buttons.GetButtons("MacEthModulesButtonList")
	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
}

func handleETHModule(Bot *tgbotapi.BotAPI, queryID string, chatID int64, msgID int, next, platform, txt string) {
	answeredCallback(Bot, queryID)
	log.Println("in handleTMmodule")
	nextIdx,_ := strconv.Atoi(next)
	prev := nextIdx -20
	prevstr := strconv.Itoa(prev)

	msg := tgbotapi.NewEditMessageText(chatID, msgID, txt)
	btns := tgbotapi.EditMessageReplyMarkupConfig{}

	if platform == "Linux" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.PersistentNavButtons("ETH-Linux-Module"+prevstr, "ETH-Linux-Module"+next, "ETH-Linux-Module"+next))
	}
	if platform == "Windows" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.PersistentNavButtons("ETH-Windows-Module"+prevstr, "ETH-Windows-Module"+next, "ETH-Windows-Module"+next))
	}
	if platform == "Mac" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.PersistentNavButtons("ETH-Mac-Module"+prevstr, "ETH-Mac-Module"+next, "ETH-Mac-Module"+next))
	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)

}

func answeredCallback(Bot *tgbotapi.BotAPI, queryId string){
	config := tgbotapi.CallbackConfig{queryId,"",false,"",0}
	Bot.AnswerCallbackQuery(config)
}

func handleDownload(Bot *tgbotapi.BotAPI, update *tgbotapi.Update, platform, typ string) {
	answeredCallback(Bot, update.CallbackQuery.ID)
	chatID := update.CallbackQuery.Message.Chat.ID
	if typ == "Doc" {
		res := tgbotapi.DocumentConfig{}
		switch platform {
		case "Linux":
			res = tgbotapi.NewDocumentShare(chatID,constants.LinuxPdfUrl)
		case "Windows":
			res = tgbotapi.NewDocumentShare(chatID,constants.WindowsPdfUrl)
		case "Mac":
			res = tgbotapi.NewDocumentShare(chatID,constants.MacPdfUrl)
		}
		Bot.Send(res)
	}
	if typ == "Video" {
		msg := tgbotapi.NewMessage(chatID,constants.VideoUrl)
		switch platform {
		case "Linux":
			msg = tgbotapi.NewMessage(chatID, constants.LinuxVideoUrl)
		case "Windows":
			msg = tgbotapi.NewMessage(chatID, constants.WindowsVideoUrl)
		case "Mac":
			msg = tgbotapi.NewMessage(chatID, constants.MacVideoUrl)
		}
		msg.ParseMode = tgbotapi.ModeHTML
		Bot.Send(msg)
	}
}
func handleEthLastModule(Bot *tgbotapi.BotAPI, queryID string, chatID int64, msgID int, prev, platform, txt string) {
	answeredCallback(Bot, queryID)
	msg := tgbotapi.NewEditMessageText(chatID, msgID, txt)
	btns := tgbotapi.EditMessageReplyMarkupConfig{}

	if platform == "Linux"{
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.LastModuleButtons("ETH-Linux-Module"+prev))
	}
	if platform == "Windows" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.LastModuleButtons("ETH-Windows-Module"+prev))
	}
	if platform == "Mac" {
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.LastModuleButtons("ETH-Mac-Module"+prev))
	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)

}