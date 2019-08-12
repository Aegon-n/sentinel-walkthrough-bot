package post_notifications

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Aegon-n/sentinel-bot/post_notifications/messages"
	"github.com/jasonlvhit/gocron"
	"go.mongodb.org/mongo-driver/mongo"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Dc      string   `xml:"dc,attr"`
	Content string   `xml:"content,attr"`
	Atom    string   `xml:"atom,attr"`
	Version string   `xml:"version,attr"`
	Cc      string   `xml:"cc,attr"`
	Channel struct {
		Text        string `xml:",chardata"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Image struct {
			Text  string `xml:",chardata"`
			URL   string `xml:"url"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"image"`
		Generator     string `xml:"generator"`
		LastBuildDate string `xml:"lastBuildDate"`
		WebMaster     string `xml:"webMaster"`
		Item          []struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
			Guid  struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			Category []string `xml:"category"`
			Creator  string   `xml:"creator"`
			PubDate  string   `xml:"pubDate"`
			Updated  string   `xml:"updated"`
			Encoded  string   `xml:"encoded"`
		} `xml:"item"`
	} `xml:"channel"`
}

var pubDate = float64(time.Now().Unix())

func CheckForNewPublication(bot *tgbotapi.BotAPI, db *mongo.Collection) {
	var body Rss
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://medium.com/feed/@harishmarri551", nil)
	if err != nil {
		log.Println("error")
	}
	req.Header.Set("User-Agent", "Chrome")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	if err := xml.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Println("Error")
	}
	defer resp.Body.Close()

	if len(body.Channel.Item) == 0 {
		fmt.Println("No Medium posts")
		return
	}
	post := body.Channel.Item[0]
	t, _ := time.Parse(time.RFC1123, post.PubDate)
	postPublishedAt := float64(t.Unix())

	if postPublishedAt > pubDate {
		pubDate = postPublishedAt
		fmt.Println("New Publication: ", post.Title)
		link := strings.Split(post.Link, "?")[0]
		txt := fmt.Sprintf(messages.MediumPost, post.Creator, post.Title, link)
		users := GetAllChatIDs(db)
		broadcastMediumPost(bot, users, txt)
	}
}
func broadcastMediumPost(bot *tgbotapi.BotAPI, chatIDs []int64, text string) {
	for _, id := range chatIDs {
		msg := tgbotapi.NewMessage(id, text)
		msg.ParseMode = tgbotapi.ModeMarkdown
		bot.Send(msg)
	}
	return
}

func MediumPostSheduler(bot *tgbotapi.BotAPI, db *mongo.Collection) {
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(CheckForNewPublication, bot, db)
	<-s.Start()
}
