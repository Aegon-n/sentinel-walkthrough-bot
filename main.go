package main

import (
	"github.com/Aegon-n/sentinel-bot/handler"
	"github.com/Aegon-n/sentinel-bot/handler/modules"
	"github.com/fatih/color"
	"strings"

	"gopkg.in/telegram-bot-api.v4"

	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("774002945:AAEHc1Gc5WfMEVWz4oilLuENzbBL7mH006A")
	if err != nil {
		log.Fatalf("error in instantiating the bot: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		color.Red("error while receiving messages: %s", err)
		return
	}

	color.Green("%s", "started the bot successfully")

	for update := range updates {
		log.Println(update.CallbackQuery)
		if update.Message != nil && update.Message.IsCommand() {
			switch(update.Message.Command()) {
			case "walkthrough":
				handler.HandlerWalkThrough(bot, &update)
			case "start":
				handler.HandleGreet(bot, &update)
			}
		}

		if update.CallbackQuery != nil {
			module := strings.Split(update.CallbackQuery.Data,"-")[0]
			log.Println(module)
			switch module {
			case "ETH":
				modules.HandleEthModules(bot, &update, strings.Split(update.CallbackQuery.Data,"-")[1])
			case "TM":
				modules.HandleTMModules(bot, &update, strings.Split(update.CallbackQuery.Data,"-")[1])
			default:
				handler.HandleCallbackQuery(bot, &update)
			}
		}
	}

}
