package helpers

import "gopkg.in/telegram-bot-api.v4"

/*func Send(bot *tgbotapi.BotAPI, update *tgbotapi.Update, typ string, ) {
	l=localize(text)
	bot.ne(l)
}*/



func GetUserName(u *tgbotapi.Update) string {
	var username string
	if u.CallbackQuery != nil {
		username = u.CallbackQuery.Message.Chat.UserName
	}
	if u.Message != nil {
		username = u.Message.From.UserName
	}
	return username
}

func GetchatID(u *tgbotapi.Update) int64 {
	var chatID int64
	if u.CallbackQuery != nil {
		chatID = u.CallbackQuery.Message.Chat.ID
	}
	if u.Message != nil {
		chatID = u.Message.Chat.ID
	}
	return chatID
}
