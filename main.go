package main

import (
	"os"
	"strings"

	"github.com/Aegon-n/sentinel-bot/handler"
	"github.com/Aegon-n/sentinel-bot/handler/modules"
	updates2 "github.com/Aegon-n/sentinel-bot/handler/updates"
	"github.com/Aegon-n/sentinel-bot/locale"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/handlers"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/helpers"
	tmExplorer "github.com/Aegon-n/sentinel-bot/tm-explorer"
	"github.com/fatih/color"
	"github.com/than-os/sentinel-bot/constants"

	tgbotapi "gopkg.in/telegram-bot-api.v4"

	"log"
)

func main() {
	locale.StartLocalizer()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_API_KEY"))
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
	color.Green("started %s successfully", bot.Self.UserName)

	db, err := dbo.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	//nodes, err := helpers.GetNodes()
	//go proxy.UpdateNodesListJob(&nodes)

	for update := range updates {

		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "walkthrough":
				handler.HandlerWalkThrough(bot, &update)
			case "start":
				handler.HandleGreet(bot, &update)
			case "help":
				log.Println("in help")
				handler.HandleHelp(bot, &update)
			/*case "locale":
			handler.HandleLocalization(bot, &update)*/
			case "tm":
				tmExplorer.HandleTMExplorer(bot, &update)

			case "updates":
				handler.HandleUpdates(bot, &update)

			case "sps":
				handlers.HandleSocks5Proxy(bot, update, db)

			case "mynode":
				handlers.ShowMyNode(bot, update, db)

			case "restart_sps":
				handlers.Restart(bot, update, db)
			case "sps_info":
				handlers.ShowMyInfo(bot, update, db)
			case "sps_wallet":
				handlers.ShowEthWallet(bot, update, db)
			}
		}

		if update.CallbackQuery != nil {
			log.Println(update.CallbackQuery)
			module := strings.Split(update.CallbackQuery.Data, "-")
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

			case "Socks5":
				handlers.HandleSocks5InlineButtons(bot, update, db)

			default:
				handler.HandleCallbackQuery(bot, &update, db)
			}
		}
		if update.Message != nil && !update.Message.IsCommand() && len(update.Message.Text) > 0 {
			handlers.Socks5InputHandler(bot, update, db)
			TMState := helpers.GetState(bot, update, constants.TMState, db)
			color.Green("******* APP STATE = %d *******", TMState)
		}

	}

}
