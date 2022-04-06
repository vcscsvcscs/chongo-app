package controllers

import (
	"github.com/vcscsvcscs/chongo-app/backend/controllers/accounts"
	"log"
	"time"
)

type Database struct {
	users accounts.UsersDB
}

//this function should always be started with go routin. Innitializes database informations and starts the deleted account collection.
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
