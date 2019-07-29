package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/buttons"
	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo/ldb"
	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo/models"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/constants"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/templates"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func GetNumaricKeyBoard() tgbotapi.ReplyKeyboardMarkup {
	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("1"),
			tgbotapi.NewKeyboardButton("2"),
			tgbotapi.NewKeyboardButton("3"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("4"),
			tgbotapi.NewKeyboardButton("5"),
			tgbotapi.NewKeyboardButton("6"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("7"),
			tgbotapi.NewKeyboardButton("8"),
			tgbotapi.NewKeyboardButton("9"),
		),
	)
	return numericKeyboard
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
				ResizeKeyboard:  true,
			}
		}
		if o.Type == constants.InlineButton {
			c.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
				InlineKeyboard: buttons.InlineButtons(o.InlineKeyboardOpts),
			}
		}
	}
	c.ParseMode = tgbotapi.ModeHTML
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

	if strings.Contains(username, "telegram") {
		return strings.TrimPrefix(username, "telegram")
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
	if len(body.List) > 20 {
		return body.List[:20], err
	}
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
	btnOpts := []models.InlineButtonOptions{
		{Label: "Sentinel Proxy Node", URL: TG_URL},
	}
	opts := models.ButtonHelper{Type: constants.InlineButton, InlineKeyboardOpts: btnOpts}
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
	Send(b, u, templates.Success, opts)
	return
}

func DisconnectNode(b *tgbotapi.BotAPI, u tgbotapi.Update, ip string, token string) {

	values := map[string]string{"account_addr": u.Message.From.UserName, "token": token}
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

func GetDataUsage(u tgbotapi.Update, ip, token string) (models.Usage, error) {
	var body models.VpnUsage
	var usage models.Usage
	values := map[string]string{"account_addr": u.Message.From.UserName, "token": token}
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
