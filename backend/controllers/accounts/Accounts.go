package accounts

import "github.com/vcscsvcscs/chongo-app/backend/sessionmanager"

type Accounts struct {
	sessionManager sessionmanager.SessionManager
}

func NewAccounts(sessionManager sessionmanager.SessionManager) Accounts {
	return Accounts{sessionManager: sessionManager}
}
