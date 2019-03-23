package modules

import (
	"github.com/Aegon-n/sentinel-bot/handler/messages"
	"github.com/Aegon-n/sentinel-bot/handler/buttons"
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

	case "ETH-Windows-Module00":
		handleETHModuleList(bot, queryID, chatID, msgID,"10","Windows", messages.EthWinListOfModulesMsg)

	case "ETH-Windows-Module10":
		handleETHModule(bot, queryID, chatID, msgID,"20","Windows", messages.EthWindowsModule10)

	case "ETH-Windows-Module20":
		handleETHModule(bot, queryID, chatID, msgID, "30","Windows", messages.EthWindowsModule20)

	case "ETH-Windows-Module30":
		handleETHModule(bot, queryID, chatID, msgID, "40","Windows", messages.EthWindowsModule30)

	case "ETH-Windows-Module40":
		handleETHModule(bot, queryID, chatID, msgID, "50","Windows", messages.EthWindowsModule40)

	case "ETH-Windows-Module50":
		handleETHModule(bot, queryID, chatID, msgID, "60","Windows", messages.EthWindowsModule50)

	case "ETH-Windows-Module60":
		handleETHModule(bot, queryID, chatID, msgID, "70","Windows", messages.EthWindowsModule60)

	}
}

func handleETHLinuxModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {

	case "ETH-Linux-Module00":
		handleETHModuleList(bot, queryID, chatID, msgID,"10","Linux", messages.EthWinListOfModulesMsg)

	case "ETH-Linux-Module10":
		handleETHModule(bot, queryID, chatID, msgID,"20","Linux", messages.EthWindowsModule10)

	case "ETH-Linux-Module20":
		handleETHModule(bot, queryID, chatID, msgID, "30","Linux", messages.EthWindowsModule20)

	case "ETH-Linux-Module30":
		handleETHModule(bot, queryID, chatID, msgID, "40","Linux", messages.EthWindowsModule30)

	case "ETH-Linux-Module40":
		handleETHModule(bot, queryID, chatID, msgID, "50","Linux", messages.EthWindowsModule40)

	case "ETH-Linux-Module50":
		handleETHModule(bot, queryID, chatID, msgID, "60","Linux", messages.EthWindowsModule50)

	case "ETH-Linux-Module60":
		handleETHModule(bot, queryID, chatID, msgID, "70","Linux", messages.EthWindowsModule60)

	}
}

func handleETHMacModules(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	module := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	switch module {

	case "ETH-Mac-Module00":
		handleETHModuleList(bot, queryID, chatID, msgID,"10","Mac", messages.EthWinListOfModulesMsg)

	case "ETH-Mac-Module10":
		handleETHModule(bot, queryID, chatID, msgID,"20","Mac", messages.EthWindowsModule10)

	case "ETH-Mac-Module20":
		handleETHModule(bot, queryID, chatID, msgID, "30","Mac", messages.EthWindowsModule20)

	case "ETH-Mac-Module30":
		handleETHModule(bot, queryID, chatID, msgID, "40","Mac", messages.EthWindowsModule30)

	case "ETH-Mac-Module40":
		handleETHModule(bot, queryID, chatID, msgID, "50","Mac", messages.EthWindowsModule40)

	case "ETH-Mac-Module50":
		handleETHModule(bot, queryID, chatID, msgID, "60","Mac", messages.EthWindowsModule50)

	case "ETH-Mac-Module60":
		handleETHModule(bot, queryID, chatID, msgID, "70","Mac", messages.EthWindowsModule60)

	}
}
func handleETHModuleList(Bot *tgbotapi.BotAPI, queryID string, chatID int64, msgID int, next, platform, txt string){
	answeredCallback(Bot, queryID)
	msg := tgbotapi.NewEditMessageText(chatID, msgID, txt)
	btn := tgbotapi.EditMessageReplyMarkupConfig{}
	if platform == "Linux"{
		btn = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.ModulesListButton("ETH-Linux-Module"+next))
	}
	if platform == "Windows" {
		btn = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.ModulesListButton("ETH-Windows-Module"+next))
	}
	if platform == "Mac" {
		btn = tgbotapi.NewEditMessageReplyMarkup(chatID,
			msgID, buttons.ModulesListButton("ETH-Mac-Module"+next))
	}
	Bot.Send(msg)
	Bot.Send(btn)
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

	Bot.Send(msg)
	Bot.Send(btns)

}

func answeredCallback(Bot *tgbotapi.BotAPI, queryId string){
	config := tgbotapi.CallbackConfig{queryId,"",false,"",0}
	Bot.AnswerCallbackQuery(config)
}