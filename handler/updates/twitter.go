package updates

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/Aegon-n/sentinel-bot/handler/messages/en_messages"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)


func Twitter_updates(bot *tgbotapi.BotAPI,update *tgbotapi.Update)  {


	queryId := update.CallbackQuery.ID

	chatID := update.CallbackQuery.Message.Chat.ID

	config := oauth1.NewConfig("cosumer Api key", "cosumer api secret key")
	token, err := oauth1.NewToken("Access Token", "access token secret")
	if (err){
		fmt.Println('token invalid');
	}

	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	client := twitter.NewClient(httpClient)

	config1 := tgbotapi.CallbackConfig{queryId,"",false,"",0}
	bot.AnswerCallbackQuery(config1)


	tt,_,err:= client.Timelines.UserTimeline(&twitter.UserTimelineParams{ScreenName:"sentinel_co",Count:3,TweetMode:"extended"})


	if(err != nil){
		log.Println("can't get the twits",err)
	}

	threetweets := [3]int{0,1,2}

	for _,i := range threetweets{

			//ind := len(tt[i].Entities.Urls)

			msg := tgbotapi.NewMessage(chatID,"@" + tt[i].User.ScreenName + " #tweeted on " +tt[i].CreatedAt+"\n"+tt[i].FullText)


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
