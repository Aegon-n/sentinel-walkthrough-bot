package modules

import (
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Kind 	string 					`json: "kind"`
	Data  map[string]interface{} 	`json: "data"`
}
func Reddit_updates(bot *tgbotapi.BotAPI,update * tgbotapi.Update){
	queryId := update.CallbackQuery.ID
	Api_url := "https://www.reddit.com/r/SENT/new.json?limit=3"

	chatID := update.CallbackQuery.Message.Chat.ID
	//msg := tgbotapi.MessageConfig{}

	client := &http.Client{}

	req, err := http.NewRequest("GET", Api_url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Chrome")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatal("Error while getting response..")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response")
	}

	var result Response
	json.Unmarshal(data, &result)

	config := tgbotapi.CallbackConfig{queryId,"",false,"",0}
	bot.AnswerCallbackQuery(config)

	nums := []int{2, 3, 4}
	for i,_ := range nums{

		urls := result.Data["children"].([]interface{})[i].(map[string]interface{})["data"].(map[string]interface{})["url"].(string)
		msg := tgbotapi.NewMessage(chatID,urls)
		msg.ParseMode = tgbotapi.ModeHTML
		bot.Send(msg)

	}


	HandleGreet(bot,update)

}
