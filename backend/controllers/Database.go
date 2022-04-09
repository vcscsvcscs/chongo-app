package controllers

import (
	"log"
	"time"

	"github.com/vcscsvcscs/chongo-app/backend/controllers/accounts"
)

type Database struct {
	users accounts.UsersDB
}

func InitCredentials(usersDB accounts.UsersDB) Database {
	db := Database{
		users: usersDB,
	}

	go db.AccountDeleteCollector()

	return db
}

//A timed garbage collector which checks every 48 hours if accounts have been deleted and, if an account was deleted 15 days ago, it deletes it.
func (db *Database) AccountDeleteCollector() {
	for range time.Tick(time.Hour * 48) {
		err := db.users.RemoveAll(time.Hour * (-1) * 336)
		if err != nil {
			log.Println(err)
		}
	}
}
