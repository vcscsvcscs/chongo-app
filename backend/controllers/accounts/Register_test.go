package accounts

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAccounts_Register(t *testing.T) {
	t.Run("invalid email", func(t *testing.T) {
		const invalidUserInput = `{"username": "asd", "name": "name", "password": "123", "email": "email", "ip": "1.1.1.1", "deleted": 0}`

		// Given
		endpoint := "/register"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Register)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(invalidUserInput))

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusNotFound, wr.Code)
		assert.Equal(t, "{\"message\":\"Email is not valid.\"}", wr.Body.String())
	})

	t.Run("duplicate email", func(t *testing.T) {
		const invalidUserInput = `{"username": "asd", "name": "name", "password": "123", "email": "a@a.com", "ip": "1.1.1.1", "deleted": 0}`

		// Given
		endpoint := "/register"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Register)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(invalidUserInput))

		setup.usersDB.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(true)

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusNotFound, wr.Code)
		assert.Equal(t, "{\"message\":\"There is already a user with this email.\"}", wr.Body.String())
	})

	t.Run("duplicate username", func(t *testing.T) {
		const invalidUserInput = `{"username": "asd", "name": "name", "password": "123", "email": "a@a.com", "ip": "1.1.1.1", "deleted": 0}`

		// Given
		endpoint := "/register"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Register)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(invalidUserInput))

		setup.usersDB.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(false)
		setup.usersDB.EXPECT().FindByUserName(gomock.Any(), gomock.Any()).Return(true)

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusNotFound, wr.Code)
		assert.Equal(t, "{\"message\":\"There is already a user with this username.\"}", wr.Body.String())
	})

	t.Run("could not insert user", func(t *testing.T) {
		const invalidUserInput = `{"username": "asd", "name": "name", "password": "123", "email": "a@a.com", "ip": "1.1.1.1", "deleted": 0}`

		// Given
		endpoint := "/register"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Register)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(invalidUserInput))

		setup.usersDB.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(false)
		setup.usersDB.EXPECT().FindByUserName(gomock.Any(), gomock.Any()).Return(false)
		setup.usersDB.EXPECT().Insert(gomock.Any()).Return(errors.New("could not insert user"))

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusNotFound, wr.Code)
		assert.Equal(t, "{\"message\":\"There was an error with our database, please try again, sry for the incovinience.\"}", wr.Body.String())
	})

	t.Run("could not insert session", func(t *testing.T) {
		const invalidUserInput = `{"username": "asd", "name": "name", "password": "123", "email": "a@a.com", "ip": "1.1.1.1", "deleted": 0}`

		// Given
		endpoint := "/register"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Register)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(invalidUserInput))

		setup.usersDB.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(false)
		setup.usersDB.EXPECT().FindByUserName(gomock.Any(), gomock.Any()).Return(false)
		setup.usersDB.EXPECT().Insert(gomock.Any()).Return(nil)
		setup.clock.EXPECT().Now().Return(time.Time{}).AnyTimes()
		setup.sessionsDB.EXPECT().Insert(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("could not insert session"))

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusInternalServerError, wr.Code)
		assert.Equal(t, "{\"message\":\"Internal server error.\"}", wr.Body.String())
	})

	t.Run("create user", func(t *testing.T) {
		const invalidUserInput = `{"username": "asd", "name": "name", "password": "123", "email": "a@a.com", "ip": "1.1.1.1", "deleted": 0}`

		// Given
		endpoint := "/register"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Register)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(invalidUserInput))

		setup.usersDB.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(false)
		setup.usersDB.EXPECT().FindByUserName(gomock.Any(), gomock.Any()).Return(false)
		setup.usersDB.EXPECT().Insert(gomock.Any()).Return(nil)
		setup.clock.EXPECT().Now().Return(time.Time{}).AnyTimes()
		setup.sessionsDB.EXPECT().Insert(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusCreated, wr.Code)
		assert.Equal(t, "{\"message\":\"User Registration successfull\",\"token\":\"314df7e428cb2f977c9ba34ef2a8e2b1\"}", wr.Body.String())
	})
}
