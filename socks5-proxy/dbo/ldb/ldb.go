package ldb

import "github.com/Aegon-n/sentinel-bot/socks5-proxy/dbo/models"

type BotDB interface {
	// ETHUserState returns the current state of the user in the bot db
	EthUserState(string) []models.KV
	// TMUserState returns the current state of the user in TM the bot db
	TMUserState(string) []models.KV
	// GetState and SetState are getter setter methods for user state in the app
	GetEthState(string) (int8, error)
	SetEthState(string, int8) error
	GetTMState(string) (int8, error)
	SetTMState(string, int8) error
	// Insert is used to store a new key-value pair
	Insert(string, string, string) error
	// Delete would remove one key-pair from the database
	Delete(string, string) error
	// Read would return a key-value pair for a query
	Read(string, string) (models.KV, error)
	// RemoveUser Would delete all of the user info
	RemoveETHUser(string) error
	// RemoveUser Would delete all of the user info
	RemoveTMUser(string) error
	// Perform a partial search for a bunch of key value pairs
	PartialSearch(string) []models.KV
	// Iterate over the entire database and find all the users
	Iterate() ([]models.User, error)
	// IterateExpired would return a slice of expired users
	IterateExpired() ([]models.ExpiredUsers, error)
	// MultiRead would return a slice of your key value
	// pairs just to avoid too much redundant code
	MultiReader([]string, string) []models.KV
	// MultiWriter would write multiple key value pairs
	// into database to avoid multiple calls to Insert() inside a method
	MultiWriter([]models.KV, string) error
}
