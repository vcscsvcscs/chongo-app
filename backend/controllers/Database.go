package controllers

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var MongoDB *mgo.Session

var Users *mgo.Collection

type User struct {
	Username       string `json:"username" bson:"username"`
	Name           string `json:"name" bson:"name"`
	Email          string `json:"email" bson:"email"`
	Password       string `json:"password" bson:"password"`
	Ip             string `json:"ip" bson:"ip"`
	DeletedAccount int64  `json:"deleted" bson:"deleted"`
}

//this function should always be started with go routin. Innitializes database informations and starts the deleted account collection.
func InnitCredentials(mongoDB *mgo.Session, dbname string, collectionname string) {
	MongoDB = mongoDB
	Users = MongoDB.DB(dbname).C(collectionname)
	AccountDeleteCollector()
}

//A timed garbage collector which checks every 48 hours if accounts have been deleted and, if an account was deleted 15 days ago, it deletes it.
func AccountDeleteCollector() {
	for range time.Tick(time.Hour * 48) {
		_, err := Users.RemoveAll(bson.M{"$and": []bson.M{{"deleted": bson.M{"$ne": 0}}, {"deleted": bson.M{"$lt": time.Now().Add(time.Hour * (-1) * 336).Unix()}}}})
		if err != nil {
			log.Println(err)
		}
	}
}
