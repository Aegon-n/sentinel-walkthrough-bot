package dbo

import (
	"fmt"

	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo/ldb"
	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo/models"
	"github.com/Aegon-n/sentinel-bot/eth-socks-proxy/helpers"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/constants"
	"github.com/syndtr/goleveldb/leveldb"
)

type Level struct {
	db *leveldb.DB
	//nodes []models.TONNode
}

func NewDB() (ldb.BotDB, error) {

	db, err := leveldb.OpenFile("./eth-socks-proxy/store", nil)
	return Level{db: db}, err
}

func (l Level) Insert(key, username, value string) error {
	k := []byte(key + username)
	v := []byte(value)
	return l.db.Put(k, v, nil)
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

func (l Level) Delete(key, username string) error {

	return nil
}

func (l Level) SetStatus(username string, status string) error {
	return l.Insert("SPSSTATUS", username, status)
}

func (l Level) GetStatus(username string) (string, error) {
	pair, err := l.Read("SPSSTATUS", username)
	if err != nil {
		return "", err
	}

	return pair.Value, nil
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
func (l Level) RemoveUser(username string) error {

	if e := l.db.Delete([]byte("SPSSTATUS"+username), nil); e != nil {
		return e
	}

	if e := l.db.Delete([]byte("NodeInfo"+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte(constants.AssignedNodeURI+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte("NodeIP"+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte("TOKEN"+username), nil); e != nil {
		return e
	}
	if e := l.db.Delete([]byte("ChatID"+username), nil); e != nil {
		return e
	}
	return nil
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
				if u.Key == "NodeIP"+username {
					participant.Node = u.Value
					participant.TelegramUsername = username
				} else if u.Key == "ChatID"+username {
					participant.ChatID = u.Value
				} else if u.Key == "TOKEN"+username {
					participant.Token = u.Value
				} 

			}
		}
		if participant.Node != "" && participant.TelegramUsername != "" && participant.ChatID != "" && participant.Token != "" {
			p = append(p, participant)
		}
	}
	return p, err
}
