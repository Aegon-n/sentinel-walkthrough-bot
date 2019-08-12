package post_notifications

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Aegon-n/sentinel-bot/post_notifications/messages"
	"github.com/jasonlvhit/gocron"
	"go.mongodb.org/mongo-driver/mongo"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
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

var last_created = float64(time.Now().Unix())

func CheckForNewPost(bot *tgbotapi.BotAPI, db *mongo.Collection) {
	var body RedditPost
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.reddit.com/u/Harish1551/new.json?limit=1", nil)
	if err != nil {
		log.Println("error")
	}
	req.Header.Set("User-Agent", "Chrome")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Println("Error")
	}
	defer resp.Body.Close()

	if len(body.Data.Children) == 0 {
		fmt.Println("no posts")
		return
	}
	post := body.Data.Children[0].Data
	if post.CreatedUTC > last_created {
		last_created = post.CreatedUTC
		fmt.Println("New Publication: ", post.Title, "\n", post.Selftext)
		txt := fmt.Sprintf(messages.RedditPost, post.Subredit, post.Title, post.Selftext, post.URL)
		users := GetAllChatIDs(db)
		fmt.Println(users)
		go broadcastRedditPost(bot, users, txt)

	}
}
func broadcastRedditPost(bot *tgbotapi.BotAPI, chatIDs []int64, text string) {
	for _, id := range chatIDs {
		msg := tgbotapi.NewMessage(id, text)
		msg.ParseMode = tgbotapi.ModeMarkdown
		bot.Send(msg)
	}
	return
}

func RedditPostSheduler(bot *tgbotapi.BotAPI, db *mongo.Collection) {
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(CheckForNewPost, bot, db)
	<-s.Start()

}
