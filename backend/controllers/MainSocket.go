package controllers

import (
	"encoding/json"
	"log"

	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager/model"

	socketio "github.com/googollee/go-socket.io"
)

var server *socketio.Server

func MainSocket(sm *sessionmanager.SessionManager) *socketio.Server {
	server = socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		//s.SetContext("") //userdata and stuff
		log.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "authenticate", func(s socketio.Conn, uinf string) {
		var info Authenticationdata
		json.Unmarshal([]byte(uinf), &info)
		token, legit := sm.IsSessionLegit(info.Token)
		if legit {
			s.SetContext(token)
		} else {
			s.Emit("error", "This token is not Legit")
		}
	})
	/*server.OnEvent("/", "update", func(s socketio.Conn, uinf string) string {
		var info string
		var update string
		json.Unmarshal([]byte(uinf), &info)
		toreturn, _ := json.Marshal(update)
		return string(toreturn)
	})*/
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("Meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		token, ok := s.Context().(model.Token)
		if ok {
			sm.DeleteSessionKey(token.Token)
		}
		log.Println("closed", reason)
	})
	return server
}
