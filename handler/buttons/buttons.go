package buttons

import "gopkg.in/telegram-bot-api.v4"


func DesktopOsButtons(data1, data2, data3 string) tgbotapi.InlineKeyboardMarkup {
	home := tgbotapi.NewInlineKeyboardButtonData("Home","Home")
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Linux",data1)
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Windows",data2)
	btn3 := tgbotapi.NewInlineKeyboardButtonData("Mac OS",data3)
	btns := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{home,btn1,btn2,btn3}},
	}
	return btns
}

func MobileOsButtons(data1, data2 string) tgbotapi.InlineKeyboardMarkup {
	home := tgbotapi.NewInlineKeyboardButtonData("Home","Home")
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Android",data1)
	btn2 := tgbotapi.NewInlineKeyboardButtonData("IOS",data2)
	btns := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{home,btn1,btn2}},
	}
	return btns
}

func TestNetButtons(data1, data2 string) tgbotapi.InlineKeyboardMarkup {
	home := tgbotapi.NewInlineKeyboardButtonData("Home","Home")
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Ethereum", data1)
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Tendermint", data2)
	btns := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{home, btn1, btn2}},
	}
	return btns
}

func PersistentNavButtons(data1, data2, data3 string) tgbotapi.InlineKeyboardMarkup {
	home := tgbotapi.NewInlineKeyboardButtonData("Home","Home")
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Prev",data1)
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Skip",data2)
	btn3 := tgbotapi.NewInlineKeyboardButtonData("Next",data3)
	btns := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{home, btn1,btn2,btn3}},
	}
	return btns
}

func ModulesListButton(data, downloadurl, videourl  string) tgbotapi.InlineKeyboardMarkup {

	home := tgbotapi.NewInlineKeyboardButtonData("Home","Home")
	btn1 := tgbotapi.NewInlineKeyboardButtonData("View Chapters",data)
	down := tgbotapi.NewInlineKeyboardButtonURL("Download Full Process PDF file", downloadurl)

	video := tgbotapi.NewInlineKeyboardButtonURL("Watch Walkthrough Video",videourl)
	btns := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{home,btn1},{down},{video}},
	}
	return btns
}
func HomeButtons(data1, data2 string) tgbotapi.InlineKeyboardMarkup {

	btn1 := tgbotapi.NewInlineKeyboardButtonData("Sentinel-Desktop App",data1)
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Sentinel-Mobile App",data2)

	btns := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{{btn1,btn2}},
	}
	return btns
}
