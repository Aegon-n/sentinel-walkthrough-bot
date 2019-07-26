package ethereum

import (
	"encoding/json"
	"fmt"
	"github.com/Aegon-n/sentinel-bot/handler/buttons"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/helpers"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/constants"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/ldb"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/models"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/services/proxy"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/templates"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleWallet(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	helpers.SetState(b, u, constants.EthState, constants.EthState1, db)

	ok := common.IsHexAddress(u.Message.Text)
	if ok {
		err := db.Insert(constants.EthAddr, u.Message.From.UserName, u.Message.Text)
		if err != nil {
			c := tgbotapi.NewMessage(u.Message.Chat.ID, "error while storing user eth address")
			_, _ = b.Send(c)
			return
		}
		helpers.Send(b, u, "Attached the ETH wallet to user successfully")
		helpers.Send(b, u, "Searching for nodes ... ")
		nodes, err := helpers.GetNodes()
		node := nodes[0]
		fmt.Println(node)
		uri := "https://t.me/socks?server=192.168.2.156&port=1080&user=client1564138446879037&pass=DZtXIT1eQ3"
		// uri, message, err := helpers.SocksProxy(b, u, node.AccountAddr, u.Message.Text) 
		// helpers.Send(b, u, message)
		btnOpts := []models.InlineButtonOptions{
			{Label: "Sentinel Proxy Node", URL: uri},
		}
		opts := models.ButtonHelper{Type: constants.InlineButton, InlineKeyboardOpts: btnOpts}
		helpers.Send(b, u, templates.Success, opts)
		return
		
		/* opts := models.ButtonHelper{
			Type:   constants.ReplyButton,
			Labels: []string{constants.TenD, constants.OneM, constants.ThreeM},
		}
		helpers.Send(b, u, templates.AskForBW, opts)

		err = db.Insert(constants.EthAddr, u.Message.From.UserName, u.Message.Text)
		if err != nil {
			helpers.Send(b, u, "could not store your wallet")
			return
		} */

		return
	}

	helpers.Send(b, u, "internal bot error")
	return
}

func HandleEthBW(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, nodes []models.List) {

	resp, err := db.Read(constants.Bandwidth, u.Message.From.UserName)

	if err != nil {
		err := db.Insert(constants.Bandwidth, u.Message.From.UserName, u.Message.Text[:2])
		if err != nil {
			return
		}
		switch u.Message.Text {
		case constants.TenD:
			helpers.SubscriptionPeriod(b, u, db,
				constants.TenDays, constants.EthNetwork, constants.NodeBasePrice, constants.TenD,
			)
		case constants.OneM:
			helpers.SubscriptionPeriod(b, u, db,
				constants.Month, constants.EthNetwork, constants.NodeMonthPrice, constants.OneM,
			)
		case constants.ThreeM:
			helpers.SubscriptionPeriod(b, u, db,
				constants.ThreeMonths, constants.EthNetwork, constants.NodeThreeMonthPrice, constants.ThreeM,
			)
		}

		helpers.Send(b, u, templates.AskToSelectANode)
		for idx, node := range nodes {
			msg := fmt.Sprintf(templates.NodeList, strconv.Itoa(idx+1), node.Location.City,node.Location.Country,
				node.NetSpeed.Download,node.IP, node.VpnType, node.AccountAddr)
			helpers.Send(b, u, msg)
		}
		return
	}

	nodeIdx, err := strconv.ParseInt(resp.Value[0:2], 10, 64)
	if err != nil {
		c := tgbotapi.NewMessage(u.Message.Chat.ID, err.Error())
		_, _ = b.Send(c)
		return
	}

	var n models.List
	for i := 0; i < len(nodes); i++ {
		if i == int(nodeIdx) {
			n = nodes[i]
			return
		}
	}

	uri := fmt.Sprintf(constants.ProxyURL, n.IP, strconv.Itoa(3000), "sentinel", "Password")
	buttonOptions := []models.InlineButtonOptions{
		{Label: "Sentinel Proxy Node", URL: uri},
	}
	msg := "you have already selected : Node " + resp.Value
	opts := models.ButtonHelper{
		Type:               constants.InlineButton,
		InlineKeyboardOpts: buttonOptions,
	}
	helpers.Send(b, u, msg, opts)
}

func AskForEthWallet(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, nodes []models.List) {
	helpers.SetState(b, u, constants.EthState, constants.EthState0, db)
	if len(nodes) == 0 {
		msg := tgbotapi.NewEditMessageText(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Message.MessageID, templates.NoEthNodes)
		btns := tgbotapi.NewEditMessageReplyMarkup(u.CallbackQuery.Message.Chat.ID,
			u.CallbackQuery.Message.MessageID,buttons.GetButtons("SocksNetworkButtonList"))
		b.Send(msg)
		b.Send(btns)
		return
	}

	err := db.Insert(constants.BlockchainNetwork, u.CallbackQuery.From.UserName, constants.EthNetwork)
	if err != nil {
		helpers.Send(b, u, "internal bot error")
		return
	}

	helpers.Send(b, u, templates.AskForEthWallet)
	return
}

func HandleTxHash(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, nodes []models.List) {
	helpers.SetState(b, u, constants.EthState, constants.EthState4, db)
	resp, err := db.Read(constants.Node, u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	UserWallet, err := db.Read(constants.EthAddr, u.Message.From.UserName)
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
	if FindTxByHash(u.Message.Text, UserWallet.Value, u, db) {
		uri := fmt.Sprintf(constants.ProxyURL, nodes[i].IP, strconv.Itoa(3000), "Sentinel", "Password")
		err := db.Insert(constants.IPAddr, u.Message.From.UserName, nodes[i].IP)
		if err != nil {
			helpers.Send(b, u, templates.Error)
			return
		}
		err = db.Insert(constants.AssignedNodeURI, u.Message.From.UserName, uri)

		if err != nil {
			helpers.Send(b, u, templates.Error)
			return
		}
		err = db.Insert(constants.IsAuth, u.Message.From.UserName, "true")
		if err != nil {
			helpers.Send(b, u, templates.Error)
			return
		}
		helpers.Send(b, u, "Thanks for submitting the TX-HASH. We're validating it")
		helpers.Send(b, u, "creating new user for "+u.Message.From.UserName+"...")

		node := nodes[i]
		err = proxy.AddUser(node.IP, u.Message.From.UserName, constants.Password, db)
		if err != nil {
			helpers.Send(b, u, "Error while creating SOCKS5 user for "+u.Message.From.UserName)
			return
		}
		pass, err := db.Read(constants.Password, u.Message.From.UserName)
		if err != nil {
			helpers.Send(b, u, templates.Error)
			return
		}
		uri = fmt.Sprintf(constants.ProxyURL, node.IP, strconv.Itoa(3000), u.Message.From.UserName, pass.Value)
		err = db.Insert(constants.IPAddr, u.Message.From.UserName, nodes[i].IP)
		if err != nil {
			helpers.Send(b, u, templates.Error)
			return
		}
		err = db.Insert(constants.AssignedNodeURI, u.Message.From.UserName, uri)
		if err != nil {
			helpers.Send(b, u, templates.Error)
			return
		}
		btnOpts := []models.InlineButtonOptions{
			{Label: "Sentinel Proxy Node", URL: uri},
		}
		opts := models.ButtonHelper{Type: constants.InlineButton, InlineKeyboardOpts: btnOpts}
		helpers.Send(b, u, templates.Success, opts)
		return
	}

	helpers.Send(b, u, "invalid transaction hash. Please try again")
}

func FindTxByHash(txHash, walletAddr string, u tgbotapi.Update, db ldb.BotDB) bool {

	wallet := "0x" + constants.ZFill + strings.TrimLeft(walletAddr, "0x")
	uri := constants.TestSentURI1 + wallet + constants.TestSendURI2 + wallet
	resp, err := http.Get(uri)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	var body models.TxReceiptList
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return false
	}

	var val bool
	user, err := db.Read(constants.EthAddr, u.Message.From.UserName)
	if err != nil {
		return false
	}
	node, err := db.Read(constants.NodeWallet, u.Message.From.UserName)
	if err != nil {
		return false
	}
	// userWallet := fmt.Sprintf("%s", w.Value)
	for _, txReceipt := range body.Results {

		if txReceipt.TransactionHash == txHash {
			// nodeWallet := "0xceb5bc384012f0eebee119d82a24925c47714fe3"
			d, e := db.Read(constants.Timestamp, u.Message.From.UserName)
			if e != nil {
				return false
			}
			duration, err := time.Parse(time.RFC3339, fmt.Sprintf("%s", d.Value))
			if err != nil {
				return false
			}
			diff := math.Ceil(duration.Sub(time.Now()).Hours() / 24)

			okWallet := strings.EqualFold(txReceipt.Topics[1], "0x"+constants.ZFill+strings.TrimLeft(user.Value, "0x"))

			okRecipient := strings.EqualFold(txReceipt.Topics[2], "0x"+constants.ZFill+strings.TrimLeft(node.Value, "0x"))
			okAmount := false
			if diff == 10 {
				okAmount = hex2int(txReceipt.Data) == uint64(1000000000)
			} else if diff == 30 {
				okAmount = hex2int(txReceipt.Data) == uint64(3000000000)
			} else {
				okAmount = hex2int(txReceipt.Data) == uint64(8000000000)
			}
			color.Red("comparison: %v%v%v%v", hex2int(txReceipt.Data), okAmount, okRecipient, okWallet)
			if okWallet && okRecipient && okAmount {
				val = true
			}
		}
	}
	return val
}

func hex2int(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return uint64(result)
}

func HandleNodeID(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, nodes []models.List) {
	NodeId := u.Message.Text
	idx, _ := strconv.Atoi(NodeId)
	if idx > len(nodes) {
		helpers.Send(b, u, templates.Error)
		return
	}
	values := []models.KV{
		{Key: constants.Node, Value: NodeId},
		{Key: constants.NodeWallet, Value: nodes[idx-1].AccountAddr},
	}
	err := db.MultiWriter(values, u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}

	kv, err := db.Read(constants.NodePrice, u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}

	msg := fmt.Sprintf(templates.AskForPayment, kv.Value)
	helpers.Send(b, u, msg)
	helpers.Send(b, u, nodes[idx-1].AccountAddr)
}
