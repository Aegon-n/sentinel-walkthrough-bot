package buttons

import (
	"encoding/json"
	"github.com/Aegon-n/sentinel-bot/handler/constants"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
)

type B struct {
	Text string
	Data string
}
var ButtonList Buttons

type Buttons struct {
	LanguageButtons				[]map[string]string		`json: "LanguageButtons"`
	AppButtonsList 				[]map[string]string  	`json: "AppButtonsList"`
	DesktopOSButtonsList 		[]map[string]string 	`json: "DesktopOSButtonsList"`
	MobileOSButtonsList 		[]map[string]string  	`json: "MobileOSButtonsList"`
	LinuxNetworkButtonList 		[]map[string]string  	`json: "LinuxNetworkButtonList"`
	WindowsNetworkButtonList 	[]map[string]string  	`json: "WindowsNetworkButtonList"`
	MacNetworkButtonList 		[]map[string]string  	`json: "MacNetworkButtonList"`
	LinuxEthModulesButtonList   []map[string]string 	`json: "LinuxEthModulesButtonList"`
	LinuxTMModulesButtonList	[]map[string]string		`json: "LinuxTMModulesButtonList"`
	WindowsEthModulesButtonList	[]map[string]string		`json: "WindowsEthModulesButtonList"`
	WindowsTMModulesButtonList	[]map[string]string		`json: "WindowsTMModulesButtonList"`
	MacEthModulesButtonList		[]map[string]string		`json: "MacEthModulesButtonList"`
	MacTMModulesButtonList		[]map[string]string		`json: "MacTMModulesButtonList"`
	AndroidModulesButtonList	[]map[string]string		`json: "AndroidModulesButtonList"`
	IOSModulesButtonList		[]map[string]string		`json: "IOSModulesButtonList"`

}

func init() {
	data, err := ioutil.ReadFile(constants.ButtonsFilePath)
	if err != nil {
		log.Fatal("File Not Found")
	}
	json.Unmarshal(data, &ButtonList)
	log.Println("Json File loaded")
	log.Println(ButtonList)
}


func PersistentNavButtons(data1, data2, data3 string) tgbotapi.InlineKeyboardMarkup {
	home := tgbotapi.NewInlineKeyboardButtonData("🏠Home","Home")
	btn1 := tgbotapi.NewInlineKeyboardButtonData("◀Prev",data1)
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Skip⏭",data2)
	btn3 := tgbotapi.NewInlineKeyboardButtonData("Next▶",data3)
	btns := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{home, btn1,btn2,btn3}},
	}
	return btns
}

func LastModuleButtons(data1 string) tgbotapi.InlineKeyboardMarkup {
	home := tgbotapi.NewInlineKeyboardButtonData("🏠Home", "Home")
	prev := tgbotapi.NewInlineKeyboardButtonData("◀Prev",data1)
	exit := tgbotapi.NewInlineKeyboardButtonData("❌Exit","Exit")

	btns := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{home,prev,exit}},
	}

	return btns
}

func GetButtons(buttontype string) tgbotapi.InlineKeyboardMarkup {
	list := [][]tgbotapi.InlineKeyboardButton{{},{}}
	if buttontype == "AppButtonsList" {
		return genBtnRow(ButtonList.AppButtonsList)

	}
	if buttontype == "DesktopOSButtonsList" {
		return genBtnRow(ButtonList.DesktopOSButtonsList)

	}
	if buttontype == "MobileOSButtonsList" {
		return genBtnRow(ButtonList.MobileOSButtonsList)

	}
	if buttontype == "LinuxNetworkButtonList" {
		return genBtnRow(ButtonList.LinuxNetworkButtonList)

	}
	if buttontype == "WindowsNetworkButtonList" {
		return genBtnRow(ButtonList.WindowsNetworkButtonList)

	}
	if buttontype == "MacNetworkButtonList" {
		return genBtnRow(ButtonList.MacNetworkButtonList)

	}
	if buttontype == "LinuxEthModulesButtonList" {
		return genModuleButtonList(ButtonList.LinuxEthModulesButtonList)

	}
	if buttontype == "LinuxTMModulesButtonList" {
		return genModuleButtonList(ButtonList.LinuxTMModulesButtonList)

	}
	if buttontype == "WindowsEthModulesButtonList" {
		return genModuleButtonList(ButtonList.WindowsEthModulesButtonList)

	}
	if buttontype == "WindowsTMModulesButtonList" {
		return genModuleButtonList(ButtonList.WindowsTMModulesButtonList)

	}
	if buttontype == "MacEthModulesButtonList" {
		return genModuleButtonList(ButtonList.MacEthModulesButtonList)

	}
	if buttontype == "MacTMModulesButtonList" {
		return genModuleButtonList(ButtonList.MacTMModulesButtonList)

	}
	if buttontype == "AndroidModulesButtonList" {
		return genModuleButtonList(ButtonList.AndroidModulesButtonList)

	}
	if buttontype == "IOSModulesButtonList" {
		return genModuleButtonList(ButtonList.IOSModulesButtonList)

	}
	if buttontype == "LanguageButtons" {
		return genModuleButtonList(ButtonList.LanguageButtons)

	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard:list}
}

func genModuleButtonList(data []map[string]string) tgbotapi.InlineKeyboardMarkup {
	list := [][]tgbotapi.InlineKeyboardButton{{},{}}
	log.Println(data)
	for i, b := range data {
		i+=1
		if i > 2{
			for k,v := range b {
				list[1] = append(list[1], tgbotapi.NewInlineKeyboardButtonData(k, v))
			}
		}else {
			for k,v := range b {
				list[0] = append(list[0], tgbotapi.NewInlineKeyboardButtonData(k, v))
			}
		}

	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: list}
}
func genBtnRow(data []map[string]string) tgbotapi.InlineKeyboardMarkup {
	list := [][]tgbotapi.InlineKeyboardButton{{},{}}
	for _, b := range data {
		for k,v := range b {
			list[0] = append(list[0], tgbotapi.NewInlineKeyboardButtonData(k, v))
		}
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: list}
}