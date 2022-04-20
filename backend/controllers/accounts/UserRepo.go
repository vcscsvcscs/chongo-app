package accounts

import (
	"time"

	"github.com/vcscsvcscs/chongo-app/backend/controllers/accounts/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepo struct {
	db *mgo.Collection
}

func NewUserRepo(db *mgo.Collection) UserRepo {
	return UserRepo{db: db}
}

func (u UserRepo) FindByEmail(email string, user *model.User) bool {
	iter := u.db.Find(bson.M{"email": email}).Iter()
	return !iter.Next(&user)
}

func (u UserRepo) FindByUserName(userName string, user *model.User) bool {
	iter := u.db.Find(bson.M{"username": userName}).Iter()
	return !iter.Next(&user)
}

func (u UserRepo) Update(userName string, t time.Time) error {
	return u.db.Update(bson.M{"username": userName}, bson.M{"$set": bson.M{"deleted": t.Unix()}})
}

func (u UserRepo) Insert(user *model.User) error {
	return u.db.Insert(&user)
}

func (u UserRepo) RemoveAll(duration time.Duration) error {
	_, err := u.db.RemoveAll(bson.M{"$and": []bson.M{{"deleted": bson.M{"$ne": 0}}, {"deleted": bson.M{"$lt": time.Now().Add(duration).Unix()}}}})

	return err
}
