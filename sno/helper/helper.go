package helper

import "gopkg.in/telegram-bot-api.v4"

func GetUserName(u tgbotapi.Update) string {
	var username string
	if u.CallbackQuery != nil {
		username = u.CallbackQuery.Message.Chat.UserName
	}
	if u.Message != nil {
		username = u.Message.From.UserName
	}
	return username
}

func GetchatID(u tgbotapi.Update) int64 {
	var chatID int64
	if u.CallbackQuery != nil {
		chatID = u.CallbackQuery.Message.Chat.ID
	}
	if u.Message != nil {
		chatID = u.Message.Chat.ID
	}
	return chatID
}
func GetMsgID(u tgbotapi.Update) int {
	var msgID int
	if u.CallbackQuery != nil {
		msgID = u.CallbackQuery.Message.MessageID
	}
	if u.Message != nil {
		msgID = u.Message.MessageID
	}
	return msgID
}

func Send(b *tgbotapi.BotAPI, u tgbotapi.Update, txt string, btns ...*tgbotapi.InlineKeyboardMarkup) {
	
	chatID := GetchatID(u)
	var msgID int
	if u.CallbackQuery != nil && len(btns) > 0 {
		msgID = u.CallbackQuery.Message.MessageID
		msg := tgbotapi.NewEditMessageText(chatID, msgID, txt)
	  if len(btns) > 0 {
	  	msg.ReplyMarkup = btns[0]
	  }
	  msg.ParseMode = tgbotapi.ModeMarkdown
	  b.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(chatID, txt)
		if len(btns) > 0 {
	  	msg.ReplyMarkup = btns[0]
	  }
	  msg.ParseMode = tgbotapi.ModeMarkdown
	  b.Send(msg)
	}
	return
}