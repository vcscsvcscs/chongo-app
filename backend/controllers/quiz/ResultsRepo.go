package quiz

import (
	"time"

	"github.com/vcscsvcscs/chongo-app/backend/controllers/quiz/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ResultRepo struct {
	db *mgo.Collection
}

func NewResultsRepo(db *mgo.Collection) ResultRepo {
	return ResultRepo{db: db}
}

func (u ResultRepo) FindByUserNameAndId(userName string, id string, result *model.Result) bool {
	iter := u.db.Find(bson.M{"id": id}).Iter()
	return !iter.Next(&result)
}

func (u ResultRepo) Update(userName string, id string, t time.Time, answers map[string]int) error {
	return u.db.Update(bson.M{"$and": []bson.M{{"id": id}, {"username": userName}}}, bson.M{"$and": []bson.M{{"$set": bson.M{"answers": answers}}, {"$set": bson.M{"date": t.Unix()}}}})
}

func (u ResultRepo) Insert(Result *model.Result) error {
	return u.db.Insert(&Result)
}

func (u ResultRepo) Remove(userName string, id string) error {
	err := u.db.Remove(bson.M{"id": id})
	return err
}
