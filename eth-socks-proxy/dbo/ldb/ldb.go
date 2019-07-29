package ldb

import "github.com/Aegon-n/sentinel-bot/eth-socks-proxy/dbo/models"

type BotDB interface {

	// Insert is used to store a new key-value pair
	Insert(string, string, string) error
	// Delete would remove one key-pair from the database
	Delete(string, string) error
	// Read would return a key-value pair for a query
	Read(string, string) (models.KV, error)

	// RemoveUser Would delete all of the user info
	SetStatus(string, string) error
	GetStatus(string) (string, error)
	// pairs just to avoid too much redundant code
	MultiReader([]string, string) []models.KV
	// MultiWriter would write multiple key value pairs
	// into database to avoid multiple calls to Insert() inside a method
	MultiWriter([]models.KV, string) error
	RemoveUser(string) error
}
