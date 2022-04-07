package sessionmanager

import (
	"time"

	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SessionRepo struct {
	db *mgo.Collection
}

func NewSessionRepo(db *mgo.Collection) SessionRepo {
	return SessionRepo{db: db}
}

func (s SessionRepo) FindByToken(token string, profile *model.Token) bool {
	return s.db.Find(bson.M{"token": token}).Iter().Next(&profile)
}

func (s SessionRepo) FindByTimeAccessed(duration time.Duration) *mgo.Iter {
	return s.db.Find(bson.M{"timeAccessed": bson.M{"$lt": time.Now().Add(duration).Unix()}}).Iter()
}

func (s SessionRepo) RemoveAll(duration time.Duration) error {
	_, err := s.db.RemoveAll(bson.M{"timeAccessed": bson.M{"$lt": time.Now().Add(duration).Unix()}})
	return err
}

func (s SessionRepo) Update(token string, t time.Time) error {
	return s.db.Update(bson.M{"token": token}, bson.M{"$set": bson.M{"timeAccessed": t.Unix()}})
}

func (s SessionRepo) Insert(token, userName string, t time.Time) error {
	return s.db.Insert(&model.Token{userName, token, t.Unix()})
}

func (s SessionRepo) Remove(token string) error {
	return s.db.Remove(bson.M{"token": token})
}
