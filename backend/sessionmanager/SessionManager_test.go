package sessionmanager

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager/mocks"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager/model"
	utilMocks "github.com/vcscsvcscs/chongo-app/backend/utilities/mocks"
	"testing"
	"time"
)

type SessionManagerTestSetup struct {
	sm    SessionManager
	db    *mocks.MockSessionsDB
	clock *utilMocks.MockClock
}

func (ts *SessionManagerTestSetup) AddUser(token, userName string) {
	newUsers := ts.sm.users
	newUsers[token] = userName

	newOnlines := ts.sm.online
	newOnlines[token] = true

	ts.sm = SessionManager{
		sessions: ts.sm.sessions,
		maxAge:   ts.sm.maxAge,
		users:    newUsers,
		online:   newOnlines,
		clock:    ts.sm.clock,
	}
}

func InitSessionManagerTestSetup(t *testing.T) *SessionManagerTestSetup {
	ctrl := gomock.NewController(t)
	db := mocks.NewMockSessionsDB(ctrl)
	clock := utilMocks.NewMockClock(ctrl)

	return &SessionManagerTestSetup{
		sm:    InitMockSessions(db, clock),
		db:    db,
		clock: clock,
	}
}

func TestSessionManager_SetSessionKeys(t *testing.T) {
	t.Run("should set session keys", func(t *testing.T) {
		// Given
		currentTime := time.Now()
		userName := "randomUserName123"

		setup := InitSessionManagerTestSetup(t)
		setup.clock.EXPECT().Now().Return(currentTime).AnyTimes()
		setup.db.EXPECT().Insert(gomock.Any(), userName, currentTime).Times(1)

		// When
		token, err := setup.sm.SetSessionKeys("1.1.1.1", userName)

		// Then
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, userName, setup.sm.GetUser(token))
	})

	t.Run("error setting session keys", func(t *testing.T) {
		// Given
		currentTime := time.Now()
		userName := "randomUserName123"

		setup := InitSessionManagerTestSetup(t)
		setup.clock.EXPECT().Now().Return(currentTime).AnyTimes()
		setup.db.EXPECT().Insert(gomock.Any(), userName, currentTime).Times(1).Return(errors.New("random error"))

		// When
		token, err := setup.sm.SetSessionKeys("1.1.1.1", userName)

		// Then
		assert.Error(t, err)
		assert.Empty(t, token)
	})
}

func TestSessionManager_IsSessionLegit(t *testing.T) {
	t.Run("legit session", func(t *testing.T) {
		// Given
		currentTime := time.Now()
		userName := "randomUserName123"
		tokenString := "randomToken123"

		setup := InitSessionManagerTestSetup(t)
		setup.AddUser(tokenString, userName)

		setup.db.EXPECT().FindByToken(tokenString, gomock.Any()).SetArg(1, model.Token{Token: tokenString}).Return(true).Times(1)
		setup.clock.EXPECT().Now().Return(currentTime).AnyTimes()
		setup.db.EXPECT().Update(tokenString, currentTime).Times(1)

		// When
		token, isLegit := setup.sm.IsSessionLegit(tokenString)

		// Then
		assert.Equal(t, tokenString, token.Token)
		assert.True(t, isLegit)
	})

	t.Run("not legit session", func(t *testing.T) {
		// Given
		currentTime := time.Now()
		userName := "randomUserName123"
		tokenString := "randomToken123"

		setup := InitSessionManagerTestSetup(t)
		setup.AddUser(tokenString, userName)

		setup.db.EXPECT().FindByToken(tokenString, gomock.Any()).Return(false).Times(1)
		setup.clock.EXPECT().Now().Return(currentTime).AnyTimes()

		// When
		_, isLegit := setup.sm.IsSessionLegit(tokenString)

		// Then
		assert.False(t, isLegit)
	})
}
