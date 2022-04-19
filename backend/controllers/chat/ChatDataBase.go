package chat

import (
	"gopkg.in/mgo.v2"
)

var ChatIds *mgo.Collection
var ChatUsers *mgo.Collection
var Chats *mgo.Collection

type chatMessage struct {
	Token   string `json:"token" bson:"token"`
	Content string `json:"msg" bson:"msg"`
	Id      string `json:"id" bson:"id"`
	Event   string `json:"event" bson:"event"`
}

type updateInfo struct {
	Chats []chatId `json:"Chats" bson:"Chats"`
	Chat  chat     `json:"chat" bson:"chat"`
	Error string   `json:"errors" bson:"errors"`
}

type chatId struct {
	Admin  string `json:"admin" bson:"admin"`
	Id     string `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name"`
	NewMsg bool   `json:"new" bson:"new"`
	Online bool   `json:"online" bson:"online"`
}

type chatUsers struct {
	Id       string `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
}

type chat struct {
	Id       string    `json:"id" bson:"id"`
	Messages []message `json:"msgs" bson:"msgs"`
}

type message struct {
	Time     int64  `json:"sent" bson:"sent"`
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
	Content  string `json:"msg" bson:"msg"`
}
