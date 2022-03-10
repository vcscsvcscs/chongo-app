package chat

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager"
	"gopkg.in/mgo.v2/bson"
)

func ChatSocket(c *gin.Context) {
	var data chatMessage
	var username string
	//Upgrade get request to webSocket protocol
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("There was an error with this connection")
		ws.Close()
	}
	defer ws.Close()
	/////////////////////////////////////////////////////////////////////////////////
	//On open
	err = ws.ReadJSON(&data)
	if err != nil {
		//log.Println("error read json")
		log.Fatal(err)
	}
	username = sessionmanager.Users[data.Token]
	var chatid chatUsers
	var mychats []chatId
	iter := ChatUsers.Find(bson.M{"username": username}).Iter()
	for iter.Next(&chatid) {
		mychats = append(mychats, chatId{Id: chatid.Id})
	}
	for i := 0; i < len(mychats); i++ {
		iter := ChatUsers.Find(bson.M{"id": mychats[i].Id}).Iter()
		for iter.Next(&mychats[i]) {

		}
		sessionmanager.Online[Base.Friends[i].ToUsername]

	}
	/////////////////////////////////////////////////////////////////////////////////
	for {
		err = ws.ReadJSON(&data)
		if err != nil {
			//log.Println("error read json")
			break
		}
		if data.Token == "close" {
			break
		}
		switch data.Event {
		case "Msg":
			///////
		case "Delete":
			//////
		}

	}
}
