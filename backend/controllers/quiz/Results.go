package quiz

import (
	"time"

	"github.com/vcscsvcscs/chongo-app/backend/controllers/quiz/model"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager"
)

type ResultsDB interface {
	FindByUserNameAndId(userName string, id string, result *model.Result) bool
	Update(userName string, id string, t time.Time, answers map[string]int) error
	Insert(result *model.Result) error
	Remove(userName string, id string) error
}

type Results struct {
	sessionManager sessionmanager.SessionManager
	db             ResultsDB
}

func NewResults(sessionManager sessionmanager.SessionManager, db ResultsDB) Results {
	return Results{
		sessionManager: sessionManager,
		db:             db,
	}
}
