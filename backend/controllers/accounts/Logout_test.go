package accounts

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAccounts_Logout(t *testing.T) {
	t.Run("user not found", func(t *testing.T) {
		// Given
		endpoint := "/logout"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Logout)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint+"?token=123", strings.NewReader(validUserInput))

		setup.sessionsDB.EXPECT().Remove(gomock.Any()).Return(errors.New("user not found"))

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusNotFound, wr.Code)
		assert.Equal(t, "{\"message\":\"There was a problem with logout, please try again and sry for your incovinience.\"}", wr.Body.String())
	})

	t.Run("user found and logged out", func(t *testing.T) {
		// Given
		endpoint := "/logout"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.Logout)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint+"?token=123", strings.NewReader(validUserInput))

		setup.sessionsDB.EXPECT().Remove(gomock.Any()).Return(nil)

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusOK, wr.Code)
		assert.Equal(t, "{\"message\":\"User Sign out successfully\"}", wr.Body.String())
	})
	
	t.Run("test test for checking codecov", func(t *testing.T) {
		assert.Equal(t, true, true)

	})
}
