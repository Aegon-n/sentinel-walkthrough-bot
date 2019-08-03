package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo/ldb"
	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo/models"
	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/helpers"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/constants"
	"github.com/Aegon-n/sentinel-bot/handler/buttons"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/templates"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func HandleSPS(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	username := helpers.GetUserName(u)
	ChatID := helpers.GetchatID(u)
	greet := fmt.Sprintf(templates.GreetingMsg, username)
	opts := buttons.GetButtons("SpsButtonsList")
	msg := tgbotapi.NewMessage(ChatID, greet+"\n\n"+"Choose an option from the list below: ")
	msg.ReplyMarkup = opts
	b.Send(msg)
	return
}
func HandleSocks5Proxy(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	username := helpers.GetUserName(u)
	ChatID := helpers.GetchatID(u)
	greet := fmt.Sprintf(templates.GreetingMsg, username)
	helpers.Send(b, u, greet)
	_, err := db.Read(constants.AssignedNodeURI, username)
	if err == nil {
		helpers.Send(b, u, templates.NodeAttachedAlready)
		return
	}
	nodes, err := helpers.GetNodes()
	if err != nil {
		helpers.Send(b, u, templates.NoEthNodes)
		return
	}
	err = db.Insert("ChatID", username, strconv.FormatInt(ChatID, 10))
	if err != nil {
		helpers.Send(b, u, templates.Error)
	}
	txt := ""
	for idx, node := range nodes {
		txt = txt + fmt.Sprintf(templates.NodeList, strconv.Itoa(idx+1), node.Location.City, node.Location.Country,
			node.NetSpeed.Download/float64(1000000), node.Load.CPU, "%")

		txt += "\n\n"
		if idx == 60 {
			helpers.Send(b, u, txt)
			txt = ""
		}
	}
	fmt.Println(txt)
	// msg.ReplyMarkup = buttons.GetNodeListButtons(len(nodes))
	db.SetStatus(username, "gotnodelist")
	helpers.Send(b, u, txt)
	msg := tgbotapi.NewMessage(ChatID, templates.AskToSelectANode)
	numericKeyboard := helpers.GetNumaricKeyBoard(len(nodes))
	msg.ReplyMarkup = numericKeyboard
	b.Send(msg)
	return

}

func isNodeID(u tgbotapi.Update) bool {
	_, err := strconv.Atoi(u.Message.Text)
	if err != nil {
		return false
	}
	return true
}

func Socks5InputHandler(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	nodes, err := helpers.GetNodes()
	if err != nil {
		helpers.Send(b, u, templates.NoEthNodes)
		return
	}
	if !u.Message.IsCommand() {
		log.Println("here...")
		if isNodeID(u) {
			go HandleNodeId(b, u, db, nodes)
			return
		}
		helpers.Send(b, u, templates.InvalidOption)
	}
}

func ShowMyNode(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	username := helpers.GetUserName(u)
	status, err := db.GetStatus(username)
	if err != nil {
		helpers.Send(b, u, templates.NoAssignedNodes)
		return
	}
	if status == constants.AssignedNodeURI {
		kv, err := db.Read(constants.AssignedNodeURI, username)
		if err != nil {
			helpers.Send(b, u, templates.NoAssignedNodes)
			return
		}
		ShowMyInfo(b, u, db)
		btnOpts := []models.InlineButtonOptions{
			{Label: "Connect", URL: kv.Value},
		}
		opts := models.ButtonHelper{
			Type: constants.InlineButton, InlineKeyboardOpts: btnOpts,
		}

		helpers.Send(b, u, templates.ConnectMessage, opts)
		return
	}
	helpers.Send(b, u, templates.NoAssignedNodes)
	return
}

func HandleNodeId(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, nodes []models.List) {
	log.Println("came here")
	status, err := db.GetStatus(u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	if status != "gotnodelist" {
		helpers.Send(b, u, "invalid command")
		return
	}
	NodeId := u.Message.Text
	idx, _ := strconv.Atoi(NodeId)
	if idx > len(nodes) {
		helpers.Send(b, u, "Invalid Option!! Please choose correct NodeID")
		return
	}

	txt := fmt.Sprintf(templates.NodeList, strconv.Itoa(idx), nodes[idx-1].Location.City, nodes[idx-1].Location.Country,
		nodes[idx-1].NetSpeed.Download/float64(1000000), nodes[idx-1].Load.CPU, "%")
	err = db.Insert("NodeInfo", u.Message.From.UserName, txt)
	if err != nil {
		log.Println("Error inserting NodeInfo")
		helpers.Send(b, u, "interal bot error")
		return
	}
	_ = db.Insert("NodeIP", u.Message.From.UserName, nodes[idx-1].IP)
	helpers.Send(b, u, txt)
	helpers.Send(b, u, "Please wait... \nGetting socks5 proxy... ")
	go helpers.SocksProxy(b, u, db, nodes[idx-1].AccountAddr)
	return
}
func Restart(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {

	// node, _ := db.Read("NodeIP", u.Message.From.UserName)
	// helpers.DisconnectNode(b, u, node.Value, kv.Value)
	err := db.RemoveUser(u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}

	go HandleSocks5Proxy(b, u, db)
	return
}
func ShowMyInfo(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	username := helpers.GetUserName(u)
	token, err := db.Read("TOKEN", username)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	NodeIP, err := db.Read("NodeIP", username)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
	usage, err := helpers.GetDataUsage(username, NodeIP.Value, token.Value)
	if err != nil {
		log.Println(err)
		helpers.Send(b, u, templates.Error)
		return
	}
	NodeInfo := ""
	nodes, _  := helpers.GetNodes()
	for _, node := range nodes {
		if node.IP == NodeIP.Value {
			NodeInfo = fmt.Sprintf(templates.NodeList, "1", node.Location.City, node.Location.Country, node.NetSpeed.Download/float64(1000000), node.Load.CPU, "%")
			break
		}
	}
	txt := fmt.Sprintf(templates.DATACONSUMPTION, usage.Down/float64(1048576))
	helpers.Send(b, u, txt)
	helpers.Send(b, u, NodeInfo)
	return
}
func DisconnectProxy(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	username := helpers.GetUserName(u)
	status, err := db.GetStatus(username)
	if err != nil {
		log.Println(err)
		helpers.Send(b, u, templates.NoAssignedNodes)
		return
	}
	if status != constants.AssignedNodeURI {
		helpers.Send(b, u, templates.NoAssignedNodes)
		return
	} 
	node, _ := db.Read("NodeIP", username)
	token,_ := db.Read("TOKEN", username)
	helpers.Send(b, u, templates.DisableProxy)
	helpers.DisconnectNode(b, username, node.Value, token.Value)
	err = db.RemoveUser(username)
	if err != nil {
		helpers.Send(b, u, templates.Error)
		return
	}
}

func AnsweredQuery(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	queryId := u.CallbackQuery.ID
	config := tgbotapi.CallbackConfig{queryId, "", false, "", 0}
	bot.AnswerCallbackQuery(config)
}
