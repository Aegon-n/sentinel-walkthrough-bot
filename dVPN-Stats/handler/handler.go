package handler

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Aegon-n/sentinel-bot/dVPN-Stats/helpers"
	"github.com/Aegon-n/sentinel-bot/dVPN-Stats/messages"
	eth "github.com/Aegon-n/sentinel-bot/eth-socks-proxy/helpers"
	"github.com/Aegon-n/sentinel-bot/sno/buttons"
	"github.com/Aegon-n/sentinel-bot/sno/helper"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleHome(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	username := helper.GetUserName(u)
	txt := fmt.Sprintf(messages.StatsHomeMsg, username) + "\n\n" + messages.ChooseOption
	btns := buttons.GetButtons("DVPNStatsButtonsList")
	helper.Send(b, u, txt, &btns)
	return
}
func HandleStats(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	module := strings.Split(u.CallbackQuery.Data, "-")[1]
	switch module {
	case "Home":
		log.Println("in stats home")
		go HandleHome(b, u)
	case "Stats":
		log.Println("in stats")
		go SendStats(b, u)
		return
	case "ActiveNodes":
		log.Println("in active nodes")
		go SendActiveNodes(b, u)
		return
	}
}

func SendStats(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	active_nodes := make(chan int)
	avg_nodes := make(chan int)
	active_sessions := make(chan int)
	avg_sessions := make(chan int)
	lastday_bandwidth := make(chan float64)
	total_bandwidth := make(chan float64)

	go helpers.GetCount("active", "nodes", active_nodes)
	go helpers.GetCount("average", "nodes", avg_nodes)
	go helpers.GetCount("active", "sessions", active_sessions)
	go helpers.GetCount("average", "sessions", avg_sessions)
	go helpers.GetUsedBandwidth("lastday", lastday_bandwidth)
	go helpers.GetUsedBandwidth("total", total_bandwidth)

	/* log.Println(<- active_nodes)
	log.Println(<- active_sessions)
	log.Println(<- lastday_bandwidth)
	log.Println(<- total_bandwidth) */
	conf := tgbotapi.CallbackConfig{CallbackQueryID: u.CallbackQuery.ID,
		Text: "Loading Statistics ...\nPlease wait ..."}
	b.AnswerCallbackQuery(conf)

	msg := fmt.Sprintf(messages.StatsMsg, <-active_nodes, <-avg_nodes,
		<-active_sessions, <-avg_sessions, <-lastday_bandwidth/1024.0, <-total_bandwidth/(1024.0*1024.0))
	btns := buttons.GetButtons("StatsFlowEndButtons")
	helper.Send(b, u, msg, &btns)
	return
}

func SendActiveNodes(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	nodes, err := eth.GetAllNodes()
	socks_nodes, _ := eth.GetNodes()
	if err != nil {
		helper.Send(b, u, "unable get active nodes list")
		return
	}
	nodes = append(nodes, socks_nodes...)
	txt := "Here it is: *Active dVPN - nodes*\n"
	for idx, node := range nodes {
		txt = txt + fmt.Sprintf(messages.NodeList, strconv.Itoa(idx+1), node.Location.City, node.Location.Country,
			node.NetSpeed.Download/float64(1000000), node.Load.CPU, "%")

		txt += "\n\n"
		if idx != 0 && idx%20 == 0 {
			helper.Send(b, u, txt)
			txt = ""
		}
	}
	btns := buttons.GetButtons("StatsFlowEndButtons")
	msg := tgbotapi.NewMessage(u.CallbackQuery.Message.Chat.ID, txt)
	msg.ReplyMarkup = btns
	b.Send(msg)
	return
}
