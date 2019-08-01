package main

import (
	"os"
	"strings"

	"github.com/Aegon-n/sentinel-bot/handler"
	"github.com/Aegon-n/sentinel-bot/handler/modules"
	updates2 "github.com/Aegon-n/sentinel-bot/handler/updates"
	"github.com/Aegon-n/sentinel-bot/locale"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo"
	dbo2 "github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/handlers"
	eth_handlers "github.com/Aegon-n/sentinel-bot/eth-socks-proxy/handler"
	eth_helpers "github.com/Aegon-n/sentinel-bot/eth-socks-proxy/helpers"
	tmExplorer "github.com/Aegon-n/sentinel-bot/tm-explorer"
	"github.com/fatih/color"

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
		log.Println(err)
		log.Fatal(err)
	}
	db2, err := dbo2.NewDB()
	if err != nil {
		log.Println("h",err)
		log.Fatal(err)
	}
	//nodes, err := helpers.GetNodes()
	go eth_helpers.ExpiredUsersJob(bot, db2)


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
				eth_handlers.HandleSocks5Proxy(bot, update, db2)

			case "mynode":
				eth_handlers.ShowMyNode(bot, update, db2)

			case "restart":
				handler.HandleGreet(bot, &update)

			case "restart_sps":
				eth_handlers.Restart(bot, update, db2)
				
			case "sps_info":
				eth_handlers.ShowMyInfo(bot, update, db2)
			/*case "sps_wallet":
				handlers.ShowEthWallet(bot, update, db) */
			case "about":
				handlers.AboutSentinel(bot, update)
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
			eth_handlers.Socks5InputHandler(bot, update, db2)
			/* TMState := helpers.GetState(bot, update, constants.TMState, db)
			color.Green("******* APP STATE = %d *******", TMState) */
		}

	}

}
