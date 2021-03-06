package modules

import (
	"fmt"
	"github.com/Aegon-n/sentinel-bot/handler/messages/en_messages"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)


func Twitter_updates(bot *tgbotapi.BotAPI,update *tgbotapi.Update)  {


	queryId := update.CallbackQuery.ID

	chatID := update.CallbackQuery.Message.Chat.ID

	config := oauth1.NewConfig("ae1P8eyT1IvX7zgoalLZNFYKO", "9m2z8BswbMf66gAnz93pSIWsOfFhAHzkltaktV3XZZxB7ACx40")
	token := oauth1.NewToken("1111215341424861184-Z0QLOM2aokMzH4DEKWNqRXPr0hxld8", "Vg8wsTHjQOTyJnGdxKXAdtg7kuM9jsBaKvnTVwkVkk3ax")
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	client := twitter.NewClient(httpClient)

	config1 := tgbotapi.CallbackConfig{queryId,"",false,"",0}
	bot.AnswerCallbackQuery(config1)


	tt,_,err:= client.Timelines.UserTimeline(&twitter.UserTimelineParams{ScreenName:"sentinel_co",Count:3,})


	if(err != nil){
		log.Println("can't get the twits",err)
	}

	threetweets := [3]int{0,1,2}

	for _,i := range threetweets{

			ind := len(tt[i].Entities.Urls)

			msg := tgbotapi.NewMessage(chatID,tt[i].User.ScreenName + " tweeted on " + tt[i].CreatedAt+"\n"+tt[i].Entities.Urls[ind-1].ExpandedURL)

			msg.ParseMode = tgbotapi.ModeHTML

			bot.Send(msg)


	}
	HandleGreet(bot,update)

}

func HandleGreet(Bot *tgbotapi.BotAPI, update *tgbotapi.Update )  {
	username := update.CallbackQuery.Message.Chat.UserName
	chatID := update.CallbackQuery.Message.Chat.ID
	txt := fmt.Sprintf(en_messages.WelcomeGreetMsg, username)+"\n"+en_messages.SelectwalkthroughMsg
	msg := tgbotapi.NewMessage(chatID,txt)
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)

}
