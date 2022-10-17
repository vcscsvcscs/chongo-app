package model

import "time"

type Result struct {
	Id       string         `json:"id" bson:"id"`
	Username string         `json:"username" bson:"username"`
	Date     time.Time      `json:"date" bson:"date"`
	Answers  map[string]int `json:"answers" bson:"answers"`
}
