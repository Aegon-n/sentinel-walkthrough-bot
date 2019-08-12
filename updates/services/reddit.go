package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/telegram-bot-api.v4"
)

type Response struct {
	Kind string                 `json: "kind"`
	Data map[string]interface{} `json: "data"`
}

func Reddit_updates(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
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

	nums := []int{2, 3, 4}
	for i, _ := range nums {

		text := result.Data["children"].([]interface{})[i].(map[string]interface{})["data"].(map[string]interface{})["selftext"].(string)
		urls := result.Data["children"].([]interface{})[i].(map[string]interface{})["data"].(map[string]interface{})["url"].(string)
		msg := tgbotapi.NewMessage(chatID, text+"\n"+urls)
		bot.Send(msg)

	}
	msg2 := tgbotapi.NewMessage(chatID, "click /updates to see updates menu\nclick /start to see home menu")
	bot.Send(msg2)
}
