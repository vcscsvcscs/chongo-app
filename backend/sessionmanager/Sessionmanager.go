package sessionmanager

import (
	"fmt"
	"log"
	"time"

	"github.com/vcscsvcscs/chongo-app/backend/utilities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var MongoDB *mgo.Session
var Sessions *mgo.Collection
var MaxAge time.Duration
var Users = make(map[string]string)
var Online = make(map[string]bool)

type Token struct {
	Username    string `json:"username" bson:"username"`
	Token       string `json:"token" bson:"token"`
	TimeAccesed int64  `json:"timeaccesed" bson:"timeaccesed"`
}

func InnitSessions(maxAge int, mongoDB *mgo.Session, dbname string, collectionname string) {
	MongoDB = mongoDB
	Sessions = MongoDB.DB(dbname).C(collectionname)
	MaxAge = time.Minute
	for i := 1; i < maxAge; i++ {
		MaxAge -= time.Minute
	}
	SessionCollector()
}

func SessionCollector() {
	for range time.Tick(time.Minute * 15) {
		var profile Token
		shouldelete := Sessions.Find(bson.M{"timeaccesed": bson.M{"$lt": time.Now().Add(MaxAge).Unix()}}).Iter()
		for shouldelete.Next(&profile) {
			delete(Online, Users[profile.Token])
			delete(Users, profile.Token)
		}
		_, err := Sessions.RemoveAll(bson.M{"timeaccesed": bson.M{"$lt": time.Now().Add(MaxAge).Unix()}})
		if err != nil {
			log.Println(err)
		}
	}
}

func IsSessionLegit(token string) (Token, bool) {
	var profile Token
	legit := Sessions.Find(bson.M{"token": token}).Iter().Next(&profile)
	if !legit {
		delete(Users, token)
	} else {
		Sessions.Update(bson.M{"token": token}, bson.M{"$set": bson.M{"timeaccesed": time.Now().Unix()}})
		Online[Users[token]] = true
	}
	return profile, legit
}

//This is the function which generates a token and returns it, The token is automaticaly added to the connected mongoDB and it is added to the local cache.
func SetSessionKeys(ClientIP string, username string) string {
	token := utilities.Md(fmt.Sprint(ClientIP, time.Now()))
	Users[token] = username
	Online[username] = true
	if err := Sessions.Insert(&Token{username, token, time.Now().Unix()}); err != nil {
		log.Println(err)
	}
	return token
}

func DeleteSessionKey(token string) bool {
	if err := Sessions.Remove(bson.M{"token": token}); err != nil {
		log.Println(err)
		return false
	}
	delete(Online, Users[token])
	delete(Users, token)

	return true
}
