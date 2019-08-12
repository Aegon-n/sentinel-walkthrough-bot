package services

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func Twitter_updates(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	chatID := update.CallbackQuery.Message.Chat.ID

	config := oauth1.NewConfig("ae1P8eyT1IvX7zgoalLZNFYKO", "9m2z8BswbMf66gAnz93pSIWsOfFhAHzkltaktV3XZZxB7ACx40")
	token := oauth1.NewToken("1111215341424861184-Z0QLOM2aokMzH4DEKWNqRXPr0hxld8", "Vg8wsTHjQOTyJnGdxKXAdtg7kuM9jsBaKvnTVwkVkk3ax")

	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	client := twitter.NewClient(httpClient)

	client.Streams.User(&twitter.StreamUserParams{})
	tt, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{ScreenName: "sentinel_co", Count: 3, TweetMode: "extended"})
	if err != nil {
		log.Println("can't get the twits", err)
	}

	threetweets := [3]int{0, 1, 2}

	for _, i := range threetweets {

		//ind := len(tt[i].Entities.Urls)

		msg := tgbotapi.NewMessage(chatID, "@"+tt[i].User.ScreenName+" #tweeted on "+tt[i].CreatedAt+"\n"+tt[i].FullText)

		msg.ParseMode = tgbotapi.ModeHTML
		bot.Send(msg)
	}
	msg2 := tgbotapi.NewMessage(chatID, "click /updates to see updates menu\nclick /start to see home menu")
	bot.Send(msg2)
}
func ReadPosts(stream *twitter.Stream) {
	forever := make(chan bool)
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	demux.HandleChan(stream.Messages)

	log.Printf(" [*] Waiting for messages.")
	<-forever
}
