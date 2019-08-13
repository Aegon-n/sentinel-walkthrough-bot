package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/Aegon-n/sentinel-bot/handler/buttons"
	"github.com/Aegon-n/sentinel-bot/handler/dbo"
	"github.com/Aegon-n/sentinel-bot/handler/helpers"
	"github.com/Aegon-n/sentinel-bot/handler/messages/en_messages"
	"github.com/Aegon-n/sentinel-bot/handler/models"
	"github.com/Aegon-n/sentinel-bot/locale"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/ldb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleGreet(Bot *tgbotapi.BotAPI, update *tgbotapi.Update, collection *mongo.Collection) {
	username := helpers.GetUserName(update)
	chatID := helpers.GetchatID(update)
	txt := fmt.Sprintf(en_messages.WelcomeGreetMsg, username) + "\n\n\n" + en_messages.SelectwalkthroughMsg
	// msg := tgbotapi.NewMessage(chatID,txt)
	// msg.ReplyMarkup = buttons.GetButtons("LanguageButtons")
	if update.CallbackQuery != nil {
		msgID := update.CallbackQuery.Message.MessageID
		msg := tgbotapi.NewEditMessageText(chatID, msgID, txt+"\n\n"+"Choose an option from the list below: ")
		btns := buttons.GetButtons("HomeButtonsList")
		msg.ReplyMarkup = &btns
		msg.ParseMode = tgbotapi.ModeMarkdown
		Bot.Send(msg)
		return
	}
	msg2 := tgbotapi.NewMessage(chatID, txt+"\n\n"+"Choose an option from the list below: ")
	msg2.ReplyMarkup = buttons.GetButtons("HomeButtonsList")
	msg2.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg2)
	var user models.BotUser

	err := collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		user = models.BotUser{UserName: username, FirstName: update.Message.Chat.FirstName, ChatID: chatID}
		insertResult, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		return
	}
	fmt.Printf("Found a single document: %+v\n", user)

}

func HandlerWalkThrough(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

	username := helpers.GetUserName(update)
	chatID := helpers.GetchatID(update)
	txt := fmt.Sprintf(en_messages.WalkthroughGreetMsg, username) + "\n\n" + en_messages.AppSelectMsg
	msg := tgbotapi.NewMessage(chatID, txt)
	msg.ReplyMarkup = buttons.GetButtons("AppButtonsList")
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
}

func HandleUpdates(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, en_messages.SelectUpdateBlog)
	msg.ReplyMarkup = buttons.GetButtons("UpdatesButtonList")
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
}
func HandleHelp(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, en_messages.HelpMsg)
	msg.ParseMode = tgbotapi.ModeHTML
	Bot.Send(msg)
}
func HandleCallbackQuery(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

	switch update.CallbackQuery.Data {

	case "Home":
		HandlerWalkThrough(bot, update)

	case "Sentinel-Desktop":
		handleAppVersion(bot, update, "Desktop")

	case "Sentinel-Mobile":
		handleAppVersion(bot, update, "Mobile")

	case "Linux":
		handleOs(bot, update, "Linux")

	case "Windows":
		handleOs(bot, update, "Windows")

	case "Mac":
		handleOs(bot, update, "Mac")

	case "Exit":
		handleExit(bot, update)

	/* case "English":
		handleLang(bot, update, "English", db)

	case "Russian":
		handleLang(bot, update, "Russian", db)

	case "Chinese":
		handleLang(bot, update, "Chinese", db) */

	default:
		chatID := update.CallbackQuery.Message.Chat.ID
		txt := "Not implemented"
		msg := tgbotapi.NewMessage(chatID, txt)
		bot.Send(msg)
	}
}

func handleHome(Bot *tgbotapi.BotAPI, update *tgbotapi.Update, username string) {

	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	msg := tgbotapi.NewEditMessageText(chatID, msgID, fmt.Sprintf(en_messages.WelcomeGreetMsg, username)+"\n"+en_messages.AppSelectMsg)
	msg.ParseMode = tgbotapi.ModeMarkdown
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.GetButtons("AppButtonsList"))

	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)
}

func handleAppVersion(Bot *tgbotapi.BotAPI, update *tgbotapi.Update, version string) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID

	msg := tgbotapi.EditMessageTextConfig{}

	replyMarkup := tgbotapi.InlineKeyboardMarkup{}
	if version == "Desktop" {
		msg = tgbotapi.NewEditMessageText(chatID, msgID, en_messages.DesktopOSSelectMsg)
		replyMarkup = buttons.GetButtons("DesktopOSButtonsList")
	}
	if version == "Mobile" {
		msg = tgbotapi.NewEditMessageText(chatID, msgID, en_messages.MobileOSSelectMsg)
		replyMarkup = buttons.GetButtons("MobileOSButtonsList")
	}
	btns := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, replyMarkup)
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)

}
func handleOs(Bot *tgbotapi.BotAPI, update *tgbotapi.Update, os string) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	msg := tgbotapi.EditMessageTextConfig{}
	btns := tgbotapi.EditMessageReplyMarkupConfig{}
	if os == "Linux" {

		msg = tgbotapi.NewEditMessageText(chatID, msgID, en_messages.LinuxNetworkSelectMsg)
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.GetButtons("LinuxNetworkButtonList"))

	}
	if os == "Windows" {
		msg = tgbotapi.NewEditMessageText(chatID, msgID, en_messages.WindowsNetworkSelectMsg)
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.GetButtons("WindowsNetworkButtonList"))

	}
	if os == "Mac" {
		msg = tgbotapi.NewEditMessageText(chatID, msgID, en_messages.MacNetworkSelectMsg)
		btns = tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, buttons.GetButtons("MacNetworkButtonList"))

	}
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
	Bot.Send(btns)
}

func answeredCallback(Bot *tgbotapi.BotAPI, queryId string) {
	config := tgbotapi.CallbackConfig{queryId, "", false, "", 0}
	Bot.AnswerCallbackQuery(config)
}
func handleExit(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	chatID := update.CallbackQuery.Message.Chat.ID
	msgID := update.CallbackQuery.Message.MessageID
	msg := tgbotapi.NewEditMessageText(chatID, msgID, en_messages.ExitMsg)
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
}
func HandleLocalization(Bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	lang := dbo.GetUserLang(update.Message.From.UserName)
	msg := tgbotapi.NewMessage(chatID, en_messages.LangSelectMsg[lang])
	msg.ReplyMarkup = buttons.GetButtons("LanguageButtons")
	Bot.Send(msg)
}

func handleLang(Bot *tgbotapi.BotAPI, update *tgbotapi.Update, lang string, db ldb.BotDB) {
	queryID := update.CallbackQuery.ID
	answeredCallback(Bot, queryID)
	err := db.Insert("lang", update.CallbackQuery.From.UserName, lang)
	if err != nil {
		log.Fatal("Error adding user language preferences..")
	}
	log.Println("Added user language preferences")
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, locale.LocalizeTemplate(en_messages.LangChosenMsg, struct{ Langchosen string }{lang}, lang))
	Bot.Send(msg)
}
