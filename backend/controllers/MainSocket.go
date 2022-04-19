package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type reciever struct {
	Token    string        `json:"token" bson:"token"`
	Types    []string      `json:"types" bson:"types"`
	Contents []interface{} `json:"contents" bson:"contents"`
}

func MainSocket(c *gin.Context) {
	var data reciever
	//var username string
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
		log.Println(&err)
	}
	/*if data.Token == "close" {
		return
	}*/
	/////////////////////////////////////////////////////////////////////////////////
	for {
		err = ws.ReadJSON(data)
		if err != nil {
			//log.Println("error read json")
			break
		}
		for _, Trequest := range data.Types {
			switch Trequest {
			case "Chat":
				///////
			case "Quiz":
				//////
			case "":
				///
			}
		}
		if data.Token == "close" {
			break
		}
	}
}
