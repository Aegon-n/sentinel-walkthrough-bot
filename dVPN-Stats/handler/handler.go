package handler

import (
	"github.com/Aegon-n/sentinel-bot/sno/buttons"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
	"strconv"
	"github.com/Aegon-n/sentinel-bot/sno/helper"
	"github.com/Aegon-n/sentinel-bot/dVPN-Stats/helpers"
	eth "github.com/Aegon-n/sentinel-bot/eth-socks-proxy/helpers"
	"github.com/Aegon-n/sentinel-bot/dVPN-Stats/messages"
)
func HandleHome(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	username := helper.GetUserName(u)
	txt := fmt.Sprintf(messages.StatsHomeMsg, username)+"\n\n"+messages.ChooseOption
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
	active_nodes, _, err := helpers.GetCount("active", "nodes")
	if err != nil {
		helper.Send(b, u, messages.UnableToGetStats)
		return
	}
	active_sessions, _, err := helpers.GetCount("active", "sessions")
	if err != nil {
		helper.Send(b, u, messages.UnableToGetStats)
		return
	}
	lastday_bandwidth, err := helpers.GetUsedBandwidth("lastday")
	if err != nil {
		helper.Send(b, u, messages.UnableToGetStats)
		return
	}
	total_bandwidth, err := helpers.GetUsedBandwidth("total")
	if err != nil {
		helper.Send(b, u, messages.UnableToGetStats)
		return
	}
	msg := fmt.Sprintf(messages.StatsMsg, active_nodes, active_sessions, lastday_bandwidth/1024.0, total_bandwidth/(1024.0*1024.0))
	helper.Send(b, u, msg)
	return
}

func SendActiveNodes(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	nodes, err := eth.GetNodes()
	if err != nil {
		helper.Send(b, u, "unable get active nodes list")
		return
	}
	txt := "Here it is: *Active dVPN - nodes*\n"
	for idx, node := range nodes {
		txt = txt + fmt.Sprintf(messages.NodeList, strconv.Itoa(idx+1), node.Location.City, node.Location.Country,
			node.NetSpeed.Download/float64(1000000), node.Load.CPU, "%")

		txt += "\n\n"
		if idx%20 == 0 {
			helper.Send(b, u, txt)
			txt = ""
		}
	}
	helper.Send(b, u, txt)
	return
}