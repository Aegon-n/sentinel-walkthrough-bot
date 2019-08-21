package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/buttons"
	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo/ldb"
	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo/models"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/constants"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/templates"
	"github.com/jasonlvhit/gocron"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

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

func GetNumaricKeyBoard(n int) tgbotapi.ReplyKeyboardMarkup {
	btnlist := [][]tgbotapi.KeyboardButton{{}, {}}
	rows := math.Ceil(float64(n) / float64(4))
	fmt.Println(rows)
	number := 0
	for j := 0; j < int(rows); j++ {
		list := []tgbotapi.KeyboardButton{}
		for i := 0; i < 4; i++ {
			number++
			if number > n {
				break
			}
			list = append(list, tgbotapi.NewKeyboardButton(strconv.Itoa(number)))
		}
		btnlist = append(btnlist, list)
	}

	fmt.Println(btnlist)
	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:        btnlist,
		OneTimeKeyboard: true,
		ResizeKeyboard:  true,
		Selective:       true,
	}
}

func Send(b *tgbotapi.BotAPI, u tgbotapi.Update, msg string, opts ...models.ButtonHelper) {
	var chatID int64
	if u.CallbackQuery != nil {
		chatID = u.CallbackQuery.Message.Chat.ID
	}
	if u.Message != nil {
		chatID = u.Message.Chat.ID
	}
	c := tgbotapi.NewMessage(chatID, msg)

	for _, o := range opts {
		if o.Type == constants.ReplyButton {
			c.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
				Keyboard:        buttons.ReplyButtons(o.Labels),
				OneTimeKeyboard: true,
				ResizeKeyboard:  false,
			}
		}
		if o.Type == constants.InlineButton {
			c.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
				InlineKeyboard: buttons.InlineButtons(o.InlineKeyboardOpts),
			}
		}
	}
	c.ParseMode = tgbotapi.ModeMarkdown

	_, _ = b.Send(c)
	//_, e := b.Send(c)
	//color.Red("***** \n ERROR: %v \n*****", e)
}

func GetTelegramUsername(username string) string {

	//username :=  fmt.Sprintf("%s", b)
	//log.Println("\n\n what does it look like? : ", username, "\n\n")
	if len(username) < 1 {
		log.Println("invalid username")
		return ""
	}

	if strings.Contains(username, constants.AssignedNodeURI) {
		return strings.TrimPrefix(username, constants.AssignedNodeURI)
	}

	return ""
}

func GetNodes() ([]models.List, error) {
	var body models.SocksResponse
	var N []models.List
	resp, err := http.Get("https://api.sentinelgroup.io/client/vpn/socks-list")
	if err != nil {
		return N, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return N, err
	}
	defer resp.Body.Close()
	/*if len(body.List) > 60 {
		return body.List[:60], err
	}*/
	return body.List, err
}
func GetAllNodes() ([]models.List, error) {
	var body models.SocksResponse
	var N []models.List
	resp, err := http.Get("https://api.sentinelgroup.io/client/vpn/list")
	if err != nil {
		return N, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return N, err
	}
	defer resp.Body.Close()
	/*if len(body.List) > 60 {
		return body.List[:60], err
	}*/
	return body.List, err
}

func GetToken(vpn_addr, telegramId string) (models.MasterResponce, error) {
	var body models.MasterResponce

	requestBody, err := json.Marshal(map[string]string{
		"device_id": telegramId,
		"vpn_addr":  vpn_addr,
	})
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post("https://api.sentinelgroup.io/client/vpn", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return body, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return body, err
	}
	defer resp.Body.Close()
	return body, err
}

func ConnectNode(telegramId string, resp models.MasterResponce) (models.VpnResponse, error) {
	var body models.VpnResponse

	requestBody, err := json.Marshal(map[string]string{
		"account_addr": telegramId,
		"vpn_addr":     resp.VpnAddr,
		"token":        resp.Token,
	})
	if err != nil {
		log.Fatalln(err)
	}
	url := fmt.Sprintf("http://%s:%s/creds", resp.IP, strconv.Itoa(resp.Port))
	res, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	log.Println(res)
	if err != nil {
		return body, err
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return body, err
	}
	defer res.Body.Close()
	return body, err
}

func SocksProxy(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, vpn_addr string) {
	telegramId := u.Message.From.UserName
	resp, err := GetToken(vpn_addr, telegramId)
	log.Println(resp)
	if err != nil {
		Send(b, u, "Sorry Unable connect node")
		return
	}
	if resp.Success == false {
		Send(b, u, "Sorry Unable connect node")
		return
	}
	_ = db.Insert("TOKEN", u.Message.From.UserName, resp.Token)
	result, err := ConnectNode(telegramId, resp)
	log.Println(result)
	if err != nil {
		Send(b, u, "Sorry Unable connect node")
		return
	}
	if result.Success == false {
		Send(b, u, "Sorry Unable connect node")
		return
	}
	fmt.Println(result)
	fmt.Println(result.Node)
	TG_URL := result.Node.Vpn.TelegramLink
	err = db.Insert(constants.AssignedNodeURI, u.Message.From.UserName, TG_URL)
	if err != nil {
		log.Println("unable to insert proxy url")
		Send(b, u, "interal bot error")
	}
	err = db.SetStatus(u.Message.From.UserName, constants.AssignedNodeURI)
	if err != nil {
		log.Println("unable to set status")
		Send(b, u, "interal bot error")
	}
	optns := [][]tgbotapi.InlineKeyboardButton{{}, {}}
	for idx, row := range []map[string]string{{"connect": TG_URL}, {"◀Back": "sps", "🏠Home": "home"}} {
		for k, v := range row {
			val := v
			if k == "connect" {
				optns[idx] = append(optns[idx], tgbotapi.InlineKeyboardButton{Text: k, URL: &val})
				continue
			}
			optns[idx] = append(optns[idx], tgbotapi.InlineKeyboardButton{Text: k, CallbackData: &val})
		}
	}

	msg := tgbotapi.NewMessage(GetchatID(u), templates.Success)
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: optns}
	b.Send(msg)
	return
}

func DisconnectNode(b *tgbotapi.BotAPI, username string, ip string, token string) {

	values := map[string]string{"account_addr": username, "token": token}
	fmt.Println(values)
	jsonValue, _ := json.Marshal(values)
	url := fmt.Sprintf("http://%s:3000/disconnect", ip)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(resp)
	if resp.StatusCode != 200 {
		log.Println("unable to disconnect")
		return
	}

}

func GetDataUsage(UserName, ip, token string) (models.Usage, error) {
	var body models.VpnUsage
	var usage models.Usage
	values := map[string]string{"account_addr": UserName, "token": token}
	fmt.Println(values)
	jsonValue, _ := json.Marshal(values)
	url := fmt.Sprintf("http://%s:3000/usage", ip)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Println(err)
		return usage, err
	}
	fmt.Println(resp)
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return body.Usage, err
	}
	defer resp.Body.Close()
	return body.Usage, err
}

func CheckAndDisconnectExpiredUsers(bot *tgbotapi.BotAPI, user models.User, db ldb.BotDB) {
	var body models.LimitResponse
	url := fmt.Sprintf("http://%s:3000/limit_reached_ids", user.Node)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Println(err)
		return
	}
	// fmt.Println(body.ClientList)
	if contains(body.ClientList, strings.ToLower(user.TelegramUsername)) {
		chatId, err := strconv.Atoi(user.ChatID)
		if err != nil {
			log.Println(err)
			return
		}
		msg := tgbotapi.NewMessage(int64(chatId), fmt.Sprintf(templates.LIMITEXCEEDED, user.TelegramUsername))
		bot.Send(msg)
		DisconnectNode(bot, user.TelegramUsername, user.Node, user.Token)
		db.RemoveUser(user.TelegramUsername)
	}
}

func CheckLimitExceededUsers(bot *tgbotapi.BotAPI, db ldb.BotDB) {
	AllUsers, err := db.Iterate()
	if err != nil {
		log.Println(err)
		return
	}
	// fmt.Println("All Users:\n", AllUsers)
	for _, user := range AllUsers {
		go CheckAndDisconnectExpiredUsers(bot, user, db)
	}

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ExpiredUsersJob(bot *tgbotapi.BotAPI, db ldb.BotDB) {
	s := gocron.NewScheduler()
	s.Every(30).Seconds().Do(CheckLimitExceededUsers, bot, db)
	<-s.Start()
}
