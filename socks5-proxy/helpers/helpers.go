package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/buttons"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/constants"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/ldb"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/models"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/templates"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
	"strings"
	"time"
)

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

func SubscriptionPeriod(b *tgbotapi.BotAPI, u tgbotapi.Update, db ldb.BotDB, t time.Duration, network, price, period string) {
	EthPairs := []models.KV{
		{
			Key: constants.Timestamp, Value: time.Now().Add(t).Format(time.RFC3339),
		},
		{
			Key: constants.NodePrice, Value: price,
		},
	}
	TMPairs := []models.KV{
		{
			Key: constants.TimestampTM, Value: time.Now().Add(t).Format(time.RFC3339),
		},
		{
			Key: constants.NodePriceTM, Value: price,
		},
	}
	if network == constants.EthNetwork {
		err := db.MultiWriter(EthPairs, u.CallbackQuery.From.UserName)
		if err != nil {
			Send(b, u, templates.BWError)
		}
		msg := fmt.Sprintf(templates.BWPeriods, period)
		Send(b, u, msg)
		return
	}
	err := db.MultiWriter(TMPairs, u.CallbackQuery.From.UserName)
	if err != nil {
		Send(b, u, templates.BWError)
	}
	msg := fmt.Sprintf(templates.BWPeriods, period)
	Send(b, u, msg)

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

func GetNodes() (models.Nodes, error) {
	var body []models.TONNode
	var N models.Nodes
	resp, err := http.Get(constants.SentinelTONURL)
	if err != nil {
		return N, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return N, err
	}
	defer resp.Body.Close()

	for _, node := range body {
		if node.Type == constants.NodeType {
			N.TMNodes = append(N.TMNodes, node)
		} else {
			N.EthNodes = append(N.EthNodes, node)
		}
	}
	return N, err
}

func SetState(b *tgbotapi.BotAPI, u tgbotapi.Update, network string, state int8, db ldb.BotDB) {
	username := ""
	if u.CallbackQuery != nil {
		username = u.CallbackQuery.From.UserName
	}
	if u.Message != nil {
		username = u.Message.From.UserName
	}
	if network == constants.TMState {
		err := db.SetTMState(username, state)
		if err != nil {
			Send(b, u, "invalid tendermint user set state")
			return
		}
		return
	}
	err := db.SetEthState(username, state)
	if err != nil {
		Send(b, u, "invalid ethereum user set state")
		return
	}
}

func GetState(b *tgbotapi.BotAPI, u tgbotapi.Update, network string, db ldb.BotDB) int8 {
	username := ""
	if u.CallbackQuery != nil {
		username = u.CallbackQuery.From.UserName
	}
	if u.Message != nil {
		username = u.Message.From.UserName
	}
	if network == constants.TMState {
		state, err := db.GetTMState(username)
		if err != nil {
			Send(b, u, "invalid tendermint user get state")
			return constants.NoState
		}

		return state
	}

	state, err := db.GetEthState(username)
	if err != nil {
		Send(b, u, "invalid ethereum user get state")
		return constants.NoState
	}

	return state
}
