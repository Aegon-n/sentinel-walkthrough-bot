package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Aegon-n/sentinel-bot/updates/messages"
	"gopkg.in/telegram-bot-api.v4"
)

type RedditPost struct {
	Kind string `json:"kind"`
	Data Data   `json:"data"`
}

type Post struct {
	ApprovedAtUtc  interface{} `json:"approved_at_utc"`
	Selftext       string      `json:"selftext"`
	Subredit       string      `json:"subreddit_name_prefixed"`
	AuthorFullname string      `json:"author_fullname"`
	Clicked        bool        `json:"clicked"`
	Title          string      `json:"title"`
	Name           string      `json:"name"`
	SelftextHTML   string      `json:"selftext_html"`
	Likes          interface{} `json:"likes"`
	Author         string      `json:"author"`
	URL            string      `json:"url"`
	CreatedUTC     float64     `json:"created_utc"`
	IsVideo        bool        `json:"is_video"`
}
type Children struct {
	Kind string `json:"kind"`
	Data Post   `json:"data"`
}
type Data struct {
	Children []Children `json:"children"`
}

func Reddit_updates(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	Api_url := "https://www.reddit.com/r/SENT/new.json?limit=3"
	var body RedditPost

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
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Println("Error")
	}
	defer resp.Body.Close()

	for _, child := range body.Data.Children {

		text := child.Data.Selftext
		urls := child.Data.URL
		date := time.Unix(int64(child.Data.CreatedUTC), 0)
		p_date := date.Format(time.RFC1123)
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(messages.RedditPost, p_date, text, urls))
		msg.ParseMode = tgbotapi.ModeMarkdown
		bot.Send(msg)

	}
	msg2 := tgbotapi.NewMessage(chatID, "click /updates to see updates menu\nclick /start to see home menu")
	bot.Send(msg2)
}
