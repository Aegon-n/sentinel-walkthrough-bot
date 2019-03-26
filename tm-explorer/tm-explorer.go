package tmExplorer

import (
	"encoding/json"
	"fmt"
	"github.com/Aegon-n/sentinel-bot/tm-explorer/models"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleTMExplorer(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	args := update.Message.CommandArguments()
	chatID := update.Message.Chat.ID
	msg := tgbotapi.MessageConfig{}

	resp, err := http.Get("http://tm-lcd.sentinelgroup.io:26657/status")
	if err != nil {
		log.Fatal("Error while getting response..")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response")
	}

	var result models.Responce
	json.Unmarshal(data, &result)

	if args == "lastblock" || args == "latestblock" {
		latestBlockHight := result.Result["sync_info"].(map[string]interface{})["latest_block_height"]
		txt := fmt.Sprintf("*Latest Block:* %v",latestBlockHight)
		msg = tgbotapi.NewMessage(chatID, txt)
	}
	if args == "validators" || args == "validator" {
		validator_adrr := result.Result["validator_info"].(map[string]interface{})["address"]
		validator_votingpower := result.Result["validator_info"].(map[string]interface{})["voting_power"]
		txt := fmt.Sprintf("*Address:* %v\n*Voting Power:* %v",validator_adrr, validator_votingpower)
		msg = tgbotapi.NewMessage(chatID, txt)
	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
}
