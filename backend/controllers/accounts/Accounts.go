package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mocks2 "github.com/vcscsvcscs/chongo-app/backend/controllers/accounts/mocks"
	"github.com/vcscsvcscs/chongo-app/backend/controllers/accounts/model"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager/mocks"
	utilMocks "github.com/vcscsvcscs/chongo-app/backend/utilities/mocks"
	"testing"
	"time"
)

//go:generate mockgen -source=Accounts.go -destination=./mocks/interfaces.go -package=mocks

type UsersDB interface {
	FindByEmail(email string, user *model.User) bool
	FindByUserName(userName string, user *model.User) bool
	Update(userName string, t time.Time) error
	Insert(user *model.User) error
	RemoveAll(duration time.Duration) error
}

type Accounts struct {
	sessionManager sessionmanager.SessionManager
	db             UsersDB
}

func NewAccounts(sessionManager sessionmanager.SessionManager, db UsersDB) Accounts {
	return Accounts{
		sessionManager: sessionManager,
		db:             db,
	}
}

type AccountsTestSetup struct {
	acc        Accounts
	sessionsDB *mocks.MockSessionsDB
	usersDB    *mocks2.MockUsersDB
	clock      *utilMocks.MockClock
	router     *gin.Engine
}

func InitAccountsTestSetup(t *testing.T) *AccountsTestSetup {
	ctrl := gomock.NewController(t)
	sessionsDB := mocks.NewMockSessionsDB(ctrl)
	clock := utilMocks.NewMockClock(ctrl)
	userDB := mocks2.NewMockUsersDB(ctrl)

	return &AccountsTestSetup{
		acc:        NewAccounts(sessionmanager.InitMockSessions(sessionsDB, clock), userDB),
		sessionsDB: sessionsDB,
		usersDB:    userDB,
		clock:      clock,
		router:     gin.Default(),
	}
}
