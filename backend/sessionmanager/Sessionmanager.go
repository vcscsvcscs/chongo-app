package sessionmanager

import (
	"fmt"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager/model"
	"log"
	"time"

	"github.com/vcscsvcscs/chongo-app/backend/utilities"
	"gopkg.in/mgo.v2"
)

//go:generate mockgen -source=Sessionmanager.go -destination=./mocks/interfaces.go -package=mocks

type SessionsDB interface {
	FindByToken(token string, profile *model.Token) bool
	FindByTimeAccessed(duration time.Duration) *mgo.Iter
	RemoveAll(duration time.Duration) error
	Update(token string, t time.Time) error
	Insert(token, userName string, t time.Time) error
	Remove(token string) error
}

type SessionManager struct {
	sessions SessionsDB
	maxAge   time.Duration
	users    map[string]string
	online   map[string]bool
	clock    utilities.Clock
}

//Initialization of the mongo database, have to set maxage in minute.
func InitSessions(maxAge int, mongoDB *mgo.Session, dbName, collectionName string, clock utilities.Clock) SessionManager {
	sm := SessionManager{
		sessions: NewSessionRepo(mongoDB.DB(dbName).C(collectionName)),
		maxAge:   time.Minute,
		users:    make(map[string]string),
		online:   make(map[string]bool),
	}

	for i := 1; i < maxAge; i++ {
		sm.maxAge -= time.Minute
	}

	go sm.sessionCollector()

	return sm
}

//Session garbage collector, deletes old invalid sessions and sets users offline.
func (sm *SessionManager) sessionCollector() {
	for range time.Tick(time.Minute * 15) {
		var profile model.Token
		shouldDelete := sm.sessions.FindByTimeAccessed(sm.maxAge)
		for shouldDelete.Next(&profile) {
			delete(sm.online, sm.users[profile.Token])
			delete(sm.users, profile.Token)
		}

		if err := sm.sessions.RemoveAll(sm.maxAge); err != nil {
			log.Println(err)
		}
	}
}

//Checks if Session is in the database, and its still valid, it also refreshes the tokens life and sets the user to online state.
func (sm *SessionManager) IsSessionLegit(token string) (model.Token, bool) {
	var profile model.Token
	legit := sm.sessions.FindByToken(token, &profile)
	if !legit {
		delete(sm.users, token)
	} else {
		sm.sessions.Update(token, sm.clock.Now())
		sm.online[sm.users[token]] = true
	}
	return profile, legit
}

//This is the function which generates a token and returns it, The token is automaticaly added to the connected mongoDB and it is added to the local cache.
func (sm *SessionManager) SetSessionKeys(ClientIP string, username string) string {
	token := utilities.Md(fmt.Sprint(ClientIP, sm.clock.Now()))
	sm.users[token] = username
	sm.online[username] = true
	if err := sm.sessions.Insert(token, username, sm.clock.Now()); err != nil {
		log.Println(err)
	}
	return token
}

//Deletes session token from the database, sets user offline.
func (sm *SessionManager) DeleteSessionKey(token string) bool {
	if err := sm.sessions.Remove(token); err != nil {
		log.Println(err)
		return false
	}
	delete(sm.online, sm.users[token])
	delete(sm.users, token)

	return true
}

func (sm *SessionManager) GetUser(token string) string {
	return sm.users[token]
}
