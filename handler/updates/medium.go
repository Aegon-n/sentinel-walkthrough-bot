package updates

import (
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type obj struct {
	ImageUrl string `json: "imageUrl"`
	P string `json:"p"`
	Title string `json:"title"`
	Link string 	`json:"link"`
}


func MediumUpdates(bot *tgbotapi.BotAPI,update *tgbotapi.Update)  {
	queryId := update.CallbackQuery.ID
	chatId := update.CallbackQuery.Message.Chat.ID
	api_url := "http://185.181.8.90:9091/feed";

	resp, err := http.Get(api_url)

	if(err != nil){
		log.Print("access denide or api url is not available",err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading response")
	}


	result := make([]obj,0)
	json.Unmarshal(data, &result)

	config := tgbotapi.CallbackConfig{queryId,"",false,"",0}
	bot.AnswerCallbackQuery(config)

	for i,res := range result{
			if i > 2 {
				break
			}
			resupdates := strings.Split(res.Link, "?")[0]
			msg := tgbotapi.NewMessage(chatId,resupdates)
			msg.ParseMode = tgbotapi.ModeMarkdown
			bot.Send(msg)
	}
	HandleGreet(bot,update)
}
