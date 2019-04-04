package handlers

import (
	"fmt"
	"github.com/Aegon-n/sentinel-bot/handler/buttons"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/constants"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/ldb"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/models"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/helpers"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/services/ethereum"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/services/proxy"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/services/tendermint"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/templates"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func HandleSocks5Proxy(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	helpers.SetState(b, u, constants.EthState, constants.EthState0, db)
	helpers.SetState(b, u, constants.TMState, constants.TMState0, db)
	_ = db.Insert(constants.TMTimeLimit, u.Message.From.UserName, time.Now().Format(time.RFC3339))
	greet := fmt.Sprintf(templates.GreetingMsg, u.Message.From.UserName)
	msg := tgbotapi.NewMessage(u.Message.Chat.ID,greet)
	msg.ReplyMarkup = buttons.GetButtons("SocksNetworkButtonList")
	b.Send(msg)
}
func HandleSocks5InlineButtons(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	nodes, err := helpers.GetNodes()
	log.Println(nodes)
	if err != nil {
		helpers.Send(b, u, templates.NoTMNodes)
		return
	}
	module := strings.Split(u.CallbackQuery.Data,"-")[2]
	switch module {
	case "Eth":
		go ethereum.AskForEthWallet(b, u, db, nodes)
		answeredQuery(b, u)
	case "TM":
		go tendermint.AskForTendermintWallet(b, u, db, nodes)
		answeredQuery(b, u)
	case "10 Days","20 Days","30 Days":
		go HandleBW(b, u, db, nodes)
		answeredQuery(b, u)
	case "1","2","3":
		go HandleNodeId(b, u, db, nodes)
		answeredQuery(b, u)
	}
}
func AboutSentinel(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	helpers.Send(b, u, templates.AboutSentinel)
}

func isEthAddr(u tgbotapi.Update) string {
	r, _ := regexp.Compile(constants.EthRegex)
	ok := common.IsHexAddress(u.Message.Text)

	if ok && r.MatchString(u.Message.Text) {
		return u.Message.Text
	}

	return ""
}

func isNodeID(u tgbotapi.Update) string {
	_, err := strconv.Atoi(u.Message.Text)
	if err != nil {
		return ""
	}

	return u.Message.Text
}

func isTxn(u tgbotapi.Update) string {
	_, err := hexutil.Decode(u.Message.Text)
	if err != nil {
		return ""
	}
	return u.Message.Text
}

func Socks5InputHandler(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	nodes, err := helpers.GetNodes()
	if err != nil {
		helpers.Send(b, u, templates.NoTMNodes)
		return
	}
	switch u.Message.Text {

	/*case isEthAddr(u):
		go ethereum.HandleWallet(b, u, db)*/
	case tendermint.IsValidTMAccount(u):
		go tendermint.HandleWallet(b, u, db)
	case isTxn(u):
		go ethereum.HandleTxHash(b, u, db, nodes)
	case tendermint.IsTMTxnHash(u):
		go tendermint.HandleTMTxnHash(b, u, db, nodes)
	default:
		if !u.Message.IsCommand() {
			helpers.Send(b, u, templates.InvalidOption)
		}
	}
}

func ShowEthWallet(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	kv, err := db.Read(constants.WalletTM, u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}

	helpers.Send(b, u, kv.Value)
}

func ShowMyNode(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	kv, err := db.Read(constants.AssignedNodeURITM, u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	btnOpts := []models.InlineButtonOptions{
		{Label: "Proxy Node", URL: kv.Value},
	}
	opts := models.ButtonHelper{
		Type: constants.InlineButton, InlineKeyboardOpts: btnOpts,
	}
	helpers.Send(b, u, templates.ConnectMessage, opts)

}

func Restart(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	kv, err := db.Read(constants.IPAddrTM, u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	err = proxy.DeleteUser(u.Message.From.UserName, kv.Value)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	err = db.RemoveETHUser(u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	err = db.RemoveTMUser(u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	greet := fmt.Sprintf(templates.GreetingMsg, u.Message.From.UserName)
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, greet)
	msg.ReplyMarkup = buttons.GetButtons("SocksNetworkButtonList")
	b.Send(msg)
}

func ShowMyInfo(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	bw, err := db.Read(constants.BandwidthTM, u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	wallet, err := db.Read(constants.WalletTM, u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	ts, err := db.Read(constants.TimestampTM, u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	log.Println(ts)
	log.Println(bw.Value)
	d, _ := time.Parse(time.RFC3339, ts.Value)
	log.Println(d)
	log.Println(math.Ceil(d.Sub(time.Now()).Hours()))
	days := math.Ceil(d.Sub(time.Now()).Hours()) / 24
	msg := fmt.Sprintf(templates.UserInfo, days, wallet.Value)
	helpers.Send(b, u, msg)
}

func HandleNodeId(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, nodes []models.TONNode) {
	log.Println("came here")
	network, err := db.Read(constants.BlockchainNetwork, u.CallbackQuery.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.NoNetworkSelected)
		return
	}
	if network.Value == constants.TenderMintNetwork {
		TMState := helpers.GetState(b, u, constants.TMState, db)
		//color.Green("******* STATE NODE ID = %d *******", TMState)
		if TMState < constants.TMState3 {
			helpers.Send(b, u, templates.FollowSequence)
			return
		}
		log.Println("Came here 2")
		tendermint.HandleTMNodeID(b, u, db, nodes)
		helpers.SetState(b, u, constants.TMState, constants.TMState4, db)
	}

	if network.Value == constants.EthNetwork {
		EthState := helpers.GetState(b, u, constants.EthState, db)
		//color.Green("******* STATE NODE ID = %d *******", EthState)
		if EthState <= constants.EthState1 {
			helpers.Send(b, u, templates.FollowSequence)
			return
		}
		ethereum.HandleNodeID(b, u, db, nodes)
		helpers.SetState(b, u, constants.EthState, constants.EthState3, db)

	}

}

func HandleBW(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, nodes []models.TONNode) {
	network, err := db.Read(constants.BlockchainNetwork, u.CallbackQuery.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.BWAttachmentError)
	}

	if network.Value == constants.TenderMintNetwork {
		state := helpers.GetState(b, u, constants.TMState, db)
		//color.Green("******* STATE HANDLE BW = %d *******", state)

		if state < constants.TMState2 {
			helpers.Send(b, u, templates.FollowSequence)
			return
		}
		tendermint.HandleBWTM(b, u, db, nodes)
		helpers.SetState(b, u, constants.TMState, constants.TMState3, db)
	}

	if network.Value == constants.EthNetwork {
		EthState := helpers.GetState(b, u, constants.EthState, db)
		//color.Green("******* STATE HANDLE BW = %d *******", EthState)
		if EthState <= constants.EthState0 {
			helpers.Send(b, u, templates.FollowSequence)
			return
		}
		ethereum.HandleEthBW(b, u, db, nodes)
		helpers.SetState(b, u, constants.EthState, constants.TMState2, db)
	}

}

func ClaimRefund(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	helpers.Send(b, u, "feature not available yet")
}
func answeredQuery(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	queryId := u.CallbackQuery.ID
	config := tgbotapi.CallbackConfig{queryId,"",false,"",0}
	bot.AnswerCallbackQuery(config)
}