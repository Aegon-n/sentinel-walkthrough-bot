package dbo

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)
var ldb *leveldb.DB

type LinkBot struct {}

func  NewDB() {
	db, err := leveldb.OpenFile("./database", nil)
	if err != nil {
		log.Fatal(err)
	}
	ldb = db
	log.Println("Successfully initialized DB..")
}

func AddUserLang(username, lang string) error {
	
	return ldb.Put([]byte(username+"lang"),[]byte(lang), nil)

}
func GetUserLang(username string) string {
	lang,err := ldb.Get([]byte(username+"lang"),nil)
	if err != nil {
		log.Println("user not found")
		return "English"
	}
	return string(lang)
}