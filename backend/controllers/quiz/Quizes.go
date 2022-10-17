package quiz

import (
	"github.com/vcscsvcscs/chongo-app/backend/controllers/quiz/model"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager"
)

type QuizDB interface {
	FindById(id string, quiz *model.Quiz) bool
	Update(id string, quiz *model.Quiz) error
	Insert(quiz *model.Quiz) error
	Remove(id string) error
}

//go:generate mockgen -source=Quizes.go -destination=./mocks/interfaces.go -package=mocks
type Quizes struct {
	sessionManager sessionmanager.SessionManager
	db             QuizDB
}

func NewQuizes(sessionManager sessionmanager.SessionManager, db QuizDB) Quizes {
	return Quizes{
		sessionManager: sessionManager,
		db:             db,
	}
}

/*
type QuizesTestSetup struct {
	quizes     Quizes
	sessionsDB *mocks.MockSessionsDB
	usersDB    *mocks2.MockUsersDB
	clock      *utilMocks.MockClock
	router     *gin.Engine
}

func InitAccountsTestSetup(t *testing.T) *QuizesTestSetup {
	ctrl := gomock.NewController(t)
	sessionsDB := mocks.NewMockSessionsDB(ctrl)
	clock := utilMocks.NewMockClock(ctrl)
	userDB := mocks2.NewMockUsersDB(ctrl)

	return &QuizesTestSetup{
		quizes:     NewAccounts(sessionmanager.InitMockSessions(sessionsDB, clock), userDB),
		sessionsDB: sessionsDB,
		usersDB:    userDB,
		clock:      clock,
		router:     gin.Default(),
	}
}*/
