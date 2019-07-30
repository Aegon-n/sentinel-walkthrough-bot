package handler

import (
	"fmt"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/constants"
	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo/ldb"
	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo/models"
	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/helpers"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/templates"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strconv"
)

func HandleSocks5Proxy(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {

	greet := fmt.Sprintf(templates.GreetingMsg, u.Message.From.UserName)
	helpers.Send(b, u, greet)
	_, err := db.Read(constants.AssignedNodeURI, u.Message.From.UserName)
	if err == nil {
		helpers.Send(b, u, templates.NodeAttachedAlready)
		return
	} 
	nodes, err := helpers.GetNodes()
	if err != nil {
		helpers.Send(b, u, templates.NoEthNodes)
		return
	}
	
	
	txt := ""
	for idx, node := range nodes {
		txt = txt + fmt.Sprintf(templates.NodeList, strconv.Itoa(idx+1), node.Location.City, node.Location.Country,
			node.NetSpeed.Download/float64(8000000))
		txt += "\n\n"
		if idx == 60 {
			helpers.Send(b, u, txt)
			txt = ""
		}
	}
	fmt.Println(nodes[0].NetSpeed.Download)
	fmt.Println(nodes[0].NetSpeed.Download/float64(8000000))
	fmt.Println(txt)
	// msg.ReplyMarkup = buttons.GetNodeListButtons(len(nodes))
	db.SetStatus(u.Message.From.UserName, "gotnodelist")
	helpers.Send(b, u, txt)
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, templates.AskToSelectANode)
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
	status, err := db.GetStatus(u.Message.From.UserName)
	if err != nil {
		helpers.Send(b, u, "There are no nodes assigned for you ..Get a node from here /sps")
		return
	}
	if status == constants.AssignedNodeURI {
		kv, err := db.Read(constants.AssignedNodeURI, u.Message.From.UserName)
		if err != nil {
			helpers.Send(b, u, "There are no nodes assigned for you ..Get a node from here /sps")
			return
		}
		btnOpts := []models.InlineButtonOptions{
			{Label: "Proxy Node", URL: kv.Value},
		}
		opts := models.ButtonHelper{
			Type: constants.InlineButton, InlineKeyboardOpts: btnOpts,
		}
		NodeInfo, _ := db.Read("NodeInfo", u.Message.From.UserName)
		helpers.Send(b, u, NodeInfo.Value)
		helpers.Send(b, u, templates.ConnectMessage, opts)
		return
	}
	helpers.Send(b, u, "There are no nodes assigned for you ..Get a node from here /sps")
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
		helpers.Send(b, u, "invalid cmd")
		return
	}
	NodeId := u.Message.Text
	idx, _ := strconv.Atoi(NodeId)
	if idx > len(nodes) {
		helpers.Send(b, u, "Invalid Option!! Please choose correct NodeID")
		return
	}
	log.Println("hereeeee")
	
	txt := fmt.Sprintf(templates.NodeList,strconv.Itoa(idx), nodes[idx-1].Location.City, nodes[idx-1].Location.Country,
														nodes[idx-1].NetSpeed.Download/float64(8000000))
	err = db.Insert("NodeInfo", u.Message.From.UserName, txt)
	if err != nil {
		log.Println("Error inserting NodeInfo")
		helpers.Send(b, u, "interal bot error")
		return
	}			
	_ = db.Insert("NodeIP", u.Message.From.UserName, nodes[idx -1].IP)								 
	helpers.Send(b, u, txt)
	helpers.Send(b, u, "Please wait... \ngetting socks5 proxy... ")
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

	go HandleSocks5Proxy(b, u , db)
	return
}
func ShowMyInfo(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB) {
	token, err := db.Read("TOKEN", u.Message.From.UserName) 
	if err != nil{
		helpers.Send(b, u, templates.Error)
		return
	}
	node, err := db.Read("NodeIP", u.Message.From.UserName) 
	if err != nil{
		helpers.Send(b, u, templates.Error)
		return
	}
	usage, err := helpers.GetDataUsage(u, node.Value, token.Value)
	if err != nil {
		log.Println(err)
		helpers.Send(b, u, templates.Error)
		return
	}
	txt := fmt.Sprintf("Usage:\ndownload: %f Mb\nupload: %f Mb\n", float64(usage.Down)/float64(800000), float64(usage.Up)/float64(800000))
	helpers.Send(b, u, txt)
	return 
}


func answeredQuery(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	queryId := u.CallbackQuery.ID
	config := tgbotapi.CallbackConfig{queryId,"",false,"",0}
	bot.AnswerCallbackQuery(config)
}
