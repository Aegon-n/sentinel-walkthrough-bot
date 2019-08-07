package main

import (
	"os"
	"strings"

	"github.com/Aegon-n/sentinel-bot/handler"
	"github.com/Aegon-n/sentinel-bot/handler/modules"

	// updates2 "github.com/Aegon-n/sentinel-bot/handler/updates"
	dbo2 "github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo"
	eth_handlers "github.com/Aegon-n/sentinel-bot/eth-socks-proxy/handler"
	eth_helpers "github.com/Aegon-n/sentinel-bot/eth-socks-proxy/helpers"
	"github.com/Aegon-n/sentinel-bot/locale"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo"
	/* "github.com/Aegon-n/sentinel-bot/socks5-proxy/handlers" */
	sno_handler "github.com/Aegon-n/sentinel-bot/sno/handler"
	stats_handler "github.com/Aegon-n/sentinel-bot/dVPN-Stats/handler"

	// tmExplorer "github.com/Aegon-n/sentinel-bot/tm-explorer"
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
		log.Println("h", err)
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
			/*case "tm":
				tmExplorer.HandleTMExplorer(bot, &update)

			case "updates":
				handler.HandleUpdates(bot, &update)
			*/
			case "sps":
				eth_handlers.HandleSocks5Proxy(bot, update, db2)

			case "mynode":
				eth_handlers.ShowMyNode(bot, update, db2)

			case "restart":
				handler.HandleGreet(bot, &update)

			case "restart_sps":
				eth_handlers.Restart(bot, update, db2)

			case "disconnect_proxy":
				eth_handlers.DisconnectProxy(bot, update, db2)
			case "about":
				sno_handler.AboutSentinel(bot, update)
			case "stats":
				stats_handler.HandleHome(bot, update)
			case "downloads":
				sno_handler.DownloadsHome(bot, update)
			case "guides":
				sno_handler.GuidesHome(bot, update)
			}
		}

		if update.CallbackQuery != nil {
			log.Println(update.CallbackQuery)
			eth_handlers.AnsweredQuery(bot, update)
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

			/*case "Medium":
				//log.Println(update)
				updates2.MediumUpdates(bot, &update)

			case "Reddit":
				updates2.Reddit_updates(bot, &update)

			case "Twitter":
				updates2.Twitter_updates(bot, &update)

			case "Socks5":
				handlers.HandleSocks5InlineButtons(bot, update, db)*/
			case "home":
				handler.HandleGreet(bot, &update)
			case "about":
				sno_handler.AboutSentinel(bot, update)
			case "sps":
				eth_handlers.HandleSPS(bot, update, db2)
			case "list_nodes":
				eth_handlers.HandleSocks5Proxy(bot, update, db2)
			case "my_node":
				eth_handlers.ShowMyNode(bot, update, db2)
			case "sno":
				sno_handler.HandleHome(bot, update)
			case "Downloads":
				sno_handler.HandleDownloads(bot, update)
			case "Guides":
				sno_handler.HandleGuides(bot, update)
			case "dVPNStats":
				stats_handler.HandleStats(bot, update)
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
