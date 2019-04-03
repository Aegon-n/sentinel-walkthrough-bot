package socks5_proxy

import (
	"github.com/Aegon-n/sentinel-bot/handler/buttons"
	"github.com/Aegon-n/sentinel-bot/handler/messages/en_messages"
	"gopkg.in/telegram-bot-api.v4"
	"strings"
)

func HandleSocks5Proxy(bot *tgbotapi.BotAPI, update *tgbotapi.Update ) {
	module := strings.Split(update.CallbackQuery.Data,"-")[2]
	switch module {
	case "Eth":
		handleSocksEth(bot, update)
	case "TM":
		handleSocksTM(bot, update)
	case "10Days","20Days","30Days":

	case "1","2","3":


	}
}
func HandleSocks5InputMsg(Bot *tgbotapi.BotAPI, update *tgbotapi.Update)  {
	input := update.Message.Text
	switch input {
	case isWalletAddress(input):

	case isTxhash(input):


	}
}
func handleSocksEth(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	msg := tgbotapi.NewEditMessageText(chatID, msgID, en_messages.Socks5EthereumMsg)
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.GetButtons("SocksNetworkButtonList"))
	Bot.Send(msg)
	Bot.Send(btns)
}
func handleSocksTM(Bot *tgbotapi.BotAPI, update *tgbotapi.Update)
