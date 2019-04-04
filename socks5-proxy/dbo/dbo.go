package dbo

import (
	"fmt"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/helpers"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/constants"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/ldb"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/models"
)

type Level struct {
	db *leveldb.DB
	//nodes []models.TONNode
}

func NewDB() (ldb.BotDB, error) {

	db, err := leveldb.OpenFile("./store", nil)
	return Level{db: db}, err
}

// state type is int16 because our app is never going to exceed the limits of int8
func (l Level) SetEthState(username string, state int8) error {
	return l.Insert(constants.EthState, username, strconv.Itoa(int(state)))
}

// state type is int16 because our app is never going to exceed the limits of int8
func (l Level) GetEthState(username string) (int8, error) {
	pair, err := l.Read(constants.EthState, username)
	if err != nil {
		return 0, err
	}
	i, e := strconv.ParseInt(pair.Value, 10, 8)

	return int8(i), e
}

func (l Level) SetTMState(username string, state int8) error {
	return l.Insert(constants.TMState, username, strconv.Itoa(int(state)))
}

func (l Level) GetTMState(username string) (int8, error) {
	pair, err := l.Read(constants.TMState, username)
	if err != nil {
		return 0, err
	}
	i, e := strconv.ParseInt(pair.Value, 10, 8)

	return int8(i), e
}

func (l Level) EthUserState(username string) []models.KV {

	keys := []string{
		constants.EthAddr, constants.Timestamp, constants.Node,
		constants.Bandwidth, constants.NodeWallet, constants.NodePrice,
		constants.IPAddr, constants.AssignedNodeURI, constants.IsAuth,
		constants.Password, constants.EthState,
	}

	return l.MultiReader(keys, username)
}

func (l Level) TMUserState(username string) []models.KV {

	keys := []string{
		constants.WalletTM, constants.TimestampTM, constants.NodeTM,
		constants.BandwidthTM, constants.NodeWalletTM, constants.NodePriceTM,
		constants.IPAddrTM, constants.AssignedNodeURITM, constants.IsAuthTM,
		constants.PasswordTM, constants.TMState,
	}

	return l.MultiReader(keys, username)
}

func (l Level) Insert(key, username, value string) error {
	k := []byte(key + username)
	v := []byte(value)
	return l.db.Put(k, v, nil)
}

func (l Level) Delete(key, username string) error {

	return nil
}

func (l Level) Read(key, username string) (models.KV, error) {
	k := []byte(key + username)
	v, e := l.db.Get(k, nil)
	if e != nil {
		return models.KV{}, e
	}

	return models.KV{
		Key:   fmt.Sprintf("%s", k),
		Value: fmt.Sprintf("%s", v),
	}, e
}

func (l Level) MultiWriter(pairs []models.KV, username string) error {
	for _, pair := range pairs {
		err := l.Insert(pair.Key, username, pair.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l Level) MultiReader(keys []string, username string) []models.KV {
	var result []models.KV
	for _, key := range keys {
		kv, err := l.Read(key, username)
		if err != nil {
			continue
			//return result, err
		}
		result = append(result, kv)
	}
	//return result, nil
	return result
}

func (l Level) IterateExpired() ([]models.ExpiredUsers, error) {
	itr := l.db.NewIterator(util.BytesPrefix([]byte(constants.Timestamp)), nil)

	var usersWithTimestamp []models.ExpiredUsers
	for itr.Next() {
		usersWithTimestamp = append(usersWithTimestamp, models.ExpiredUsers{
			Key: fmt.Sprintf("%s", itr.Key()), Value: fmt.Sprintf("%s", itr.Value()),
		})
	}
	itr.Release()
	err := itr.Error()
	if err != nil {
		return usersWithTimestamp, err
	}

	return usersWithTimestamp, err
}

func (l Level) PartialSearch(key string) []models.KV {
	var values []models.KV

	itr := l.db.NewIterator(util.BytesPrefix([]byte(key)), nil)

	for itr.Next() {
		values = append(values, models.KV{
			Value: fmt.Sprintf("%s", itr.Value()),
			Key: fmt.Sprintf("%s", itr.Key()),
		})
	}

	itr.Release()

	return values
}

func (l Level) Iterate() ([]models.User, error) {
	itr := l.db.NewIterator(nil, nil)

	var p []models.User
	var w []models.KV
	for itr.Next() {
		key := fmt.Sprintf("%s", itr.Key())
		value := fmt.Sprintf("%s", itr.Value())
		w = append(w, models.KV{Key: key, Value: value})
	}
	defer itr.Release()
	err := itr.Error()

	if err != nil {
		return []models.User{}, err
	}

	for _, user := range w {
		username := helpers.GetTelegramUsername(user.Key)
		var participant models.User
		if username != "" {
			for _, u := range w {
				if u.Key == constants.EthAddr+username {
					participant.EthAddr = u.Value
					participant.TelegramUsername = username
				} else if u.Key == constants.Timestamp+username {
					t, err := time.Parse(time.RFC3339, u.Value)
					if err != nil {
						return []models.User{}, err
					}
					participant.Timestamp = t
				} else if u.Key == constants.Node+username {

				}
			}
		}
		if participant.EthAddr != "" && participant.TelegramUsername != "" {
			p = append(p, participant)
		}
	}

	return p, err
}

func (l Level) RemoveETHUser(username string) error {
	if e := l.db.Delete([]byte(constants.Timestamp+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte(constants.IsAuth+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte(constants.Node+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte(constants.Password+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte(constants.Bandwidth+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte(constants.AssignedNodeURI+username), nil); e != nil {
		return e
	}
	return nil
}

func (l Level) RemoveTMUser(username string) error {
	if e := l.db.Delete([]byte(constants.Timestamp+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte(constants.IsAuthTM+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte(constants.NodeTM+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte(constants.PasswordTM+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte(constants.BandwidthTM+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte(constants.AssignedNodeURITM+username), nil); e != nil {
		return e
	}
	return nil
}
