package quiz

import (
	"github.com/vcscsvcscs/chongo-app/backend/controllers/quiz/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type QuizRepo struct {
	db *mgo.Collection
}

func NewQuizRepo(db *mgo.Collection) QuizRepo {
	return QuizRepo{db: db}
}

func (u QuizRepo) FindById(id string, quiz *model.Quiz) bool {
	iter := u.db.Find(bson.M{"id": id}).Iter()
	return !iter.Next(&quiz)
}

func (u QuizRepo) Update(id string, quiz *model.Quiz) error {
	return u.db.Update(bson.M{"id": id}, quiz)
}

func (u QuizRepo) Insert(quiz *model.Quiz) error {
	return u.db.Insert(&quiz)
}

func (u QuizRepo) Remove(id string) error {
	err := u.db.Remove(bson.M{"id": id})
	return err
}
