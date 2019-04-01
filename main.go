package main

import (
	"github.com/fatih/color"
	"github.com/Aegon-n/sentinel-bot/handler"
	"github.com/Aegon-n/sentinel-bot/handler/dbo"
	"github.com/Aegon-n/sentinel-bot/handler/modules"
	updates2 "github.com/Aegon-n/sentinel-bot/handler/updates"
	"github.com/Aegon-n/sentinel-bot/tm-explorer"
	"strings"

	"gopkg.in/telegram-bot-api.v4"

	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("828946943:AAF7En0tUxYR6Mw2NfsPKmGzvzVnh1wlv3M")
	if err != nil {
		log.Fatalf("error in instantiating the bot: %v", err)
	}
	dbo.NewDB()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60


	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		color.Red("error while receiving messages: %s", err)
		return
	}

	color.Green("%s", "started the bot successfully")

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			switch(update.Message.Command()) {
			case "walkthrough":
				handler.HandlerWalkThrough(bot, &update)
			case "start":
				handler.HandleGreet(bot, &update)
			case "locale":
				handler.HandleLocalization(bot, &update)
			case "tm":
				tmExplorer.HandleTMExplorer(bot, &update)

			case "updates":
				handler.HandleUpdates(bot, &update)
			}
		}

		if update.CallbackQuery != nil {
			log.Println(update.CallbackQuery)
			module := strings.Split(update.CallbackQuery.Data,"-")
			log.Println(module)

			switch module[0] {
			case "ETH":
				log.Println(update)
				modules.HandleEthModules(bot, &update, module[1])
			case "TM":
				modules.HandleTMModules(bot, &update, module[1])
			case "Mobile":
				modules.HandleMobileModules(bot, &update, module[1])

			case "Medium":
					//log.Println(update)
					updates2.MediumUpdates(bot, &update)

			case "Reddit":
					updates2.Reddit_updates(bot, &update)

			case "Twitter":

				updates2.Twitter_updates(bot, &update)

			default:
				handler.HandleCallbackQuery(bot, &update)
			}
		}
	}

}
