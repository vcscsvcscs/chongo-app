package accounts

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vcscsvcscs/chongo-app/backend/controllers/accounts/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const validUserInput = `{"username": "asd", "name": "name", "password": "123", "email": "email"}`

func TestAccounts_Login(t *testing.T) {
	t.Run("email not found", func(t *testing.T) {
		// Given
		endpoint := "/login"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Login)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(validUserInput))

		user := model.User{
			Username:       "asd",
			Name:           "name",
			Email:          "email",
			Password:       "123",
			Ip:             "1.1.1.",
			DeletedAccount: 0,
		}
		setup.usersDB.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).SetArg(1, user).Return(false).Times(1)

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusNotFound, wr.Code)
		assert.Equal(t, "{\"message\":\"No user with this email.\"}", wr.Body.String())
	})

	t.Run("invalid password", func(t *testing.T) {
		// Given
		endpoint := "/login"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Login)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(validUserInput))

		user := model.User{
			Username:       "asd",
			Name:           "name",
			Email:          "email",
			Password:       "123",
			Ip:             "1.1.1.",
			DeletedAccount: 0,
		}
		setup.usersDB.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).SetArg(1, user).Return(true).Times(1)

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusNotFound, wr.Code)
		assert.Equal(t, "{\"message\":\"Password is wrong.\"}", wr.Body.String())
	})

	t.Run("deleted account", func(t *testing.T) {
		// Given
		endpoint := "/login"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Login)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(validUserInput))

		user := model.User{
			Username:       "asd",
			Name:           "name",
			Email:          "email",
			Password:       "$2a$12$FMJjkYIkArXuJ20skrA6buvq2H9luiN62.LM46eSZrLUD4skTFw5m",
			Ip:             "1.1.1.",
			DeletedAccount: 1,
		}
		setup.usersDB.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).SetArg(1, user).Return(true).Times(1)

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusNotFound, wr.Code)
		assert.Equal(t, "{\"message\":\"Account was deleted at: 1\"}", wr.Body.String())
	})

	t.Run("correct login", func(t *testing.T) {
		// Given
		endpoint := "/login"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Login)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(validUserInput))

		user := model.User{
			Username:       "asd",
			Name:           "name",
			Email:          "email",
			Password:       "$2a$12$FMJjkYIkArXuJ20skrA6buvq2H9luiN62.LM46eSZrLUD4skTFw5m",
			Ip:             "1.1.1.",
			DeletedAccount: 0,
		}
		setup.usersDB.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).SetArg(1, user).Return(true).Times(1)
		setup.clock.EXPECT().Now().Return(time.Time{}).AnyTimes()
		setup.sessionsDB.EXPECT().Insert(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusOK, wr.Code)
		assert.Equal(t, "{\"message\":\"User Sign In successfully\",\"token\":\"314df7e428cb2f977c9ba34ef2a8e2b1\"}", wr.Body.String())
	})
}
