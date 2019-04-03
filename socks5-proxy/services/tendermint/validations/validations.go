package validations

import (
	"encoding/json"
	"fmt"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/constants"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/ldb"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/models"
	"math"
	"net/http"
	"strconv"
	"time"
)

func CheckTMBalance(address string) (float64, bool) {
	var body models.TMMsg
	resp, err := http.Get(fmt.Sprintf(constants.TMBalanceURL, address))
	if err != nil {
		return 0, false
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return 0, false
	}

	userBalance, err := strconv.ParseFloat(body.Value.Coins[0].Amount, 64)
	if err != nil || userBalance < constants.MinBal {
		return userBalance / math.Pow(10, 8), false
	}

	return userBalance / math.Pow(10, 8), true
}

func IsUniqueWallet(wallet, username string, db ldb.BotDB) bool {
	pairs := db.PartialSearch(constants.WalletTM)
	for _, pair := range pairs {
		if pair.Value == wallet && pair.Key != constants.WalletTM+username {
			return false
		}
		return true
	}
	return false
}

func CheckTXNTimeStamp(hash, wallet, timeLimit string) bool {
	var body models.Transactions
	url := fmt.Sprintf(constants.GetTXNFromMN, wallet)
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return false
	}

	for _, txn := range body.List {
		txnTimestamp, err := time.Parse(time.RFC3339, txn.Timestamp)
		if err != nil {
			return false
		}
		l, _ := time.Parse(time.RFC3339, timeLimit)

		if txn.Hash == hash && int(l.Sub(txnTimestamp).Minutes()) < constants.TimeLimit {
			return true
		}
	}
	return false
}