package chat

import (
	"encoding/json"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

var server *socketio.Server

func Serverstart() *socketio.Server {
	server = socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("") //userdata and stuff
		log.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "update", func(s socketio.Conn, uinf string) string {
		var info updateInfo
		var update updateInfo
		json.Unmarshal([]byte(uinf), &info)
		toreturn, _ := json.Marshal(update)
		return string(toreturn)
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})
	return server
}
