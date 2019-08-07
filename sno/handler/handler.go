package handler

import (
	"github.com/Aegon-n/sentinel-bot/sno/buttons"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
	"github.com/Aegon-n/sentinel-bot/sno/helper"
	"github.com/Aegon-n/sentinel-bot/sno/messages"
)
func HandleHome(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	username := helper.GetUserName(u)
	txt := fmt.Sprintf(messages.SNOHomeMsg, username)+"\n\n"+messages.ChooseOption
	btns := buttons.GetButtons("SNOButtons")
	helper.Send(b, u, txt, &btns)
	return
}
func HandleDownloads(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	module := strings.Split(u.CallbackQuery.Data, "-")

	switch module[1] {
		case "Home":
			log.Println("in downloads home")
			go DownloadsHome(b, u)
			return
		case "Desktop":
			log.Println("in Desktop Downloads")
			go DesktopDownloads(b, u)
			return
		case "Mobile":
			log.Println("in mobile Downloads")
			go SendDownloadLink(b, u, "Mobile")
			return
		case "Linux":
			log.Println("in Linux downloads")
			go SendDownloadLink(b, u, "Linux")
      return 
		case "Windows":
			log.Println("in Windows downloads")
			go SendDownloadLink(b, u, "Windows")
			return
		case "MacOS":
			log.Println("in MacOS downloads")
			go SendDownloadLink(b, u, "MacOS")
			return
		default:
			helper.Send(b, u, "Not Implemented")
	}
	return
}

func HandleGuides(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	module := strings.Split(u.CallbackQuery.Data, "-")

	switch module[1] {
	  case "Home":
	  	log.Println("in Home guide")
	  	go GuidesHome(b, u)
		case "dVPN":
			log.Println("in dVPN guide")
			go SendGuide(b, u, "dVPN")
			return
		case "Hub":
			log.Println("in Hub guide")
			go SendGuide(b, u, "Hub")
			return
 }
}

func DownloadsHome(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	// username := helper.GetUserName(u)
	txt := messages.DownloadsHomeMsg+"\n\n"+messages.ChooseVersion
	btns := buttons.GetButtons("DownloadsButtonsList")
	helper.Send(b, u, txt, &btns)
	return
}

func GuidesHome(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	// username := helper.GetUserName(u)
	txt := messages.GuidesHomeMsg+"\n\n"+messages.ChooseOption
	btns := buttons.GetButtons("GuidesButtonsList")
	helper.Send(b, u, txt, &btns)
	return
}

func DesktopDownloads(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	// username := helper.GetUserName(u)
	txt := messages.DownloadsHomeMsg+"\n\n"+messages.ChooseVersion
	btns := buttons.GetButtons("DesktopDownloadsButtonsList")
	helper.Send(b, u, txt, &btns)
	return
}

func SendDownloadLink(b *tgbotapi.BotAPI, u tgbotapi.Update, version string) {
	log.Println(version)
	if version == "Mobile" {
		btns := buttons.GetButtons("MobileFlowEndButtons")
		helper.Send(b, u, messages.MobileDownMsg, &btns)
		
	} else if version == "Linux" {
		btns := buttons.GetButtons("DesktopFlowEndButtons")
		helper.Send(b, u, messages.LinuxDownMsg, &btns)
		
	}else if version == "Windows" {
		btns := buttons.GetButtons("DesktopFlowEndButtons")
		helper.Send(b, u, messages.WinDownMsg, &btns)
		
	}else if version == "MacOS" {
		btns := buttons.GetButtons("DesktopFlowEndButtons")
		helper.Send(b, u, messages.MacOSDownMsg, &btns)
	}
	return
}

func SendGuide(b *tgbotapi.BotAPI, u tgbotapi.Update, guide string) {
	log.Println(guide)
	btns := buttons.GetButtons("GuideFlowEndButtons")
	if guide == "dVPN" {
		helper.Send(b, u, messages.DVPNGuide, &btns)
	} else if guide == "Hub" {
		helper.Send(b, u, messages.HubGuide, &btns)
	}
	return
}
func AboutSentinel(b *tgbotapi.BotAPI, u tgbotapi.Update) {
	log.Println("In about")
	chatID := helper.GetchatID(u)
	btns := buttons.GetButtons("AboutButtons")
	if u.CallbackQuery != nil {
		msgID := u.CallbackQuery.Message.MessageID
		msg := tgbotapi.NewEditMessageText(chatID, msgID, messages.AboutSentinel)
		msg.ReplyMarkup = &btns
		msg.ParseMode = tgbotapi.ModeMarkdown
		b.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatID, messages.AboutSentinel)
	msg.ReplyMarkup = btns
	msg.ParseMode = tgbotapi.ModeMarkdown
	b.Send(msg)
	return
}