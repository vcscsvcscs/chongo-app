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

func TestAccounts_DeleteAccount(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		// Given
		endpoint := "/deleteaccount"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.DeleteAcc)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint+"?token=123", strings.NewReader(validUserInput))

		setup.sessionsDB.EXPECT().Remove(gomock.Any()).Return(errors.New("not found"))

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusNotFound, wr.Code)
		assert.Equal(t, "{\"message\":\"There was a problem with deleting your account, please try again and sry for your incovinience.\"}", wr.Body.String())
	})

	t.Run("can't update", func(t *testing.T) {
		// Given
		endpoint := "/deleteaccount"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.DeleteAcc)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint+"?token=123", strings.NewReader(validUserInput))

		setup.sessionsDB.EXPECT().Remove(gomock.Any()).Return(nil)
		setup.usersDB.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("can't update"))

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusNotFound, wr.Code)
		assert.Equal(t, "{\"message\":\"There was a problem with deleting your account, please try again and sry for your incovinience.\"}", wr.Body.String())
	})

	t.Run("successfully update", func(t *testing.T) {
		// Given
		endpoint := "/deleteaccount"

		setup := InitAccountsTestSetup(t)
		setup.router.POST(endpoint, setup.acc.DeleteAcc)

		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, endpoint+"?token=123", strings.NewReader(validUserInput))

		setup.sessionsDB.EXPECT().Remove(gomock.Any()).Return(nil)
		setup.usersDB.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

		// When
		setup.router.ServeHTTP(wr, req)

		// Then
		assert.Equal(t, http.StatusOK, wr.Code)
		assert.Equal(t, "{\"message\":\"User account deleted succesfully, we delete every data permanently after 2 weeks of calming period\"}", wr.Body.String())
	})
}
