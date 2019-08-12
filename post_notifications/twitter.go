package post_notifications

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Aegon-n/sentinel-bot/handler/models"
	"github.com/Aegon-n/sentinel-bot/post_notifications/messages"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/telegram-bot-api.v4"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func getClient(creds *Credentials) (*twitter.Client, error) {

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)

	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}

func TwitterConfig() *twitter.Stream {
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	fmt.Printf("%+v\n", creds)

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	fmt.Printf("%+v\n", client)
	stream, err := client.Streams.Filter(&twitter.StreamFilterParams{
		Follow: []string{"1867993700"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stream:")
	fmt.Printf("%+v\n", stream)
	return stream
}

func UpdatePosts(bot *tgbotapi.BotAPI, stream *twitter.Stream, collection *mongo.Collection) {
	forever := make(chan bool)
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.IDStr)
		fmt.Println("New Tweet:\n", tweet.User.Name, "\n", tweet.CreatedAt, "\n",
			"\n", tweet.Text)
		var txt string
		txt = fmt.Sprintf(messages.TwitterMsg, tweet.User.Name, tweet.Text,
			tweet.User.ScreenName, tweet.IDStr)
		if len(tweet.Text) > 2 && tweet.Text[:2] == "RT" {
			txt = fmt.Sprintf(messages.RetweetMsg, tweet.User.Name, tweet.Text,
				tweet.User.ScreenName, tweet.IDStr)
		}
		users := GetAllChatIDs(collection)
		broadcastTweet(bot, users, txt)
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	demux.HandleChan(stream.Messages)

	log.Printf(" [*] Waiting for messages.")
	<-forever
}

func GetAllChatIDs(collection *mongo.Collection) []int64 {

	var user_list []int64
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Println(err)
		return []int64{}
	}
	for cur.Next(context.TODO()) {

		var user models.BotUser
		err := cur.Decode(&user)
		if err != nil {
			log.Println(err)
			return []int64{}
		}

		user_list = append(user_list, user.ChatID)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	cur.Close(context.TODO())
	return user_list
}

func broadcastTweet(bot *tgbotapi.BotAPI, chatIDs []int64, text string) {
	for _, id := range chatIDs {
		msg := tgbotapi.NewMessage(id, text)
		msg.ParseMode = tgbotapi.ModeHTML
		bot.Send(msg)
	}
	return
}
