package tendermint

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/Aegon-n/sentinel-bot/handler/buttons"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/constants"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/ldb"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/models"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/helpers"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/services/proxy"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/services/tendermint/validations"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/templates"
	"github.com/ethereum/go-ethereum/common"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func AskForTendermintWallet(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, nodes []models.List) {
	if len(nodes) == 0 {
		msg := tgbotapi.NewEditMessageText(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Message.MessageID, templates.NoTMNodes)
		btns := tgbotapi.NewEditMessageReplyMarkup(u.CallbackQuery.Message.Chat.ID,
			u.CallbackQuery.Message.MessageID, buttons.GetButtons("SocksNetworkButtonList"))
		b.Send(msg)
		b.Send(btns)
		return
	}

	err := db.Insert(constants.BlockchainNetwork, u.CallbackQuery.From.UserName, constants.TenderMintNetwork)
	if err != nil {
		helpers.Send(b, u, "internal bot error")
		return
	}

	helpers.SetState(b, u, constants.TMState, constants.TMState1, db)
	helpers.Send(b, u, templates.AskForTMWallet)
}

func IsValidTMAccount(u tgbotapi.Update) string {
	ok := strings.HasPrefix(u.Message.Text, constants.TMPrefix)
	l := len(u.Message.Text)

	if ok && l == constants.TMWalletLength {
		return u.Message.Text
	}

	return ""
}

func IsTMTxnHash(u tgbotapi.Update) string {
	ok := common.IsHexAddress(u.Message.Text)
	// sample tx hash = 158AAFD03A6493B922216A7F5AAC8FA0865F7643
	if ok && len(u.Message.Text) == constants.TMHashLength {
		return u.Message.Text
	}

	return ""
}

func getTMTxn(hash string) (models.TMTxn, bool) {
	var txnResp models.TMTxn
	url := fmt.Sprintf(constants.TMTxnURL, hash)
	resp, err := http.Get(url)
	if err != nil {
		return txnResp, false
	}
	if err = json.NewDecoder(resp.Body).Decode(&txnResp); err != nil {
		return txnResp, false
	}

	return txnResp, true
}

func HandleTMTxnHash(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, nodes []models.List) {
	state := helpers.GetState(b, u, constants.TMState, db)
	//color.Green("******* STATE BW = %d *******", state)

	if state < constants.TMState4 {
		helpers.Send(b, u, templates.FollowSequence)
		return
	}

	resp, err := db.Read(constants.NodeTM, u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}

	strToInt, err := strconv.Atoi(resp.Value)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}

	i := strToInt - 1
	if IsValidTMTxn(b, u, db) {
		url := fmt.Sprintf(constants.ProxyURL, nodes[i].IP, strconv.Itoa(3000), "Sentinel", "Password")

		values := []models.KV{
			{Key: constants.IPAddrTM, Value: nodes[i].IP},
			{Key: constants.AssignedNodeURITM, Value: url},
			{Key: constants.IsAuthTM, Value: "true"},
		}
		err := db.MultiWriter(values, u.Message.From.UserName)
		if err != nil {
			helpers.Send(b, u, templates.Error)
			return
		}

		helpers.Send(b, u, "Thanks for submitting the TX-HASH. We're validating it")

		pair, err := db.Read(constants.WalletTM, u.Message.From.UserName)
		if err != nil {
			helpers.Send(b, u, templates.Error)
			return
		}

		tl, err := db.Read(constants.TMTimeLimit, u.Message.From.UserName)
		if err != nil {
			helpers.Send(b, u, templates.Error)
			return
		}

		//timeLimit, _ := time.Parse(time.RFC3339, tl.Value)
		checkInTime := validations.CheckTXNTimeStamp(u.Message.Text, pair.Value, tl.Value)

		if !checkInTime {
			helpers.Send(b, u, "time is up for submitting the transaction")
			return
		}

		helpers.Send(b, u, "creating new user for "+u.Message.From.UserName+"...")
		node := nodes[i]
		err = proxy.AddUser(node.IP, u.Message.From.UserName, constants.PasswordTM, db)
		if err != nil {
			log.Println("here1")
			helpers.Send(b, u, templates.Error)
			return
		}
		pass, err := db.Read(constants.PasswordTM, u.Message.From.UserName)
		if err != nil {
			log.Println("here2")
			helpers.Send(b, u, templates.Error)
			return
		}
		url = fmt.Sprintf(constants.ProxyURL, nodes[i].IP, strconv.Itoa(3000), u.Message.From.UserName, pass.Value)

		kv := []models.KV{
			{
				Key:   constants.IPAddrTM,
				Value: nodes[i].IP,
			},
			{
				Key:   constants.AssignedNodeURITM,
				Value: url,
			},
		}

		err = db.MultiWriter(kv, u.Message.From.UserName)
		if err != nil {
			log.Println("here3")
			helpers.Send(b, u, templates.Error)
			return
		}
		btnOpts := []models.InlineButtonOptions{
			{Label: "Sentinel Proxy", URL: url},
		}
		opts := models.ButtonHelper{
			Type:               constants.InlineButton,
			InlineKeyboardOpts: btnOpts,
		}
		helpers.Send(b, u, templates.Success, opts)
		helpers.SetState(b, u, constants.TMState, constants.TMState5, db)
		return
	}

	helpers.Send(b, u, "invalid txn hash. please try again")
}

func IsValidTMTxn(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) bool {

	username := u.Message.From.UserName
	hash := u.Message.Text
	txn, ok := getTMTxn(hash)

	if !ok {
		helpers.Send(b, u, templates.TXNNotFound)
		return false
	}

	userWallet, err := db.Read(constants.WalletTM, username)

	if err != nil {
		return false
	}

	recipientWallet, err := db.Read(constants.NodeWalletTM, username)

	if err != nil {
		return false
	}

	amount, err := db.Read(constants.NodePriceTM, username)
	if err != nil {
		return false
	}

	if len(txn.Tx.Value.Msg) > 0 {
		okWallet := txn.Tx.Value.Msg[0].Value.From == userWallet.Value
		okRecipient := txn.Tx.Value.Msg[0].Value.To == recipientWallet.Value
		okAmount := parseTxnAmount(txn.Tx.Value.Msg[0].Value.Coins[0].Amount) == amount.Value

		if okWallet && okRecipient && okAmount {
			return true
		}
	}
	return false
}

func HandleWallet(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	TMState := helpers.GetState(b, u, constants.TMState, db)
	//color.Green("******* STATE HANDLE WALLET = %d *******", TMState)
	if TMState < constants.TMState1 {
		helpers.Send(b, u, templates.FollowSequence)
		return
	}
	helpers.SetState(b, u, constants.TMState, constants.TMState2, db)

	bal, ok := validations.CheckTMBalance(u.Message.Text)
	if !ok {
		msg := fmt.Sprintf(templates.NoMinBal, bal)
		helpers.Send(b, u, msg)
		return
	}

	if IsValidTMAccount(u) != "" {
		err := db.Insert(constants.WalletTM, u.Message.From.UserName, u.Message.Text)
		if err != nil {
			helpers.Send(b, u, templates.Error)
			return
		}

		unique := validations.IsUniqueWallet(u.Message.Text, u.Message.From.UserName, db)

		if !unique {
			helpers.Send(b, u, templates.NotUniqueWallet)
			return
		}
		helpers.Send(b, u, "Attached Tendermint wallet to user successfully")
		msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Please select bandwidth")
		msg.ReplyMarkup = buttons.GetButtons("BandwidthSelect")
		b.Send(msg)
		return
	}
	helpers.Send(b, u, templates.Error)
	return
}

func HandleBWTM(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, nodes []models.List) {
	bw := strings.Split(u.CallbackQuery.Data, "-")[2][:2]
	log.Println(bw)
	err := db.Insert(constants.BandwidthTM, u.CallbackQuery.From.UserName, bw)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	switch bw + " Days" {
	case constants.TenD:
		helpers.SubscriptionPeriod(b, u, db,
			constants.TenDays, constants.TenderMintNetwork, constants.NodeBasePrice, constants.TenD,
		)
	case constants.OneM:
		helpers.SubscriptionPeriod(b, u, db,
			constants.Month, constants.TenderMintNetwork, constants.NodeMonthPrice, constants.OneM,
		)
	case constants.ThreeM:
		helpers.SubscriptionPeriod(b, u, db,
			constants.ThreeMonths, constants.TenderMintNetwork, constants.NodeThreeMonthPrice, constants.ThreeM,
		)
	}

	helpers.Send(b, u, templates.AskToSelectANode)
	msg := tgbotapi.MessageConfig{}
	for idx, node := range nodes {
		txt := fmt.Sprintf(templates.NodeList, strconv.Itoa(idx+1), node.Location.City, node.Location.Country,
			node.NetSpeed.Download, node.IP, node.VpnType, node.AccountAddr)
		msg = tgbotapi.NewMessage(u.CallbackQuery.Message.Chat.ID, txt)
	}
	msg.ReplyMarkup = buttons.GetNodeListButtons(len(nodes))
	b.Send(msg)
	return
}

func HandleTMNodeID(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, nodes []models.List) {
	NodeId := strings.Split(u.CallbackQuery.Data, "-")[2]
	idx, _ := strconv.Atoi(NodeId)
	log.Println("came here 3")
	if idx > len(nodes) {
		log.Println("this is the Error")
		helpers.Send(b, u, templates.Error)
		return
	}

	values := []models.KV{
		{Key: constants.NodeTM, Value: NodeId},
		{Key: constants.NodeWalletTM, Value: nodes[idx-1].AccountAddr},
	}
	err := db.MultiWriter(values, u.CallbackQuery.From.UserName)
	if err != nil {
		log.Println("this is the Error2")
		helpers.Send(b, u, templates.Error)
		return
	}

	kv, err := db.Read(constants.NodePriceTM, u.CallbackQuery.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	msg := fmt.Sprintf(templates.AskForPayment, kv.Value)

	helpers.Send(b, u, msg)
	helpers.Send(b, u, nodes[idx-1].AccountAddr)
}

func parseTxnAmount(amount string) string {
	f, e := strconv.ParseFloat(amount, 64)
	if e != nil {
		return ""
	}
	return fmt.Sprintf("%0.0f", f*math.Pow(10, -8))
}
