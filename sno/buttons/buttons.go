package buttons

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"sort"

	"github.com/Aegon-n/sentinel-bot/sno/constants"
	"gopkg.in/telegram-bot-api.v4"
)

type B struct {
	Text string
	Data string
}

type Buttons struct {
	SNOButtons                  []map[string]string `json: "SNOButtons"`
	DownloadsButtonsList        []map[string]string `json: "DownloadsButtonsList"`
	DesktopDownloadsButtonsList []map[string]string `json: "DesktopDownloadsButtonsList"`
	GuidesButtonsList           []map[string]string `json: "GuidesButtonsList"`
	DVPNStatsButtonsList        []map[string]string `json: "DVPNStatsButtonsList"`
	AboutButtons                []map[string]string `json: "AboutButtons"`
	DesktopFlowEndButtons       []map[string]string `json: "DesktopFlowEndButtons"`
	MobileFlowEndButtons        []map[string]string `json: "MobileFlowEndButtons"`
	GuideFlowEndButtons         []map[string]string `json: "GuideFlowEndButtons"`
	StatsFlowEndButtons         []map[string]string `json: "StatsFlowEndButtons"`
	UpdatesHomeBtns             []map[string]string `json: "UpdatesHomeBtns"`
	UpdatesFlowEndBtns          []map[string]string `json: "UpdatesFlowEndBtns"`
}

var ButtonList Buttons

func init() {
	data, err := ioutil.ReadFile(constants.ButtonsFilePath)
	if err != nil {
		log.Fatal("File Not Found")
	}
	json.Unmarshal(data, &ButtonList)
	log.Println("Json File loaded")
	log.Println(ButtonList)
}

func GetButtons(BtnType string) tgbotapi.InlineKeyboardMarkup {
	btnlist := getField(&ButtonList, BtnType).Interface().([]map[string]string)
	list := make([][]tgbotapi.InlineKeyboardButton, int(len(btnlist)))
	for idx, b := range btnlist {
		var keys []string
		for k, _ := range b {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			list[idx] = append(list[idx], tgbotapi.NewInlineKeyboardButtonData(k, b[k]))
		}
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: list}
}

func getField(BtnList *Buttons, field string) reflect.Value {
	r := reflect.ValueOf(BtnList)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}
